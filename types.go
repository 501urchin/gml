package gml

import (
	"context"
	"io"
)

type GmlElement interface {
	Attributes(attributes ...Attr) GmlElement
	Children(children ...GmlElement) GmlElement
	RenderHtml(ctx context.Context) string
	Render(ctx context.Context, w io.Writer) error
	RenderBestEffort(ctx context.Context, w io.Writer) error
}

type gmlElement struct {
	void          bool
	tag           string
	attributes    []Attr
	attrKeysLen   int
	attrValuesLen int
	children      []GmlElement
}

