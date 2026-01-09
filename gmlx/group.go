package gmlx

import (
	"bytes"
	"context"
	"io"

	"github.com/501urchin/gml"
)

type group struct {
	gmlElements []gml.GmlElement
}

func Group(gmlElements ...gml.GmlElement) (g *group) {
	g = &group{}
	g.gmlElements = append(g.gmlElements, gmlElements...)
	return
}

func (t *group) Attributes(_ ...gml.Attr) gml.GmlElement {
	return t
}

func (t *group) Children(_ ...gml.GmlElement) gml.GmlElement {
	return t
}

func (t *group) RenderHtml(ctx context.Context) (html []byte, err error) {
	if t == nil {
		return
	}

	var buf bytes.Buffer
	for _, elm := range t.gmlElements {
		err := elm.Render(ctx, &buf)
		if err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

func (t *group) Render(ctx context.Context, w io.Writer) error {
	if t == nil {
		return nil
	}

	for _, elm := range t.gmlElements {
		err := elm.Render(ctx, w)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *group) RenderBestEffort(ctx context.Context, w io.Writer) error {
	if t == nil {
		return nil
	}

	for _, elm := range t.gmlElements {
		err := elm.Render(ctx, w)
		if err != nil {
			continue
		}
	}

	return nil
}
