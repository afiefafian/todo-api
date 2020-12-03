package response

import "net/http"

type Response map[string]interface{}

type ResultFormat struct {
	Status  int         `json:"status"`
	Name    string      `json:"name"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Meta    Response    `json:"meta,omitempty"`
}

type paginateMeta struct {
}

// resultSuccess is a sample of result success
// Beware edit the result
// Please create new variable based on sample below
// res := resultSuccess
var ResultSuccess = ResultFormat{
	Status:  http.StatusOK,
	Name:    http.StatusText(http.StatusOK),
	Message: http.StatusText(http.StatusOK),
}

// resultErr is a sample of result success
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

var ResultInternalServerErr = ResultFormat{
	Status:  http.StatusInternalServerError,
	Name:    "InternalServerErr",
	Message: http.StatusText(http.StatusInternalServerError),
}

var ResultValidationErr = ResultFormat{
	Status:  http.StatusBadRequest,
	Name:    "ValidationErr",
	Message: "Validation Error",
}
