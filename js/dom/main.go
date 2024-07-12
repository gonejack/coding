//go:build js

package main

import (
	_ "embed"

	"golang.org/x/exp/slices"
)

func main() {
	e := dom.GetWindow().Document().QuerySelector(".some-element")
	e.SetAttribute("class", "abc")
	e.GetAttribute("class")
	e.ParentNode()
	dom.WrapHTMLElement(e.Underlying())
}

func prepend[T any](arr []T, items ...T) (brr []T) {
	return slices.Insert(arr, 0, items...)
}
