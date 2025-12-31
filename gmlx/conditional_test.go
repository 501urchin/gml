package gmlx

import (
	"testing"

	"github.com/501urchin/gml"
)

func TestConditional(t *testing.T) {
	elm := gml.P().Children(gml.Content("hello"))

	t.Run("true", func(t *testing.T) {
		condElm := If(true, func() gml.GmlElement { return elm })
		if condElm.RenderHtml(t.Context()) != elm.RenderHtml(t.Context()) {
			t.Error("faield to retun correct element")
		}
	})
	t.Run("false", func(t *testing.T) {
		condElm := If(false, func() gml.GmlElement { return elm })
		if condElm.RenderHtml(t.Context()) != "" {
			t.Error("faield to retun correct element")
		}
	})
}
