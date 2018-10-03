package headers_test

import (
	"net/http"
	"testing"

	"github.com/jomoespe/webc-spikes/pkg/headers"
)

func TestCanFilterMyHeader(t *testing.T) {
	var baseHeader http.Header = map[string][]string{
		"header-1": []string{"value 1.1", "value 1.2"},
		"header-2": []string{"value 2.1"},
		"header-3": []string{"value 3.1", "value 3.2", "value 3.3"},
	}

	cases := [...]struct {
		name   string
		header http.Header
		filter []string
		want   int // expected number of headers after applying the filter
	}{
		{"Empty header and no filter", http.Header{}, []string{}, 0},
		{"Empty header and filter with two values", http.Header{}, []string{"header-1", "headerr-2"}, 0},
		{"Header with three headers and no filter", baseHeader, []string{}, 3},
		{"Header with three headers and removing first", baseHeader, []string{"header-1"}, 2},
		{"Header with three headers and removing middle", baseHeader, []string{"header-2"}, 2},
		{"Header with three headers and removing last", baseHeader, []string{"header-3"}, 2},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			h := headers.MyHeader{}
			h.Copy(&tt.header)
			h.Filter(tt.filter)
			if len(h) != tt.want {
				t.Fatalf("Got: %d, want: %d", len(h), tt.want)
			}
		})
	}
}

func BenchmarkCopyMyHeaders(b *testing.B) {
	var header http.Header = map[string][]string{
		"header-1": []string{"value 1.1", "value 1.2"},
		"header-2": []string{"value 2.1"},
		"header-3": []string{"value 3.1", "value 3.2", "value 3.3"},
		"header-4": []string{"value 1.1", "value 1.2"},
		"header-5": []string{"value 2.1"},
		"header-6": []string{"value 3.1", "value 3.2", "value 3.3", "value 3.1", "value 3.2", "value 3.3"},
	}

	for n := 0; n < b.N; n++ {
		h := headers.MyHeader{}
		h.Copy(&header)
	}
}

func BenchmarkFilterMyHeader(b *testing.B) {
	b.StopTimer()
	var header http.Header = map[string][]string{
		"header-1": []string{"value 1.1", "value 1.2"},
		"header-2": []string{"value 2.1"},
		"header-3": []string{"value 3.1", "value 3.2", "value 3.3"},
		"header-4": []string{"value 1.1", "value 1.2"},
		"header-5": []string{"value 2.1"},
		"header-6": []string{"value 3.1", "value 3.2", "value 3.3", "value 3.1", "value 3.2", "value 3.3"},
	}
	filter := []string{"header-2"}
	h := headers.MyHeader{}
	h.Copy(&header)
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		h.Filter(filter)
	}
}
