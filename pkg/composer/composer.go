package composer

import (
	"fmt"
	"io"
	"sync"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type predicate func(*ComposableNode) bool

type visitor func(*sync.WaitGroup, *ComposableNode)

type composer func(*sync.WaitGroup, *ComposableNode, predicate, visitor)

type ComposableNode struct {
	html.Node
}

func Parse(r io.Reader) (c *ComposableNode, err error) {
	if n, err := html.Parse(r); err == nil {
		c = &ComposableNode{*n}
	}
	return
}

func (n *ComposableNode) Compose() {
	composer := concurrentComposer()
	predicate := tag(atom.Section)
	visitor := printNodeData()
	var wg sync.WaitGroup
	wg.Add(1)
	go composer(&wg, n, predicate, visitor)
	wg.Wait()
}

func (n *ComposableNode) IsTag(a atom.Atom) bool {
	return n.Type == html.ElementNode && n.DataAtom == a
}

// attribute returns a predicate function that cheks if a node contains an attribute
func (n *ComposableNode) hasAttr(key string) bool {
	contains := func(attrs []html.Attribute, key string) bool {
		for _, a := range attrs {
			if a.Key == key {
				return true
			}
		}
		return false
	}
	return n.Type == html.ElementNode && contains(n.Attr, key)
}

func tag(a atom.Atom) predicate {
	return func(n *ComposableNode) bool {
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
	return func(n *ComposableNode) bool {
		return n.Type == html.ElementNode && contains(n.Attr, key)
	}
}

// printNodeData returns visitor that print the node.Data
func printNodeData() visitor {
	return func(wg *sync.WaitGroup, n *ComposableNode) {
		defer wg.Done()
		fmt.Println(n.Data)
	}
}

func concurrentComposer() composer {
	var c composer
	c = func(wg *sync.WaitGroup, node *ComposableNode, p predicate, v visitor) {
		defer wg.Done()
		if p(node) {
			wg.Add(1)
			go v(wg, node)
		} else {
			for child := node.FirstChild; child != nil; child = child.NextSibling {
				wg.Add(1)
				go c(wg, &ComposableNode{*child}, p, v)
			}
		}
	}
	return c
}
