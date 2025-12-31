package gml

import (
	"strings"
)

func (a Attr) RenderHtml() string {
	var b strings.Builder
	b.Grow(len(a.Key) + len(a.Value) + 3)
	b.WriteString(a.Key)

	if a.Value != "" {
		b.WriteByte('=')
		b.WriteByte('"')
		b.WriteString(a.Value)
		b.WriteByte('"')
	}
	return b.String()
}

func Attribute(key, value string) Attr {
	return Attr{Key: key, Value: value}
}
