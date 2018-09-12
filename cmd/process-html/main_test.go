package main_test

import (
	"os"

	"golang.org/x/net/html"

	"testing"
)

func TestProcessHTMLTasgs(t *testing.T) {
	tt := [...]struct {
		filename  string
		validator func(n *html.Node) bool
		total     int
	}{
		{"../../testdata/process-html/example-1.html", tag("section"), 3},
		{"../../testdata/process-html/example-1.html", tag("p"), 2},
		{"../../testdata/process-html/example-1.html", attribute("data-src"), 3},
	}

	for _, test := range tt {
		b, _ := os.Open(test.filename)
		doc, err := html.Parse(b)
		if err != nil {
			t.Fatalf("cannot parse testfile %s", test.filename)
		}
		tags := 0
		var f func(func(*html.Node) bool, *html.Node)
		f = func(predicate func(*html.Node) bool, n *html.Node) {
			if predicate(n) {
				tags++
			} else {
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					f(predicate, c)
				}
			}
		}
		f(test.validator, doc)

		if test.total != tags {
			t.Fatalf("expected: %d, got: %v", test.total, tags)
		}
	}
}

// tag returns a predicate function that cheks if a node is for a tag type
func tag(tagType string) func(*html.Node) bool {
	return func(n *html.Node) bool {
		return n.Type == html.ElementNode && n.Data == tagType
	}
}

// tag returns a predicate function that cheks if a node contains a attribute
func attribute(attr string) func(*html.Node) bool {
	containsAttr := func(attr string, attrs []html.Attribute) bool {
		for _, a := range attrs {
			if a.Key == attr {
				return true
			}
		}
		return false
	}
	return func(n *html.Node) bool {
		return n.Type == html.ElementNode && containsAttr(attr, n.Attr)
	}
}

// TODO Add DOM content replacement test
