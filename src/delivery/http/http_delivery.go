package http

import (
	"net/http"

	middleware "todo_api/src/delivery/http/middleware"
	pgsqlRepo "todo_api/src/repository/postgres"
	services "todo_api/src/services"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/go-pg/pg/v10"
	"github.com/julienschmidt/httprouter"
)

// InitRoute initialize http route handler
func InitRoute(r *httprouter.Router, db *pg.DB) {
	r.GlobalOPTIONS = http.HandlerFunc(middleware.CORS)

	// User route
	userRepo := pgsqlRepo.NewPostgresUserRepository(db)
	userServices := services.NewUserServices(userRepo)
	userHTTPRouter(r, userServices)
}

// JSONResult return result in json format
// Accepted resultByte type data is slice byte
func JSONResult(w http.ResponseWriter, resultByte []byte, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(resultByte)
}

func validateRequest(m interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
