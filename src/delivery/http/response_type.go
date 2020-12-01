package http

import "net/http"

type response map[string]interface{}

type resultFormat struct {
	Status  int         `json:"status"`
	Name    string      `json:"name"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   response    `json:"error,omitempty"`
	Meta    response    `json:"meta,omitempty"`
}

type paginateMeta struct {
}

// resultSuccess is a sample of result success
// Beware edit the result
// Please create new variable based on sample below
// res := resultSuccess
var resultSuccess = resultFormat{
	Status:  http.StatusOK,
	Name:    http.StatusText(http.StatusOK),
	Message: http.StatusText(http.StatusOK),
}

// resultErr is a sample of result success
var resultErr = resultFormat{
	Status: http.StatusBadRequest,
	Name:   http.StatusText(http.StatusBadRequest),
}
