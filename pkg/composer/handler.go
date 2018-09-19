// Copyright 2010 Jomoespe. All rights reserved.
// Use of this source code is governed by a WTFPL-style
// license that can be found in the LICENSE file.

package composer

import (
	"fmt"
	"html"
	"net/http"
)

var routes = map[string]string{
	"/google":  "http://www.google.com",
	"/boxever": "https://www.boxever.com",
}

func route(w http.ResponseWriter, r *http.Request) {
	url := html.EscapeString(r.URL.Path)
	target, ok := routes[url]
	if !ok {
		// path to route don't found. We send a not found status
		http.Error(w, "", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello, %s", target)
}
