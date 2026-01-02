package gml

import (
	"io"
)

type Attr struct {
	Key   []byte
	Value []byte
}

type AttributeIface interface {
	Render(w io.Writer) (err error)
}

var attrBegin = []byte(`="`)
var attrEnd = []byte(`"`)

func (a Attr) Render(w io.Writer) (err error) {
	if len(a.Key) == 0 {
		return
	}
	
	_, err = w.Write(a.Key)
	if err != nil {
		return
	}

	if len(a.Value) != 0 {
		_, err = w.Write(attrBegin)
		if err != nil {
			return
		}

		_, err = w.Write(a.Value)
		if err != nil {
			return
		}

		_, err = w.Write(attrEnd)
		if err != nil {
			return
		}
	}

	return nil
}

// escapes attributes by default
func Attribute(key string, value ...string) Attr {
	a := Attr{Key: []byte(key)}

	if len(value) != 0 {
		a.Value = []byte(value[0])
	}

	return a
}
