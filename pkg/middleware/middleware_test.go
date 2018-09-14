package middleware

import (
	"net/http"
	"net/http/httptest"
)

func ExampleMiddleware() {

	req, _ := http.NewRequest("GET", "http://sample.com/path/resource", nil)
	w := httptest.NewRecorder()

	h := http.HandlerFunc(TheHandler)
	TheMiddleware(h).ServeHTTP(w, req)

	// Output:
	// the-middleware: before
	// the-handler
	// the-middleware: after
}
