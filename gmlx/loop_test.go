package gmlx

import (
	"fmt"
	"testing"

	"github.com/501urchin/gml"
)

func TestMap(t *testing.T) {
	items := []string{"hello", "wrld", "gml"}

	t.Run("Map without keys", func(t *testing.T) {
		elm := Map(items, func(idx int, item string) gml.GmlElement {
			return gml.P().Children(gml.Content(item))
		})

		html := elm.RenderHtml(t.Context())

		exp := ""
		for _, item := range items {
			exp += fmt.Sprintf("<p>%s</p>", item)
		}

		if html != exp {
			t.Fatalf("returned html doesn't match expected: got: %v -> expected: %v", html, exp)
		}
	})

	t.Run("MapKeyed with keys", func(t *testing.T) {
		elmKeyed := MapKeyed(items, func(idx int, item string) gml.GmlElement {
			return gml.P().Children(gml.Content(item))
		})

		htmlKeyed := elmKeyed.RenderHtml(t.Context())

		expKeyed := ""
		for i, item := range items {
			expKeyed += fmt.Sprintf("<p key=\"%d\">%s</p>", i, item)
		}

		if htmlKeyed != expKeyed {
			t.Fatalf("returned html doesn't match expected: got: %v -> expected: %v", htmlKeyed, expKeyed)
		}
	})
}
