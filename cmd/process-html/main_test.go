// Copyright 2018 Jomoespe. All rights reserved.
// Use of this source code is governed by a WTFPL-style
// license that can be found in the LICENSE file.

package main_test

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"testing"
)

type composer func(*sync.WaitGroup, *html.Node, predicate, visitor)

type predicate func(*html.Node) bool

type visitor func(*sync.WaitGroup, *html.Node)

func TestProcessHTMLTasgs(t *testing.T) {
	tt := [...]struct {
		filename  string
		predicate predicate
		total     int
	}{
		{"../../testdata/process-html/example-1.html", tag(atom.Section), 3},
		{"../../testdata/process-html/example-1.html", tag(atom.P), 2},
		{"../../testdata/process-html/example-1.html", attribute("data-src"), 3},
	}
	for _, test := range tt {
		b, _ := os.Open(test.filename)
		defer b.Close()
		doc, err := html.Parse(b)
		if err != nil {
			t.Fatalf("cannot parse html test file %s", test.filename)
		}

		tags := 0
		var f func(*html.Node, predicate)
		f = func(n *html.Node, predicate predicate) {
			if predicate(n) {
				tags++
			} else {
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					f(c, predicate)
				}
			}
		}
		f(doc, test.predicate)

		if test.total != tags {
			t.Fatalf("expected: %d, got: %v", test.total, tags)
		}
	}
}

func ExampleVisitNodes() {
	filename := "../../testdata/process-html/example-1.html"

	b, _ := os.Open(filename)
	defer b.Close()
	doc, err := html.Parse(b)
	if err != nil {
		fmt.Printf("cannot parse html test file %s", filename)
	}

	composer := concurrentComposer()
	var wg sync.WaitGroup
	wg.Add(1)
	go composer(&wg, doc, tag(atom.Section), printNodeData())
	wg.Wait()

	// Output:
	// section
	// section
	// section
}

// printNodeData returns visitor that print the node.Data
func printNodeData() visitor {
	rand.Seed(time.Now().UnixNano())
	return func(wg *sync.WaitGroup, n *html.Node) {
		defer wg.Done()
		fmt.Println(n.Data)
	}
}

// tag returns a predicate function that cheks if a node is for a tag type
func tag(a atom.Atom) predicate {
	return func(n *html.Node) bool {
		return n.Type == html.ElementNode && n.DataAtom == a
	}
}

// attribute returns a predicate function that cheks if a node contains an attribute
func attribute(key string) predicate {
	contains := func(attrs []html.Attribute, key string) bool {
		for _, a := range attrs {
			if a.Key == key {
				return true
			}
		}
		return false
	}
	return func(n *html.Node) bool {
		return n.Type == html.ElementNode && contains(n.Attr, key)
	}
}

//
func concurrentComposer() composer {
	var c composer
	c = func(wg *sync.WaitGroup, node *html.Node, p predicate, v visitor) {
		defer wg.Done()
		if p(node) {
			wg.Add(1)
			go v(wg, node)
		} else {
			for child := node.FirstChild; child != nil; child = child.NextSibling {
				wg.Add(1)
				go c(wg, child, p, v)
			}
		}
	}
	return c
}

// TODO Add DOM content replacement test
