package http

import (
	"encoding/json"
	"net/http"
	"todo_api/src/delivery/http/response"
	"todo_api/src/entity"

	"github.com/julienschmidt/httprouter"
)

// UserHandler handling user services
type UserRegisterHandler struct {
	Router               *httpRouter
	RegistrationServices entity.RegistrationServices
}

func userRegisterHTTPRouter(r *httpRouter, rs entity.RegistrationServices) {
	handler := &UserRegisterHandler{
		Router:               r,
		RegistrationServices: rs,
	}

	r.Router.POST("/users/register/", handler.RegisterUser)
	r.Router.POST("/users/register/resend-code", handler.ResendCode)
	r.Router.POST("/users/register/verify", handler.VerifyCode)
}

// RegisterUser register user and send confirmation code
func (u *UserRegisterHandler) RegisterUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res := response.ResultSuccess
	res.Message = "Success send registration code"

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

	response.JSONResult(w, &res)
	return
}

// RegisterUser register user and send confirmation code
func (u *UserRegisterHandler) ResendCode(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res := response.ResultSuccess
	res.Message = "Success resend registration code"

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

	response.JSONResult(w, &res)
	return
}

// VerifyCode verify registration code
func (u *UserRegisterHandler) VerifyCode(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res := response.ResultSuccess
	res.Message = "Success register new user"

	var user entity.RegistrationVerifyValidation
	json.NewDecoder(r.Body).Decode(&user)

	// Validation
	if ok := response.ValidateAndReturnErrResultHTTPWithErrField(w, &user); !ok {
		return
	}

	ctx := r.Context()
	email := user.Email
	code := user.Code
	if err := u.RegistrationServices.VerifyCode(ctx, email, code); err != nil {
		response.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.JSONResult(w, &res)
	return
}
