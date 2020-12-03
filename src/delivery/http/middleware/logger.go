package middleware

import (
	"log"
	"net/http"
	"time"
)

// HTTPLogger middleware print http request to console
func HTTPLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
		h.ServeHTTP(w, r)
	})
}
