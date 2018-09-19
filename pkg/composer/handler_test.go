// Copyright 2010 Jomoespe. All rights reserved.
// Use of this source code is governed by a WTFPL-style
// license that can be found in the LICENSE file.

package composer_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/jomoespe/webc-spikes/pkg/composer"

	"testing"
)

func TestHandler(t *testing.T) {
	cases := [...]struct {
		name       string
		url        string
		wantStatus int
		wantBody   string
	}{
		{"must salutate Google.com", "http://xxx/google", http.StatusOK, "Hello, http://www.google.com"},
		{"must salutate Boxever.com", "https://xxx/boxever", http.StatusOK, "Hello, https://www.boxever.com"},
		{"path doesn't exist", "http://xxx/unknown", http.StatusNotFound, "\n"},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", tt.url, nil)
			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(composer.Route)
			handler.ServeHTTP(rr, req)

			if rr.Code != tt.wantStatus {
				t.Fatalf("Status: Got [%d], Want [%d]", rr.Code, tt.wantStatus)
			}

			got := rr.Body.String()
			if got != tt.wantBody {
				t.Fatalf("Body: Got [%s], Want [%s]", got, tt.wantBody)
			}
		})
	}
}
