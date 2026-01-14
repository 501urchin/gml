package gmlx

import (
	"context"
	"testing"

	"github.com/501urchin/gml"
)

func TestCache(t *testing.T) {
	ctx := context.Background()

	dep := 1
	depFunc := func(dep int) gml.GmlElement { return gml.Div().Children(gml.Content(dep)) }

	// Call Cache first time, should populate cache
	nlm := Cache(ctx, dep, depFunc)
	html, err := nlm.RenderHtml(ctx)
	if err != nil {
		t.Fatalf("RenderHtml failed: %v", err)
	}

	// Verify cache has the stored value
	found := false
	for _, v := range cacheStore.store {
		if string(html) == v {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("failed to set cache")
	}

	// Call Cache again with same dependency, should hit cache
	nlm2 := Cache(ctx, dep, depFunc)
	html2, err := nlm2.RenderHtml(ctx)
	if err != nil {
		t.Fatalf("RenderHtml failed on cache hit: %v", err)
	}

	if string(html2) != string(html) {
		t.Fatal("cache hit returned different value")
	}

	// Change dependency, cache should be invalidated
	dep = 5
	nlm3 := Cache(ctx, dep, depFunc)
	html3, err := nlm3.RenderHtml(ctx)
	if err != nil {
		t.Fatalf("RenderHtml failed on new dependency: %v", err)
	}

	if string(html3) == string(html) {
		t.Fatalf("cache was not invalidated when dependency changed: %q - %q", html3, html)
	}
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
			Cache(ctx, dep, func(dep int) gml.GmlElement { return gml.Div().Children(gml.Content(dep)) }).RenderHtml(ctx)
		}
	})

}
