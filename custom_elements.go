package gml

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/501urchin/gml/internal"
)

// text content
type content struct {
	html string
}

var null = ""

// Content creates a content struct representing the inner content of an HTML gmlElement.
// It can also be used to render raw HTML strings directly.
func Content(t any) content {
	var parsed string

	switch v := t.(type) {
	case string:
		parsed = v
	case fmt.Stringer:
		parsed = v.String()
	case nil:
		parsed = null
	case []byte:
		parsed = internal.BytesToString(v)
	case bool:
		parsed = strconv.FormatBool(v)
	case int:
		parsed = strconv.FormatInt(int64(v), 10)
	case int8:
		parsed = strconv.FormatInt(int64(v), 10)
	case int16:
		parsed = strconv.FormatInt(int64(v), 10)
	case int32:
		parsed = strconv.FormatInt(int64(v), 10)
	case int64:
		parsed = strconv.FormatInt(v, 10)
	case uint:
		parsed = strconv.FormatUint(uint64(v), 10)
	case uint8:
		parsed = strconv.FormatUint(uint64(v), 10)
	case uint16:
		parsed = strconv.FormatUint(uint64(v), 10)
	case uint32:
		parsed = strconv.FormatUint(uint64(v), 10)
	case uint64:
		parsed = strconv.FormatUint(v, 10)
	case float32:
		parsed = strconv.FormatFloat(float64(v), 'g', -1, 32)
	case float64:
		parsed = strconv.FormatFloat(v, 'g', -1, 64)
	default:
		// fall back to rendering it as json

		dt, _ := json.Marshal(v)
		parsed = internal.BytesToString(dt)
	}
	return content{html: parsed}
}

// function is used to render raw html
func Raw(t string) content {
	return content{html: t}
}

func (t content) Attributes(_ ...Attr) GmlElement {
	return t
}
func (t content) Children(_ ...GmlElement) GmlElement {
	return t
}

func (t content) RenderHtml(ctx context.Context) (html []byte, err error) {
	return internal.StringToBytes(t.html), nil
}

func (t content) Render(_ context.Context, w io.Writer) error {
	if len(t.html) == 0 {
		return nil
	}
	_, err := w.Write(internal.StringToBytes(t.html))
	return err
}

func (t content) RenderBestEffort(_ context.Context, w io.Writer) error {
	if len(t.html) == 0 {
		return nil
	}
	_, err := w.Write(internal.StringToBytes(t.html))
	return err
}

// empty element
type empty struct{}

func Empty() empty {
	return empty{}
}
func (e empty) Attributes(attributes ...Attr) GmlElement {
	return e
}
func (e empty) Children(children ...GmlElement) GmlElement {
	return e
}

var emptySlice = []byte{}

func (empty) RenderHtml(ctx context.Context) (html []byte, err error) {
	return emptySlice, nil
}
func (empty) Render(ctx context.Context, w io.Writer) error {
	return nil
}
func (empty) RenderBestEffort(ctx context.Context, w io.Writer) error {
	return nil
}
