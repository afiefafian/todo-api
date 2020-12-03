package http

import (
	"encoding/json"
	"net/http"
	response "todo_api/src/delivery/http/response"
	"todo_api/src/entity"

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
	//r.Router.GET("/users/:id", handler.GetUserByID)
	//r.Router.POST("/users", handler.Store)
	//r.Router.PUT("/users/:id", handler.Update)
	//r.Router.DELETE("/users/:id", handler.Delete)
}

// FetchUsers http routing handler for get user
func (u *UserHandler) FetchUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res := response.ResultSuccess

	ctx := r.Context()
	data, err := u.UserServices.Fetch(ctx)
	if err != nil {
		response.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	res.Data = data
	response.JSONResult(w, res)
	return
}

// GetUserByID : http routing handler for get user
func (u *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	res := response.ResultSuccess

	ctx := r.Context()
	userID := ps.ByName("id")
	data, err := u.UserServices.GetByID(ctx, userID)
	if err != nil {
		response.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	res.Data = data
	response.JSONResult(w, &res)
	return
}

// Store will store the article by given request body
func (u *UserHandler) Store(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	res := response.ResultSuccess
	res.Message = "Success create an user"

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

	response.JSONResult(w, &res)
	return
}

// Update will update user data by given request body
func (u *UserHandler) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	res := response.ResultSuccess
	res.Message = "Success update an user"

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

	response.JSONResult(w, &res)
	return
}

// Delete will delete the user by id
func (u *UserHandler) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	res := response.ResultSuccess
	res.Message = "Success delete an user"

	// Delete user
	ctx := r.Context()
	userID := ps.ByName("id")
	if err := u.UserServices.Delete(ctx, userID); err != nil {
		response.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.JSONResult(w, &res)
	return
}
