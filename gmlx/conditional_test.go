package gmlx

import (
	"bytes"
	"testing"

	"github.com/501urchin/gml"
)

func TestConditional(t *testing.T) {
	elm := gml.P().Children(gml.Content("hello"))

	t.Run("true", func(t *testing.T) {
		condElm := If(true, func() gml.GmlElement { return elm })
		h1, err := condElm.RenderHtml(t.Context())
		if err != nil {
			t.Fatal(err)
		}
		h2, err := elm.RenderHtml(t.Context())
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(h1, h2) {
			t.Error("faield to retun correct element")
		}
	})
	t.Run("false", func(t *testing.T) {
		condElm := If(false, func() gml.GmlElement { return elm })
		h1, err := condElm.RenderHtml(t.Context())
		if err != nil {
			t.Fatal(err)
		}
		if len(h1) != 0 {
			t.Error("faield to retun correct element")
		}
	})
}
