package http

import (
	"encoding/json"
	"fmt"
	"github.com/afiefafian/todo-api/src/delivery/http/response"
	"github.com/afiefafian/todo-api/src/entity"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// UserHandler handling user services
type UserRegisterHandler struct {
	Router               *httpRouter
	UserServices         entity.UserServices
	UserAuthServices     entity.UserAuthServices
	RegistrationServices entity.RegistrationServices
}

func userRegisterHTTPRouter(r *httpRouter, rs entity.RegistrationServices, ua entity.UserAuthServices, u entity.UserServices) {
	handler := &UserRegisterHandler{
		Router:               r,
		UserServices:         u,
		RegistrationServices: rs,
		UserAuthServices:     ua,
	}

	r.Router.POST("/users/register/", handler.RegisterUser)
	r.Router.POST("/users/register/resend-code", handler.ResendCode)
	r.Router.POST("/users/register/verify", handler.VerifyCode)
}

// RegisterUser register user and send confirmation code
func (u *UserRegisterHandler) RegisterUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user entity.Registration
	json.NewDecoder(r.Body).Decode(&user)

	// Validation
	if ok := response.ValidateAndReturnErrResultHTTPWithErrField(w, &user); !ok {
		return
	}

	// Send register otp to user email
	ctx := r.Context()
	if _, err := u.RegistrationServices.Register(ctx, &user); err != nil {
		response.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := response.ResultSuccess("Success send registration code")
	response.JSONResult(w, &res)
	return
}

// RegisterUser register user and send confirmation code
func (u *UserRegisterHandler) ResendCode(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var user entity.RegistrationResendValidation
	json.NewDecoder(r.Body).Decode(&user)

	// Validation
	if ok := response.ValidateAndReturnErrResultHTTPWithErrField(w, &user); !ok {
		return
	}

	// Send register otp to user email
	ctx := r.Context()
	email := user.Email
	if _, err := u.RegistrationServices.ResendCode(ctx, email); err != nil {
		response.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := response.ResultSuccess("Success resend registration code")
	response.JSONResult(w, &res)
	return
}

// VerifyCode verify registration code
func (u *UserRegisterHandler) VerifyCode(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var userVerify entity.RegistrationVerifyValidation
	json.NewDecoder(r.Body).Decode(&userVerify)

	// Validation
	if ok := response.ValidateAndReturnErrResultHTTPWithErrField(w, &userVerify); !ok {
		return
	}

	ctx := r.Context()
	email := userVerify.Email
	code := userVerify.Code
	if err := u.RegistrationServices.VerifyCode(ctx, email, code); err != nil {
		response.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user entity.User
	var err error
	if user, err = u.UserServices.GetByEmail(ctx, email); err != nil {
		response.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	var token string
	identifier := fmt.Sprintf("%s", user.ID)
	if token, err = u.UserAuthServices.GenerateAuthToken(identifier); err != nil {
		response.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Password = ""
	data := response.Response{
		"user": user,
		"token": response.Response{
			"jwt":     token,
			"refresh": "",
		},
	}

	res := response.ResultSuccess("Success register new user")
	res.SetData(data)

	response.JSONResult(w, &res)
	return
}
