package gml

import (
	"bytes"
)

type Attr struct {
	Key   []byte
	Value []byte
}

func (a Attr) RenderHtml() []byte {
	var b bytes.Buffer
	b.Grow(len(a.Key) + len(a.Value) + 3)
	b.Write(a.Key)

	if len(a.Value) != 0 {
		b.WriteByte('=')
		b.WriteByte('"')
		b.Write(a.Value)
		b.WriteByte('"')
	}
	return b.Bytes()
}

// escapes attributes by default
func Attribute(key string, value ...string) Attr {
	a := Attr{Key: []byte(key)}

	if len(value) != 0 {
		a.Value = []byte(value[0])
	}

	return a
}
