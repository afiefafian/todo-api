package http

import (
	"encoding/json"
	"net/http"
	"reflect"

	"todo_api/src/delivery/http/middleware"
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
	r.Router.GlobalOPTIONS = http.HandlerFunc(middleware.CORS)
	// User route
	userRepo := pgsqlRepo.NewPostgresUserRepository(db)
	userServices := services.NewUserServices(userRepo)
	userRegisterHTTPRouter(r, userServices)
	userHTTPRouter(r, userServices)

	return r
}

// JSONResult return result in json format
// Accepted resultByte type data is slice byte
func JSONResult(w http.ResponseWriter, resultPtr interface{}) (err error) {
	// Parse data
	resultBytePtr, code := parse(resultPtr)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(*resultBytePtr)
	return
}

// Parsing result data to byte and return status code
func parse(resultBytePtr interface{}) (*[]byte, int) {
	result, err := json.Marshal(resultBytePtr)
	if err != nil {
		response := resultFormat{
			Status:  http.StatusBadRequest,
			Name:    http.StatusText(http.StatusBadRequest),
			Message: "Failed parse data",
		}
		errRes, _ := json.Marshal(response)
		return &errRes, http.StatusBadRequest
	}

	code := http.StatusOK
	if resultBytePtr != nil {
		code = getStatusCodeFromResult(resultBytePtr)
	}

	return &result, code
}

// setStatusCodeFromResult set status code from result data
func getStatusCodeFromResult(resultBytePtr interface{}) int {
	resultReflect := reflect.ValueOf(resultBytePtr)
	// Get reflect data from pointer
	if resultReflect.Kind() == reflect.Ptr {
		resultReflect = resultReflect.Elem()
	}

	if resultReflect.Kind() == reflect.Struct {
		// Get info by field name
		if _, valid := resultReflect.Type().FieldByName("Status"); valid {
			status := resultReflect.FieldByName("Status").Int()
			return int(status)
		}
	}

	return http.StatusOK
}
