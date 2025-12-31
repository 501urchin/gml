package gml

import "io"

// text element
type text struct {
	htmlString string
}

func Text(t string) text {
	return text{htmlString: t}
}

func (t text) Render() string {
	return t.htmlString
}

func (t text) RenderStream(w io.Writer) error {
	_, err := w.Write([]byte(t.htmlString))
	return err
}

func Meta(attributes []Attr) Element {
	return newElement("meta", true, attributes)
}

func Link(attributes []Attr) Element {
	return newElement("link", true, attributes)
}
