package middleware

import (
	"fmt"
	"net/http"
)

func TheHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("the-handler")
}

func TheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("the-middleware: before")
		next.ServeHTTP(w, r)
		fmt.Println("the-middleware: after")
	})
}
