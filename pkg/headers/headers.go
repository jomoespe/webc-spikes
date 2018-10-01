// Copyright 2010 Jomoespe. All rights reserved.
// Use of this source code is governed by a WTFPL-style
// license that can be found in the LICENSE file.

// Package headers contains samples about how to deal with HTTP request and
// response headers. For example:
//   · Filter request headers from a []string
package headers

import (
	"net/http"
)

func Filter(h http.Header, excludes []string) http.Header {
	header := make(map[string][]string)
	for key, value := range h {
		header[key] = value
	}
	for _, exclude := range excludes {
		delete(header, exclude)
	}
	return header
}
