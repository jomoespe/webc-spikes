package headers_test

import (
	"net/http"
	"testing"

	"github.com/jomoespe/webc-spikes/pkg/headers"
)

func TestCanFilterHeaders(t *testing.T) {
	headerWithThreeHeaders := map[string][]string{
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
		{"Header with three headers and no filter", headerWithThreeHeaders, []string{}, 3},
		{"Header with three headers and removing first", headerWithThreeHeaders, []string{"header-1"}, 2},
		{"Header with three headers and removing middle", headerWithThreeHeaders, []string{"header-2"}, 2},
		{"Header with three headers and removing last", headerWithThreeHeaders, []string{"header-3"}, 2},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			h := headers.Filter(tt.header, tt.filter)
			if len(h) != tt.want {
				t.Fatalf("Got: %d, want: %d", len(h), tt.want)
			}
		})
	}
}

func BenchmarkFilter(b *testing.B) {
	b.StopTimer()

	for i := 0; i < 100; i++ {
	}

	header := map[string][]string{
		"header-1": []string{"value 1.1", "value 1.2"},
		"header-2": []string{"value 2.1"},
		"header-3": []string{"value 3.1", "value 3.2", "value 3.3"},
	}
	filter := []string{"header-2"}
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		headers.Filter(header, filter)
	}
}
