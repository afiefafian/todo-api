package http

import (
	"encoding/json"
	"net/http"
	"todo_api/src/entity"

	"github.com/julienschmidt/httprouter"
)

// UserHandler handling user services
type UserRegisterHandler struct {
	UserServices entity.UserServices
}

func userRegisterHTTPRouter(r *httprouter.Router, u entity.UserServices) {
	handler := &UserRegisterHandler{
		UserServices: u,
	}

	r.POST("/users/register", handler.RegisterUser
	//r.POST("/users/register/resend-code", handler.RegisterUser)
	//r.POST("/users/register/verify", handler.RegisterUser)
}

// RegisterUser
func (u *UserRegisterHandler) RegisterUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()
	data, err := u.UserServices.Fetch(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Marshal data to []byte
	result, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write result to http.ResponseWrite
	w.Write(result)
	return
}
