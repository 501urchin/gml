package gml

import (
	"testing"
)

func TestFor(t *testing.T) {
	tag := `<p>hello</p>`
	expected := ""
	elm := P().Children(Content("hello"))

	html := For(func() []HtmlElement {
		var buf []HtmlElement
		for range 4 {
			expected += tag
			buf = append(buf, elm)
		}
		return buf
	}).RenderHtml(t.Context())

	if html != expected {
		t.Fatalf("returned html doesnt match expected: got: %v -> expected: %v", html, expected)
	}
}
