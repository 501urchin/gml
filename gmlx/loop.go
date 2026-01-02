package gmlx

import (
	"bytes"
	"context"
	"io"
	"strconv"

	"github.com/501urchin/gml"
)

type mapComponent[S any] struct {
	items []S
	iter  func(index int, item S) gml.GmlElement
	keyed bool
}

func Map[S any](items []S, fn func(index int, item S) gml.GmlElement) *mapComponent[S] {
	return &mapComponent[S]{items: items, iter: fn, keyed: false}
}

func MapKeyed[S any](items []S, fn func(index int, item S) gml.GmlElement) *mapComponent[S] {
	return &mapComponent[S]{items: items, iter: fn, keyed: true}
}

func (m *mapComponent[S]) Attributes(_ ...gml.Attr) gml.GmlElement {
	return m
}

func (m *mapComponent[S]) Children(_ ...gml.GmlElement) gml.GmlElement {
	return m
}

func (m *mapComponent[S]) RenderHtml(ctx context.Context) (html []byte, err error) {
	if m == nil {
		return
	}

	var buf bytes.Buffer
	err = m.Render(ctx, &buf)
	if err != nil {
		return
	}

	return buf.Bytes(), nil
}

func (m *mapComponent[S]) Render(ctx context.Context, w io.Writer) error {
	if m == nil {
		return nil
	}

	for idx, item := range m.items {
		res := m.iter(idx, item)
		if m.keyed {
			res.Attributes(gml.Attribute("key", strconv.Itoa(idx)))
		}

		err := res.Render(ctx, w)
		if err != nil {
			return err
		}
	}

	return nil
}
func (m *mapComponent[S]) RenderBestEffort(ctx context.Context, w io.Writer) error {
	if m == nil {
		return nil
	}

	for idx, item := range m.items {
		res := m.iter(idx, item)
		if m.keyed {
			res.Attributes(gml.Attribute("key", strconv.Itoa(idx)))
		}

		err := res.Render(ctx, w)
		if err != nil {
			continue
		}
	}

	return nil
}
