package gml

import (
	"context"
	"io"
)

// text content
type content struct {
	htmlString string
}

// Content creates a content struct representing the inner content of an HTML gmlElement.
// It can also be used to render raw HTML strings directly.
func Content(t string) content {
	return content{htmlString: t}
}

func (t content) Attributes(_ ...Attr) GmlElement {
	return t
}
func (t content) Children(_ ...GmlElement) GmlElement {
	return t
}

func (t content) RenderHtml(ctx context.Context) string {
	return t.htmlString
}

func (t content) Render(_ context.Context, w io.Writer) error {
	_, err := w.Write([]byte(t.htmlString))
	return err
}

func (t content) RenderBestEffort(_ context.Context, w io.Writer) error {
	_, err := w.Write([]byte(t.htmlString))
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

func (empty) RenderHtml(ctx context.Context) string {
	return ""
}
func (empty) Render(ctx context.Context, w io.Writer) error {
	return nil
}
func (empty) RenderBestEffort(ctx context.Context, w io.Writer) error {
	return nil
}
