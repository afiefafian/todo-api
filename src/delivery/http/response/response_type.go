package response

import "net/http"

// Response map type declaration
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
var ResultSuccess = func(msg string) *ResultFormat {
	message := http.StatusText(http.StatusOK)
	if msg != "" {
		message = msg
	}

	return &ResultFormat{
		Status:  http.StatusOK,
		Name:    http.StatusText(http.StatusOK),
		Message: message,
	}
}

// ResultErr is a sample of success result
var ResultErr = func(msg string) *ResultFormat {
	message := http.StatusText(http.StatusBadRequest)
	if msg != "" {
		message = msg
	}

	return &ResultFormat{
		Status:  http.StatusBadRequest,
		Name:    http.StatusText(http.StatusBadRequest),
		Message: message,
	}
}

// ResultNotFound is a template for Not Found error
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

// SetStatus set status code on ResultFormat struct
func (r *ResultFormat) SetStatus(status int) {
	r.Status = status
}

// SetName set name on ResultFormat struct
func (r *ResultFormat) SetName(name string) {
	r.Name = name
}

// SetMessage set message on ResultFormat struct
func (r *ResultFormat) SetMessage(msg interface{}) {
	r.Message = msg
}

// SetData set data on ResultFormat struct
func (r *ResultFormat) SetData(data interface{}) {
	r.Data = data
}

// SetErrors ser error message on ResultFormat struct
func (r *ResultFormat) SetErrors(errors interface{}) {
	r.Errors = errors
}
