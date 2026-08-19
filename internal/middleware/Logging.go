package middleware

import (
	"log"
	"net/http"
)

func Logging(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL.String())
		next.ServeHTTP(w, r)
	})
}
