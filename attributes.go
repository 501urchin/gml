package html

import "strings"

type Attr struct {
	Key   string
	Value string
}

func (a Attr) Render() string {
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

type HtmlAttribute interface {
	Render() string
}

func Attribute(key, value string) Attr {
	return Attr{Key: key, Value: value}
}
