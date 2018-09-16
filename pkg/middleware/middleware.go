package middleware

import (
	"context"
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

type contextKeyType struct{}

// The key where the header value is stored. This is globally unique since
// it uses a custom unexported type. The struct{} costs zero allocations.
var contextKey = contextKeyType(struct{}{})

func TheMiddlewareWithContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), contextKey, "the value in the context")
		r = r.WithContext(ctx)

		fmt.Println("the-middleware: before")
		next.ServeHTTP(w, r)
		fmt.Println("the-middleware: after")
	})
}

func TheHandlerWithContext(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("context value=%v\n", r.Context().Value(contextKey))
}
