package gmlx

import (
	"fmt"
	"io"
	"testing"

	"github.com/501urchin/gml"
)

func TestMap(t *testing.T) {
	items := []string{"hello", "wrld", "gml"}

	t.Run("Map without keys", func(t *testing.T) {
		elm := Map(items, func(idx int, item string) gml.GmlElement {
			return gml.P().Children(gml.Content(item))
		})

		html, err := elm.RenderHtml(t.Context())
		if err != nil {
			t.Fatal(err)
		}

		exp := ""
		for _, item := range items {
			exp += fmt.Sprintf("<p>%s</p>", item)
		}

		if string(html) != exp {
			t.Fatalf("returned html doesn't match expected: got: %v -> expected: %v", html, exp)
		}
	})
}

func BenchmarkMap(b *testing.B) {
	items := []string{"hello", "wrld", "gml"}
	ctx := b.Context()
	for b.Loop() {
		Map(items, func(idx int, item string) gml.GmlElement {
			return gml.P().Children(gml.Content(item))
		}).Render(ctx, io.Discard)
	}

}
