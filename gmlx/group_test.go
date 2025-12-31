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

	html := elm.RenderHtml(t.Context())

	if exp := elm1.RenderHtml(t.Context()) + elm2.RenderHtml(t.Context()); html != exp {
		t.Errorf("failled to render group: got %q but expected %q", html, exp)
	}
}
