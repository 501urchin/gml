package gml

import (
	"html"
	"strings"
)

func (a Attr) RenderHtml() string {
	var b strings.Builder
	b.Grow(len(a.Key) + len(a.Value) + 3)
	b.WriteString(a.Key)

	if a.Value != "" {
		b.WriteByte('=')
		b.WriteByte('"')
		if a.raw {
			b.WriteString(a.Value)
		} else {
			b.WriteString(html.EscapeString(a.Value))
		}
		b.WriteByte('"')
	}
	return b.String()
}

func Attribute(key string, value ...string) Attr {
	a := Attr{Key: key, raw: false}

	if len(value) != 0 {
		a.Value = value[0]
	}

	return a
}

func AttributeNoEscape(key string, value ...string) Attr {
	a := Attr{Key: key, raw: true}

	if len(value) != 0 {
		a.Value = value[0]
	}

	return a
}
