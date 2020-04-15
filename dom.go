// Package dom provides QuerySelector related functions to be used
// with https://godoc.org/golang.org/x/net/html .
package dom

import (
	"golang.org/x/net/html"
)

// QuerySelector provides basic functionality of the JavaScript
// document|node.querySelector() method, finding the first HTML Element that
// matches the specified CSS selector in the tree starting with Node n.
//
// If the input selector string is invalid an error is returned.
func QuerySelector(n *html.Node, selector string) (*html.Node, error) {
	return nil, nil
}
