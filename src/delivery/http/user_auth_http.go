package http

import (
	"encoding/json"
	"github.com/afiefafian/todo-api/src/delivery/http/response"
	"github.com/afiefafian/todo-api/src/entity"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// UserHandler handling user services
type UserAuthHandler struct {
	Router           *httpRouter
	UserAuthServices entity.UserAuthServices
}

func userAuthHTTPRouter(r *httpRouter, u entity.UserAuthServices) {
	handler := &UserAuthHandler{
		Router:           r,
		UserAuthServices: u,
	}

	r.Router.POST("/users/auth", handler.Authentication)
	r.Router.POST("/users/logout", handler.Logout)
}

// FetchUsers http routing handler for get user
func (u *UserAuthHandler) Authentication(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var userLogin entity.UserLogin
	json.NewDecoder(r.Body).Decode(&userLogin)
	// Validation
	if ok := response.ValidateAndReturnErrResultHTTPWithErrField(w, &userLogin); !ok {
		return
	}

	var (
		user  entity.User
		token string
		err   error
		ctx   = r.Context()
	)

	if user, token, err = u.UserAuthServices.Authentication(ctx, &userLogin); err != nil {
		response.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	data := response.Response{
		"user": user,
		"token": response.Response{
			"jwt":     token,
			"refresh": "",
		},
	}
	res := response.ResultSuccess("")
	res.SetData(data)

	response.JSONResult(w, res)
	return
}

// FetchUsers http routing handler for get user
func (u *UserAuthHandler) Logout(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	res := response.ResultSuccess

	//ctx := r.Context()
	//data, err := u.UserServices.Fetch(ctx)
	//if err != nil {
	//	response.JSONError(w, err.Error(), http.StatusBadRequest)
	//	return
	//}

	//res.Data = data
	response.JSONResult(w, res)
	return
}
