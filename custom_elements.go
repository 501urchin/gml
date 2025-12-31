package gml

import (
	"io"
	"strings"
)

// text content
type text struct {
	htmlString string
}

func Text(t string) text {
	return text{htmlString: t}
}

func (t text) RenderHtml() string {
	return t.htmlString
}

func (t text) Render(w io.Writer) error {
	_, err := w.Write([]byte(t.htmlString))
	return err
}


// for loop
type forComponent struct {
	elements []Element
}

func For(fn func() []Element) forComponent {
	return forComponent{elements: fn()}
}

func (t forComponent) RenderHtml() string {
	var buf strings.Builder

	for _, elm := range t.elements {
		_, err := buf.WriteString(elm.RenderHtml())
		if err != nil {
			return ""
		}
	}

	return buf.String()
}

func (t forComponent) Render(w io.Writer) error {
	for _, elm := range t.elements {
		err := elm.Render(w)
		if err != nil {
			return err
		}
	}

	return nil
}
