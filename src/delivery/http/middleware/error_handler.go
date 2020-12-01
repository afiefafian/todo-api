package middleware

import "net/http"

// JSONHeader set response header type data to json
func ErrorHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
