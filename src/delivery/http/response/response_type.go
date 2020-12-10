package response

import "net/http"

type Response map[string]interface{}

// ResultFormat struct definition
type ResultFormat struct {
	status  int         `json:"status"`
	name    string      `json:"name"`
	message interface{} `json:"message"`
	data    interface{} `json:"data,omitempty"`
	errors  interface{} `json:"errors,omitempty"`
	meta    interface{} `json:"meta,omitempty"`
	trace   interface{} `json:"trace,omitempty"`
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
		status:  http.StatusOK,
		name:    http.StatusText(http.StatusOK),
		message: message,
	}
}

// ResultErr is a sample of success result
var ResultErr = func(msg string) *ResultFormat {
	message := http.StatusText(http.StatusBadRequest)
	if msg != "" {
		message = msg
	}

	return &ResultFormat{
		status:  http.StatusBadRequest,
		name:    http.StatusText(http.StatusBadRequest),
		message: message,
	}
}

// resultNotFound is a template for Not Found error
var ResultNotFound = ResultFormat{
	status:  http.StatusNotFound,
	name:    "NotFound",
	message: http.StatusText(http.StatusNotFound),
}

// ResultInternalServerErr is a template for Error 500 Internal Server Error
var ResultInternalServerErr = ResultFormat{
	status:  http.StatusInternalServerError,
	name:    "InternalServerErr",
	message: http.StatusText(http.StatusInternalServerError),
}

// ResultValidationErr is a template for validation form error
var ResultValidationErr = ResultFormat{
	status:  http.StatusBadRequest,
	name:    "ValidationErr",
	message: "Validation Error",
}

func (r *ResultFormat) SetStatus(status int) {
	r.status = status
}

func (r *ResultFormat) SetName(name string) {
	r.name = name
}

func (r *ResultFormat) SetMessage(msg interface{}) {
	r.message = msg
}

func (r *ResultFormat) SetData(data interface{}) {
	r.data = data
}

func (r *ResultFormat) SetErrors(errors interface{}) {
	r.errors = errors
}
