package response

import "net/http"

type Response map[string]interface{}

// ResultFormat struct definition
type ResultFormat struct {
	Status  int         `json:"status"`
	Name    string      `json:"name"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
	Trace   interface{} `json:"trace,omitempty"`
}

type paginateMeta struct {
}

// ResultSuccess is a sample of result success.
// Beware! Don't change the variable directly.
// Please create new variable based on sample below.
// example: res := resultSuccess
var ResultSuccess = ResultFormat{
	Status:  http.StatusOK,
	Name:    http.StatusText(http.StatusOK),
	Message: http.StatusText(http.StatusOK),
}

// ResultErr is a sample of success result
var ResultErr = ResultFormat{
	Status:  http.StatusBadRequest,
	Name:    http.StatusText(http.StatusBadRequest),
	Message: http.StatusText(http.StatusBadRequest),
}

// resultNotFound is a template for Not Found error
var ResultNotFound = ResultFormat{
	Status:  http.StatusNotFound,
	Name:    "NotFound",
	Message: http.StatusText(http.StatusNotFound),
}

// ResultInternalServerErr is a template for Error 500 Internal Server Error
var ResultInternalServerErr = ResultFormat{
	Status:  http.StatusInternalServerError,
	Name:    "InternalServerErr",
	Message: http.StatusText(http.StatusInternalServerError),
}

// ResultValidationErr is a template for validation form error
var ResultValidationErr = ResultFormat{
	Status:  http.StatusBadRequest,
	Name:    "ValidationErr",
	Message: "Validation Error",
}
