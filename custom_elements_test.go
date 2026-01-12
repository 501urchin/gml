package gml

import (
	"testing"
)

func TestContent(t *testing.T) {
	t.Run("type handling", func(t *testing.T) {
		ctx := t.Context()

		cases := []struct {
			input    any
			expected string
		}{
			{"hello", "hello"},
			{[]byte("world"), "world"},
			{true, "true"},
			{42, "42"},
			{int8(-8), "-8"},
			{int16(16), "16"},
			{int32(-32), "-32"},
			{int64(64), "64"},
			{uint(100), "100"},
			{uint8(8), "8"},
			{uint16(16), "16"},
			{uint32(32), "32"},
			{uint64(64), "64"},
			{nil, ""},
			{3.14, "3.14"},
			{float32(2.71), "2.71"},
			{struct{ A int }{A: 1}, `{"A":1}`},
		}

		for _, c := range cases {
			got, _ := Content(c.input).RenderHtml(ctx)
			if string(got) != c.expected {
				t.Errorf("Content(%v) = %q; want %q", c.input, got, c.expected)
			}
		}
	})

	t.Run("variable inputs", func(t *testing.T) {
		ctx := t.Context()
		expect := "hello12"

		r, _ := Content("hello", 12).RenderHtml(ctx)

		if string(r) != expect {
			t.Errorf("failed to render parameters of diffirent types: expected %q but got %q", expect, r)
		}
	})
}

func BenchmarkContent(b *testing.B) {
	for b.Loop() {
		Content("hello", 12, 13, 14, false, 45.3).RenderHtml(b.Context())
	}
}
