package gml

import (
	"context"
	"io"
)

type GmlElement interface {
	// Attributes returns the receiver with the provided attributes appended.
	Attributes(attributes ...Attr) GmlElement

	// Children returns the receiver with the provided child elements appended.
	Children(children ...GmlElement) GmlElement

	// Render writes the rendered HTML representation of the element to w.
	// If rendering this element or any of its children returns an error,
	// rendering stops immediately and that error is returned.
	Render(ctx context.Context, w io.Writer) error

	// RenderBestEffort writes the rendered HTML representation of the element to w.
	// Unlike Render, this method continues rendering remaining children even if
	// one or more children return an error. The returned error, if non-nil,
	// represents at least one failure during rendering.
	RenderBestEffort(ctx context.Context, w io.Writer) error

	// RenderHtml renders the element and returns the resulting HTML as a byte slice.
	// This is intended for use cases that require raw HTML output rather than
	// streaming to an io.Writer.
	RenderHtml(ctx context.Context) ([]byte, error)
}
