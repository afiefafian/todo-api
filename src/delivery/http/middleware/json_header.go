package middleware

import "net/http"

// JSONHeader set response header type data to json
func JSONHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
