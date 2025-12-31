package gml

import (
	"context"
	"io"
)

type HtmlElement interface {
	Attributes(attributes ...Attr) HtmlElement
	Children(children ...HtmlElement) HtmlElement
	RenderHtml(ctx context.Context) string
	Render(ctx context.Context, w io.Writer) error
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
