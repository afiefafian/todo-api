package http

import (
	"github.com/afiefafian/todo-api/src/delivery/http/middleware"
	"github.com/afiefafian/todo-api/src/delivery/http/response"
	pgsqlRepo "github.com/afiefafian/todo-api/src/repository/postgres"
	redisRepo "github.com/afiefafian/todo-api/src/repository/redis"
	"github.com/afiefafian/todo-api/src/services"
	"net/http"

	"github.com/go-redis/redis/v8"

	"github.com/go-pg/pg/v10"
	"github.com/julienschmidt/httprouter"
)

type httpRouter struct {
	Router httprouter.Router
}

// NewRouter initialize http route handler
func NewRouter(db *pg.DB, inMem *redis.Client) *httpRouter {
	r := &httpRouter{
		Router: *httprouter.New(),
	}

	// Set Global Middleware
	r.Router.GlobalOPTIONS = http.HandlerFunc(middleware.CORS)
	r.Router.PanicHandler = response.PanicHandler()
	r.Router.NotFound = http.HandlerFunc(response.NotFoundHandler)
	r.Router.MethodNotAllowed = http.HandlerFunc(response.NotFoundHandler)

	setRoute(r, db, inMem)

	return r
}

func setRoute(r *httpRouter, db *pg.DB, inMem *redis.Client) {
	// Init repository
	userRepo := pgsqlRepo.NewPostgresUserRepository(db)
	registrationRepo := pgsqlRepo.NewPostgresRegistrationRepository(db)
	tokenRepo := pgsqlRepo.NewPostgresTokenRepository(db)

	memRegistrationRepo := redisRepo.NewRedisRegistrationRepository(inMem)

	// Init services
	rs := services.NewRegistrationServices(userRepo, registrationRepo, tokenRepo, memRegistrationRepo)
	us := services.NewUserServices(userRepo)
	aus := services.NewUserAuthServices(userRepo)

	// Set Route
	userHTTPRouter(r, us)
	userAuthHTTPRouter(r, aus)
	userRegisterHTTPRouter(r, rs, aus, us)
}
