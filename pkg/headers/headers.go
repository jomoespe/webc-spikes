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

type MyHeader http.Header

func (h MyHeader) Copy(header *http.Header) {
	for key, value := range *header {
		// TODO should y copy the values to a temp before that (https://golang.org/src/net/http/header.go clone function)
		h[key] = value
	}
}

func (h MyHeader) Filter(excludes []string) {
	for _, exclude := range excludes {
		delete(h, exclude)
	}
}
