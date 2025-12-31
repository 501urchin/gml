package gml

import "io"

type HtmlElement interface {
	Render() string
	RenderStream(w io.Writer) error
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
	Render() string
}
