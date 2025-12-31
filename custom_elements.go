package gml

import (
	"bytes"
	"context"
	"io"
)

// text content
type content struct {
	htmlString string
}

func Content(t string) content {
	return content{htmlString: t}
}

func (t content) Attributes(_ ...Attr) HtmlElement {
	return nil
}
func (t content) Children(_ ...HtmlElement) HtmlElement {
	return nil
}

func (t content) RenderHtml(ctx context.Context) string {
	return t.htmlString
}

func (t content) Render(_ context.Context, w io.Writer) error {
	_, err := w.Write([]byte(t.htmlString))
	return err
}

// for loop
type forComponent struct {
	elements []HtmlElement
}

func For(fn func() []HtmlElement) *forComponent {
	return &forComponent{elements: fn()}
}

func (t *forComponent) Attributes(_ ...Attr) HtmlElement {
	return nil
}
func (t *forComponent) Children(_ ...HtmlElement) HtmlElement {
	return nil
}

func (t *forComponent) RenderHtml(ctx context.Context) string {
	var buf bytes.Buffer
	for _, elm := range t.elements {
		err := elm.Render(ctx, &buf)
		if err != nil {
			return ""
		}
	}

	return buf.String()
}

func (t *forComponent) Render(ctx context.Context, w io.Writer) error {
	for _, elm := range t.elements {
		err := elm.Render(ctx, w)
		if err != nil {
			return err
		}
	}

	return nil
}
