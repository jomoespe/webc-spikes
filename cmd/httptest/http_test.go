package main_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// a http handelr to test
	f := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", "jomoespe")
	}

	cases := [...]struct {
		name string
		want string
	}{
		{"a test with httptest", "jomoespe"},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "http://sample.com/path/resource", nil)
			w := httptest.NewRecorder()
			f(w, req)
			got := w.Body.String()
			if got != tt.want {
				t.Errorf("Got %s, Want %s", got, tt.want)
			}
		})
	}
}
