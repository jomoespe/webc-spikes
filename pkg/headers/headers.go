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
	// create an map with the filter to excludes
	filter := make(map[string]struct{}, len(excludes))
	for _, s := range excludes {
		filter[s] = struct{}{}
	}

	header := make(map[string][]string)
	for key, value := range h {
		if _, ok := filter[key]; !ok {
			header[key] = value
		}
	}
	return header
}
