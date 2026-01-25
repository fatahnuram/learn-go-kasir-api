package middleware

import "net/http"

func DefaultHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Server", "Go/1.24 net/http")
		next.ServeHTTP(w, r)
	})
}
