// Copyright 2010 Jomoespe. All rights reserved.
// Use of this source code is governed by a WTFPL-style
// license that can be found in the LICENSE file.

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

func ExampleMiddlewareWithValue() {
	req, _ := http.NewRequest("GET", "http://sample.com/path/resource", nil)
	w := httptest.NewRecorder()

	h := http.HandlerFunc(TheHandlerWithContext)
	TheMiddlewareWithContext(h).ServeHTTP(w, req)

	// Output:
	// the-middleware: before
	// context value=the value in the context
	// the-middleware: after
}
