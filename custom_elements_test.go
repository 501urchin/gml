package gml

import (
	"testing"
)

func TestContent(t *testing.T) {
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

	ctx := t.Context()
	for _, c := range cases {
		got, _ := Content(c.input).RenderHtml(ctx)
		if string(got) != c.expected {
			t.Errorf("Content(%v) = %q; want %q", c.input, got, c.expected)
		}
	}
}
