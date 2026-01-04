package gmlx

import (
	"context"
	"io"

	"github.com/501urchin/gml"
)

type ifInternal struct {
	cond     bool
	ifFunc   func() gml.GmlElement
	elseFunc func() gml.GmlElement
}

func If(cond bool, fn func() gml.GmlElement) *ifInternal {
	return &ifInternal{cond: cond, ifFunc: fn}
}

func (i *ifInternal) Else(fn func() gml.GmlElement) *ifInternal {
	i.elseFunc = fn
	return i
}

func (i *ifInternal) Attributes(attributes ...gml.Attr) gml.GmlElement     { return i }
func (i *ifInternal) Children(attributes ...gml.GmlElement) gml.GmlElement { return i }
func (i *ifInternal) RenderBestEffort(ctx context.Context, w io.Writer) error {
	return i.Render(ctx, w)
}

func (i *ifInternal) Render(ctx context.Context, w io.Writer) error {
	if i.ifFunc == nil {
		return gml.Empty().Render(ctx, w)
	}

	if i.cond {
		return i.ifFunc().Render(ctx, w)
	}

	if i.elseFunc != nil && !i.cond {
		return i.elseFunc().Render(ctx, w)
	}

	return nil
}
func (i *ifInternal) RenderHtml(ctx context.Context) ([]byte, error) {
	if i.ifFunc == nil {
		return gml.Empty().RenderHtml(ctx)
	}

	if i.cond {
		return i.ifFunc().RenderHtml(ctx)
	}

	if i.elseFunc != nil && !i.cond {
		return i.elseFunc().RenderHtml(ctx)
	}

	return nil, nil
}
