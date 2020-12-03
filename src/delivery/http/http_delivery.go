package http

import (
	"net/http"
	"todo_api/src/delivery/http/middleware"
	response2 "todo_api/src/delivery/http/response"
	pgsqlRepo "todo_api/src/repository/postgres"
	"todo_api/src/services"

	"github.com/go-pg/pg/v10"
	"github.com/julienschmidt/httprouter"
)

type httpRouter struct {
	Router httprouter.Router
}

// NewRouter initialize http route handler
func NewRouter(db *pg.DB) *httpRouter {
	r := &httpRouter{*httprouter.New()}

	// Set Global Middleware
	r.Router.GlobalOPTIONS = http.HandlerFunc(middleware.CORS)
	r.Router.PanicHandler = response2.PanicHandler()
	r.Router.NotFound = http.HandlerFunc(response2.NotFoundHandler)

	setRoute(r, db)

	return r
}

func setRoute(r *httpRouter, db *pg.DB) {
	// Init repository
	userRepo := pgsqlRepo.NewPostgresUserRepository(db)
	registrationRepo := pgsqlRepo.NewPostgresRegistrationRepository(db)

	// Init services
	rs := services.NewRegistrationServices(userRepo, registrationRepo)
	us := services.NewUserServices(userRepo)

	// Set Route
	userHTTPRouter(r, us)
	userRegisterHTTPRouter(r, rs)
}
