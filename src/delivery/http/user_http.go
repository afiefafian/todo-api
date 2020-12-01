package http

import (
	"encoding/json"
	"net/http"
	"todo_api/src/entity"

	"github.com/julienschmidt/httprouter"
)

// UserHandler handling user services
type UserHandler struct {
	UserServices entity.UserServices
}

func userHTTPRouter(r *httprouter.Router, u entity.UserServices) {
	handler := &UserHandler{
		UserServices: u,
	}

	r.GET("/users", handler.FetchUsers)
	r.GET("/users/:id", handler.GetUserByID)
	r.POST("/users", handler.Store)
	r.PUT("/users/:id", handler.Update)
	r.DELETE("/users/:id", handler.Delete)
}

// FetchUsers http routing handler for get user
func (u *UserHandler) FetchUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

// GetUserByID : http routing handler for get user
func (u *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	response := resultFormat{
		Status:  http.StatusOK,
		Name:    http.StatusText(http.StatusOK),
		Message: http.StatusText(http.StatusOK),
	}

	ctx := r.Context()
	userID := ps.ByName("id")
	data, err := u.UserServices.GetByID(ctx, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Data = data

	// Marshal data to []byte
	result, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResult(w, result, http.StatusOK)
	return
}

// Store will store the article by given request body
func (u *UserHandler) Store(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	response := resultFormat{
		Status:  http.StatusCreated,
		Name:    http.StatusText(http.StatusOK),
		Message: "Success create an user",
	}

	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)

	// Validation
	var ok bool
	if ok, err = validateRequest(&user); !ok {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create user
	ctx := r.Context()
	if err := u.UserServices.Store(ctx, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Marshal
	result, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResult(w, result, http.StatusCreated)
	return
}

// Update will update user data by given request body
func (u *UserHandler) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	response := resultFormat{
		Status:  http.StatusOK,
		Name:    http.StatusText(http.StatusOK),
		Message: "Success update an user",
	}

	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)

	// Validation
	var ok bool
	if ok, err = validateRequest(&user); !ok {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update user
	ctx := r.Context()
	userID := ps.ByName("id")
	if err := u.UserServices.Update(ctx, &user, userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Marshal result
	result, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResult(w, result, http.StatusCreated)
	return
}

// Delete will delete the user by id
func (u *UserHandler) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	response := resultFormat{
		Status:  http.StatusOK,
		Name:    http.StatusText(http.StatusOK),
		Message: "Success delete an user",
	}

	// Delete user
	ctx := r.Context()
	userID := ps.ByName("id")
	if err := u.UserServices.Delete(ctx, userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Marshal result
	result, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResult(w, result, http.StatusOK)
	return
}
