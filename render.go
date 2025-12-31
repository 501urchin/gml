package gml

import (
	"io"
	"strings"
)

func newElement(tag string, void bool, attributes []Attr, children ...HtmlElement) Element {
	elm := Element{}
	elm.tag = tag
	elm.void = void

	if l := len(attributes); l != 0 {
		elm.attributes = make([]Attr, 0, l)

		for _, attr := range attributes {
			elm.attributes = append(elm.attributes, attr)
			elm.attrKeysLen += len(attr.Key)
			elm.attrValuesLen += len(attr.Value) + 3
		}
	}

	if l := len(children); l != 0 {
		elm.children = make([]HtmlElement, 0, l)
		elm.children = append(elm.children, children...)
	}

	return elm
}

func (d Element) RenderHtml() string {
	var builder strings.Builder

	if d.tag == "" {
		return ""
	}

	l := len(d.attributes)
	builder.Grow(
		6 + (2 * len(d.tag)) + // opening and closing tag
			l + // space rune between attr
			d.attrKeysLen +
			d.attrValuesLen,
	)

	builder.WriteRune('<')
	builder.WriteString(d.tag)

	// render child attributes
	for _, attr := range d.attributes {
		builder.WriteRune(' ')
		builder.WriteString(attr.RenderHtml())
	}
	builder.WriteString(">")

	// if the elm is void (self closing) return early since it wont have children or a closing tag
	if d.void {
		return builder.String()
	}

	// render child elements
	for _, child := range d.children {
		builder.WriteString(child.RenderHtml())
	}

	// closing tag
	builder.WriteString("</")
	builder.WriteString(d.tag)
	builder.WriteRune('>')

	return builder.String()
}

func (d Element) Render(w io.Writer) error {
	if d.tag == "" {
		return nil
	}

	// open tag
	if _, err := w.Write([]byte("<")); err != nil {
		return err
	}
	if _, err := w.Write([]byte(d.tag)); err != nil {
		return err
	}

	// attributes
	for _, attr := range d.attributes {
		if _, err := w.Write([]byte(" ")); err != nil {
			return err
		}
		if _, err := w.Write([]byte(attr.RenderHtml())); err != nil {
			return err
		}
	}

	if _, err := w.Write([]byte(">")); err != nil {
		return err
	}

	if d.void {
		return nil
	}

	// children
	for _, child := range d.children {
		if err := child.Render(w); err != nil {
			return err
		}
	}

	// closing tag
	if _, err := w.Write([]byte("</")); err != nil {
		return err
	}
	if _, err := w.Write([]byte(d.tag)); err != nil {
		return err
	}
	if _, err := w.Write([]byte(">")); err != nil {
		return err
	}

	return nil
}
