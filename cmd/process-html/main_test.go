package main_test

import (
	"os"

	"golang.org/x/net/html"

	"testing"
)

func TestProcessHTMLTasgs(t *testing.T) {
	tt := [...]struct {
		filename string
		tag      string
		total    int
	}{
		{"../../testdata/process-html/example-1.html", "section", 3},
		{"../../testdata/process-html/example-1.html", "p", 2},
	}

	for _, test := range tt {
		b, _ := os.Open(test.filename)
		doc, err := html.Parse(b)
		if err != nil {
			t.Fatalf("cannot parse testfile %s", test.filename)
		}
		// TODO here to process the HTML looking for tag type
		//		t.Errorf("Type: %v", doc.Type)

		tags := 0

		var f func(*html.Node)
		f = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == test.tag {
				tags++
			} else {
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					f(c)
				}
			}
		}
		f(doc)

		if test.total != tags {
			t.Fatalf("expected: %d, got: %v", test.total, tags)
		}
	}
}

// TODO Add another test to look for 'data-src' attribute
