package html

import (
	"strings"
)

type HtmlElement interface {
	Render() string
}

type Element struct {
	void          bool
	tag           string
	attributes    []Attr
	attrKeysLen   int
	attrValuesLen int
	children      []HtmlElement
}

func (d Element) Render() string {
	builder := strings.Builder{}
	l := len(d.attributes)
	builder.Grow(
		6 + (2 * len(d.tag)) + // opening and closing tag
			(l) + // space rune between attr
			d.attrKeysLen +
			d.attrValuesLen,
	)

	builder.WriteRune('<')
	builder.WriteString(d.tag)

	// attributes - 10% allocs
	for _, attr := range d.attributes {
		builder.WriteRune(' ')
		builder.WriteString(attr.Render())
	}
	builder.WriteString(">")

	if d.void {
		return builder.String()
	}
	// chldren - 42% allocs
	for _, child := range d.children {
		builder.WriteString(child.Render())
	}

	// closing tag
	builder.WriteString("</")
	builder.WriteString(d.tag)
	builder.WriteRune('>')

	return builder.String()
}

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
