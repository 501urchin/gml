package gml

import (
	"context"
	"io"
	"strings"

	"github.com/501urchin/gml/pkg"
)

// text content
type content struct {
	html string
}

// Content creates a content struct representing the inner content of an HTML GmlElement.
// It can also be used to render raw HTML values directly.
// Content accepts a variadic list of values; all arguments are concatenated in order
// without adding any separators or whitespace.
//
// For example, Content("hello", 12) renders as: hello12
func Content(t ...any) content {
	var builder strings.Builder

	for _, t := range t {
		builder.WriteString(pkg.ConvertToString(t))
	}

	return content{html: builder.String()}
}

// function is used to render raw html
func Raw(t string) content {
	return content{html: t}
}

func (t content) Attributes(_ ...Attr) GmlElement {
	return t
}
func (t content) Children(_ ...GmlElement) GmlElement {
	return t
}

func (t content) RenderHtml(ctx context.Context) (html []byte, err error) {
	return pkg.StringToBytes(t.html), nil
}

func (t content) Render(_ context.Context, w io.Writer) error {
	if len(t.html) == 0 {
		return nil
	}
	_, err := w.Write(pkg.StringToBytes(t.html))
	return err
}

func (t content) RenderBestEffort(_ context.Context, w io.Writer) error {
	if len(t.html) == 0 {
		return nil
	}
	_, err := w.Write(pkg.StringToBytes(t.html))
	return err
}

// empty element
type empty struct{}

func Empty() empty {
	return empty{}
}
func (e empty) Attributes(attributes ...Attr) GmlElement {
	return e
}
func (e empty) Children(children ...GmlElement) GmlElement {
	return e
}

var emptySlice = []byte{}

func (empty) RenderHtml(ctx context.Context) (html []byte, err error) {
	return emptySlice, nil
}
func (empty) Render(ctx context.Context, w io.Writer) error {
	return nil
}
func (empty) RenderBestEffort(ctx context.Context, w io.Writer) error {
	return nil
}
