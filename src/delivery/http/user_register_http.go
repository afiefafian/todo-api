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

	r.Router.POST("/users/register", handler.RegisterUser)
	//r.Router.POST("/users/register/resend-code", handler.RegisterUser)
	//r.Router.POST("/users/register/verify", handler.RegisterUser)
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
	if err := u.RegistrationServices.Register(ctx, &user); err != nil {
		response.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.JSONResult(w, &res)
	return
}
