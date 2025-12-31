package gml

import "io"

type HtmlElement interface {
	RenderHtml() string
	Render(w io.Writer) error
}

type Element struct {
	void          bool
	tag           string
	attributes    []Attr
	attrKeysLen   int
	attrValuesLen int
	children      []HtmlElement
}

type Attr struct {
	Key   string
	Value string
}

type HtmlAttribute interface {
	RenderHtml() string
}
