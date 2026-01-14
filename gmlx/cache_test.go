package gmlx

import (
	"testing"

	"github.com/501urchin/gml"
)

func TestCache(t *testing.T) {
	dep := 1
	elm := gml.Div().Children(gml.Content(dep))

	elm = Cache(t.Context(), elm, dep)
	elm = Cache(t.Context(), elm, dep)
	elm = Cache(t.Context(), elm, dep)
	dep = 5
	elm = Cache(t.Context(), elm, dep)

}

func BenchmarkCache(b *testing.B) {
	dep := 1
	elm := gml.Div().Children(gml.Content(dep))

	b.Run("normal", func(b *testing.B) {
		ctx := b.Context()
		for b.Loop() {
			elm.RenderHtml(ctx)
		}
	})

	b.Run("cached", func(b *testing.B) {
		ctx := b.Context()
		for b.Loop() {
			Cache(ctx, elm, dep).RenderHtml(ctx)
		}
	})

}
