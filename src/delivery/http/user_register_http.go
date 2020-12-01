package http

import (
	"net/http"
	"todo_api/src/entity"

	"github.com/julienschmidt/httprouter"
)

// UserHandler handling user services
type UserRegisterHandler struct {
	Router       *httpRouter
	UserServices entity.UserServices
}

func userRegisterHTTPRouter(r *httpRouter, u entity.UserServices) {
	handler := &UserRegisterHandler{
		Router:       r,
		UserServices: u,
	}

	r.Router.POST("/users/register", handler.RegisterUser)
	//r.Router.POST("/users/register/resend-code", handler.RegisterUser)
	//r.Router.POST("/users/register/verify", handler.RegisterUser)
}

// RegisterUser by user data
func (u *UserRegisterHandler) RegisterUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	response := resultSuccess
	response.Name = "Success send registration code"

	// Delete user
	//ctx := r.Context()
	//userID := ps.ByName("id")
	//if err := u.UserServices.Delete(ctx, userID); err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	JSONResult(w, &response)
	return
}
