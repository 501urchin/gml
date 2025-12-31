package gml

import (
	"bytes"
	"context"
	"io"
	"slices"
)

func newElement(tag string, void bool) *Element {
	elm := Element{}
	elm.tag = tag
	elm.void = void
	return &elm
}

func (elm *Element) Attributes(attributes ...Attr) HtmlElement {
	if l := len(attributes); l != 0 {
		elm.attributes = slices.Grow(elm.attributes, l)

		for _, attr := range attributes {
			elm.attributes = append(elm.attributes, attr)
			elm.attrKeysLen += len(attr.Key)
			elm.attrValuesLen += len(attr.Value) + 3
		}
	}

	return elm
}
func (elm *Element) Children(children ...HtmlElement) HtmlElement {
	if elm.void {
		return elm
	}

	if l := len(children); l != 0 {
		elm.children = slices.Grow(elm.children, l)
		elm.children = append(elm.children, children...)
	}

	return elm
}

func (d *Element) RenderHtml(ctx context.Context) string {
	var builder bytes.Buffer
	err := d.Render(ctx, &builder)
	if err != nil {
		return ""
	}

	return builder.String()
}

func (d *Element) Render(ctx context.Context, w io.Writer) error {
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
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if _, err := w.Write([]byte(" ")); err != nil {
				return err
			}
			if _, err := w.Write([]byte(attr.RenderHtml())); err != nil {
				return err
			}
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
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := child.Render(ctx, w); err != nil {
				return err
			}
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
