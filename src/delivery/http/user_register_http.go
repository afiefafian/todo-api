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
	Router *httpRouter
	D      RegisterServicesDependencies
}

type RegisterServicesDependencies struct {
	UserServices         entity.UserServices
	RegistrationServices entity.RegistrationServices
}

func userRegisterHTTPRouter(r *httpRouter, d RegisterServicesDependencies) {
	handler := &UserRegisterHandler{
		Router: r,
		D:      d,
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

	// Delete user
	//ctx := r.Context()
	//userID := ps.ByName("id")
	//if err := u.UserServices.Delete(ctx, userID); err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	response.JSONResult(w, &res)
	return
}
