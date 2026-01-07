package gml

import (
	"io"

	"github.com/501urchin/gml/internal"
)

type Attr struct {
	Key   string
	Value string
}

func Attribute(key string, value ...string) Attr {
	a := Attr{Key: key}

	if len(value) != 0 {
		a.Value = value[0]
	}

	return a
}

var attrBegin = []byte(`="`)
var attrEnd = []byte(`"`)

func (a Attr) Render(w io.Writer) (err error) {
	if len(a.Key) == 0 {
		return
	}

	_, err = w.Write(internal.StringToBytes(a.Key))
	if err != nil {
		return
	}

	if a.Value != "" {
		_, err = w.Write(attrBegin)
		if err != nil {
			return
		}

		_, err = w.Write(internal.StringToBytes(a.Value))
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
