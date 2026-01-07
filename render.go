package gml

import (
	"bytes"
	"context"
	"io"
)

type gmlElement struct {
	children   []GmlElement // 24 bytes
	attributes []Attr       // 24 bytes
	tag        uint8        // 1 byte
	void       bool         // 1 byte
}

var leftBracket = []byte("<")
var rightBracket = []byte(">")
var closingTag = []byte("</")
var space = []byte(" ")

func newgmlElement(tag uint8, void bool) *gmlElement {
	elm := gmlElement{}
	elm.tag = tag
	elm.void = void
	return &elm
}

func (elm *gmlElement) Attributes(attributes ...Attr) GmlElement {
	if l := len(attributes); l != 0 {
		elm.attributes = append(elm.attributes, attributes...)
	}

	return elm
}

func (elm *gmlElement) Children(children ...GmlElement) GmlElement {
	if elm.void {
		return elm
	}

	if l := len(children); l != 0 {
		elm.children = append(elm.children, children...)
	}

	return elm
}

func (d *gmlElement) renderInternal(ctx context.Context, w io.Writer, bestEffort bool) error {
	if d == nil {
		return nil
	}

	// open tag
	if _, err := w.Write(leftBracket); err != nil {
		return err
	}

	tag := elementNames[d.tag]

	if _, err := w.Write(tag); err != nil {
		return err
	}

	// attributes
	for _, attr := range d.attributes {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if _, err := w.Write(space); err != nil {
				return err
			}

			if err := attr.Render(w); err != nil {
				return err
			}
		}

	}

	if _, err := w.Write(rightBracket); err != nil {
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
			if child == nil {
				continue
			}

			if err := child.Render(ctx, w); err != nil {
				if bestEffort {
					continue
				} else {
					return err
				}
			}
		}
	}

	// closing tag
	if _, err := w.Write(closingTag); err != nil {
		return err
	}

	if _, err := w.Write(tag); err != nil {
		return err
	}

	if _, err := w.Write(rightBracket); err != nil {
		return err
	}

	return nil
}

func (d *gmlElement) RenderHtml(ctx context.Context) (html []byte, err error) {
	if d == nil {
		return
	}

	var builder bytes.Buffer
	err = d.Render(ctx, &builder)
	if err != nil {
		return
	}

	return builder.Bytes(), nil
}

func (d *gmlElement) Render(ctx context.Context, w io.Writer) error {
	if d == nil {
		return nil
	}

	return d.renderInternal(ctx, w, false)
}

func (d *gmlElement) RenderBestEffort(ctx context.Context, w io.Writer) error {
	if d == nil {
		return nil
	}

	return d.renderInternal(ctx, w, true)
}
