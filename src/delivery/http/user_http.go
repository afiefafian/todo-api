package http

import (
	"encoding/json"
	"github.com/afiefafian/todo-api/src/delivery/http/response"
	"github.com/afiefafian/todo-api/src/entity"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// UserHandler handling user services
type UserHandler struct {
	Router       *httpRouter
	UserServices entity.UserServices
}

func userHTTPRouter(r *httpRouter, u entity.UserServices) {
	handler := &UserHandler{
		Router:       r,
		UserServices: u,
	}

	r.Router.GET("/users", handler.FetchUsers)
	r.Router.GET("/users/:id", handler.GetUserByID)
	r.Router.POST("/users", handler.Store)
	r.Router.PUT("/users/:id", handler.Update)
	r.Router.DELETE("/users/:id", handler.Delete)
}

// FetchUsers http routing handler for get user
func (u *UserHandler) FetchUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()
	data, err := u.UserServices.Fetch(ctx)
	if err != nil {
		response.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := response.ResultSuccess("")
	res.SetData(data)
	response.JSONResult(w, res)
	return
}

// GetUserByID : http routing handler for get user
func (u *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	userID := ps.ByName("id")
	data, err := u.UserServices.GetByID(ctx, userID)
	if err != nil {
		response.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := response.ResultSuccess("")
	res.SetData(data)
	response.JSONResult(w, &res)
	return
}

// Store will store the article by given request body
func (u *UserHandler) Store(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user entity.User
	json.NewDecoder(r.Body).Decode(&user)

	// Validation
	if ok := response.ValidateAndReturnErrResultHTTPWithErrField(w, &user); !ok {
		return
	}

	// Create user
	ctx := r.Context()
	if err := u.UserServices.Store(ctx, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := response.ResultSuccess("Success create an user")
	response.JSONResult(w, &res)
	return
}

// Update will update user data by given request body
func (u *UserHandler) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user entity.User
	json.NewDecoder(r.Body).Decode(&user)

	// Validation
	if ok := response.ValidateAndReturnErrResultHTTPWithErrField(w, &user); !ok {
		return
	}

	// Update user
	ctx := r.Context()
	userID := ps.ByName("id")
	if err := u.UserServices.Update(ctx, &user, userID); err != nil {
		response.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := response.ResultSuccess("Success update an user")
	response.JSONResult(w, &res)
	return
}

// Delete will delete the user by id
func (u *UserHandler) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Delete user
	ctx := r.Context()
	userID := ps.ByName("id")
	if err := u.UserServices.Delete(ctx, userID); err != nil {
		response.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := response.ResultSuccess("Success delete an user")
	response.JSONResult(w, &res)
	return
}
