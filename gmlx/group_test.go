package gmlx

import (
	"testing"

	"github.com/501urchin/gml"
)

func TestGroup(t *testing.T) {
	elm1 := gml.P().Children(gml.Content("hello"))
	elm2 := gml.P().Children(gml.Content("hi"))
	elm := Group(
		elm1,
		elm2,
	)

	html, err := elm.RenderHtml(t.Context())
	if err != nil {
		t.Fatal(err)
	}

	h1, err := elm1.RenderHtml(t.Context())
	if err != nil {
		t.Fatal(err)
	}

	h2, err := elm2.RenderHtml(t.Context())
	if err != nil {
		t.Fatal(err)
	}

	if exp := string(h1) + string(h2); string(html) != exp {
		t.Errorf("failled to render group: got %q but expected %q", html, exp)
	}
}

func BenchmarkGroup(b *testing.B) {
	elm1 := gml.P().Children(gml.Content("hello"))
	elm2 := gml.P().Children(gml.Content("hi"))
	elm := Group(
		elm1,
		elm2,
	)

	ctx := b.Context()

	for b.Loop() {
		elm.RenderHtml(ctx)
	}

}
