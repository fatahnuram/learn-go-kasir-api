package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func SimpleLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s\n",
			time.Now().Format("2006-01-02 15:04:05"),
			r.Method,
			r.URL.Path,
		)
		next.ServeHTTP(w, r)
	})
}
