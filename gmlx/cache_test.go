package gmlx

import (
	"context"
	"testing"

	"github.com/501urchin/gml"
)

func TestMemoize(t *testing.T) {
	ctx := context.Background()

	dep := 1
	depFunc := func(dep int) gml.GmlElement { return gml.Div().Children(gml.Content(dep)) }

	// Call Memoize first time, should populate Memoize
	nlm := Memoize(ctx, dep, depFunc)
	html, err := nlm.RenderHtml(ctx)
	if err != nil {
		t.Fatalf("RenderHtml failed: %v", err)
	}

	// Verify Memoize has the stored value
	found := false
	for _, v := range memoStore.store {
		if string(html) == v {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("failed to set Memoize")
	}

	// Call Memoize again with same dependency, should hit Memoize
	nlm2 := Memoize(ctx, dep, depFunc)
	html2, err := nlm2.RenderHtml(ctx)
	if err != nil {
		t.Fatalf("RenderHtml failed on Memoize hit: %v", err)
	}

	if string(html2) != string(html) {
		t.Fatal("Memoize hit returned different value")
	}

	// Change dependency, Memoize should be invalidated
	dep = 5
	nlm3 := Memoize(ctx, dep, depFunc)
	html3, err := nlm3.RenderHtml(ctx)
	if err != nil {
		t.Fatalf("RenderHtml failed on new dependency: %v", err)
	}

	if string(html3) == string(html) {
		t.Fatalf("Memoize was not invalidated when dependency changed: %q - %q", html3, html)
	}
}

func BenchmarkMemoize(b *testing.B) {
	dep := 1
	ctx := b.Context()
	for b.Loop() {
		Memoize(ctx, dep, func(dep int) gml.GmlElement { return gml.Div().Children(gml.Content(dep)) }).RenderHtml(ctx)
	}

}
