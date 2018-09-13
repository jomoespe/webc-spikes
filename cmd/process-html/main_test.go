package main_test

import (
	"fmt"
	"os"
	"sync"

	"golang.org/x/net/html"

	"testing"
)

type predicate func(*html.Node) bool
type visitor func(*sync.WaitGroup, *html.Node)

func TestProcessHTMLTasgs(t *testing.T) {
	tt := [...]struct {
		filename  string
		predicate predicate
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
		f(test.predicate, doc)

		if test.total != tags {
			t.Fatalf("expected: %d, got: %v", test.total, tags)
		}
	}
}

func ExampleVisitNodes() {
	filename := "../../testdata/process-html/example-1.html"
	isSection := tag("section")
	printNode := printNodeData()

	b, _ := os.Open(filename)
	doc, err := html.Parse(b)
	if err != nil {
		fmt.Printf("cannot parse testfile %s", filename)
	}

	var f func(*sync.WaitGroup, *html.Node, predicate, visitor)
	f = func(wg *sync.WaitGroup, node *html.Node, p predicate, v visitor) {
		defer wg.Done()
		if p(node) {
			wg.Add(1)
			go v(wg, node)
		} else {
			for child := node.FirstChild; child != nil; child = child.NextSibling {
				wg.Add(1)
				f(wg, child, p, v)
			}
		}
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go f(&wg, doc, isSection, printNode)
	wg.Wait()
	// Output:
	// section
	// section
	// section
}

// printNodeData returns a node visitor that print the node.Data
func printNodeData() visitor {
	return func(wg *sync.WaitGroup, n *html.Node) {
		defer wg.Done()
		fmt.Println(n.Data)
	}
}

// tag returns a predicate function that cheks if a node is for a tag type
func tag(tagType string) predicate {
	return func(n *html.Node) bool {
		return n.Type == html.ElementNode && n.Data == tagType
	}
}

// tag returns a predicate function that cheks if a node contains a attribute
func attribute(attr string) predicate {
	contains := func(attr string, attrs []html.Attribute) bool {
		for _, a := range attrs {
			if a.Key == attr {
				return true
			}
		}
		return false
	}
	return func(n *html.Node) bool {
		return n.Type == html.ElementNode && contains(attr, n.Attr)
	}
}

// TODO Add DOM content replacement test
