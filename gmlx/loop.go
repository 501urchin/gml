package gmlx

import (
	"bytes"
	"context"
	"io"

	"github.com/501urchin/gml"
)

type mapComponent[S any] struct {
	items []S
	iter  func(index int, item S) gml.GmlElement
}

func Map[S any](items []S, fn func(index int, item S) gml.GmlElement) mapComponent[S] {
	return mapComponent[S]{items: items, iter: fn}
}

func (m mapComponent[S]) Attributes(_ ...gml.Attr) gml.GmlElement {
	return m
}

func (m mapComponent[S]) Children(_ ...gml.GmlElement) gml.GmlElement {
	return m
}

func (m mapComponent[S]) RenderHtml(ctx context.Context) (html []byte, err error) {
	var buf bytes.Buffer
	err = m.Render(ctx, &buf)
	if err != nil {
		return
	}

	return buf.Bytes(), nil
}

func (m mapComponent[S]) Render(ctx context.Context, w io.Writer) error {
	for idx, item := range m.items {
		res := m.iter(idx, item)
		err := res.Render(ctx, w)
		if err != nil {
			return err
		}
	}

	return nil
}
func (m mapComponent[S]) RenderBestEffort(ctx context.Context, w io.Writer) error {
	for idx, item := range m.items {
		res := m.iter(idx, item)
		err := res.Render(ctx, w)
		if err != nil {
			continue
		}
	}

	return nil
}
