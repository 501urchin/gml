package gml

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
)

// text content
type content struct {
	html []byte
}

var null = []byte{}

// Content creates a content struct representing the inner content of an HTML gmlElement.
// It can also be used to render raw HTML strings directly.
func Content(t any) content {
	var parsed []byte

	switch v := t.(type) {
	case string:
		parsed = []byte(v)
	case fmt.Stringer:
		parsed = []byte(v.String())
	case nil:
		parsed = null
	case []byte:
		parsed = v
	case bool:
		parsed = strconv.AppendBool(parsed, v)
	case int:
		parsed = strconv.AppendInt(parsed, int64(v), 10)
	case int8:
		parsed = strconv.AppendInt(parsed, int64(v), 10)
	case int16:
		parsed = strconv.AppendInt(parsed, int64(v), 10)
	case int32:
		parsed = strconv.AppendInt(parsed, int64(v), 10)
	case int64:
		parsed = strconv.AppendInt(parsed, v, 10)
	case uint:
		parsed = strconv.AppendUint(parsed, uint64(v), 10)
	case uint8:
		parsed = strconv.AppendUint(parsed, uint64(v), 10)
	case uint16:
		parsed = strconv.AppendUint(parsed, uint64(v), 10)
	case uint32:
		parsed = strconv.AppendUint(parsed, uint64(v), 10)
	case uint64:
		parsed = strconv.AppendUint(parsed, v, 10)
	case float32:
		parsed = strconv.AppendFloat(parsed, float64(v), 'g', -1, 32)
	case float64:
		parsed = strconv.AppendFloat(parsed, v, 'g', -1, 64)
	default:
		// fall back to rendering it as json
		parsed, _ = json.Marshal(v)
	}
	return content{html: parsed}
}

// function is used to render raw html
func Raw(t string) content {
	return content{html: []byte(t)}
}

func (t content) Attributes(_ ...Attr) GmlElement {
	return t
}
func (t content) Children(_ ...GmlElement) GmlElement {
	return t
}

func (t content) RenderHtml(ctx context.Context) (html []byte, err error) {
	return t.html, nil
}

func (t content) Render(_ context.Context, w io.Writer) error {
	if len(t.html) == 0 {
		return nil
	}
	_, err := w.Write(t.html)
	return err
}

func (t content) RenderBestEffort(_ context.Context, w io.Writer) error {
	if len(t.html) == 0 {
		return nil
	}
	_, err := w.Write(t.html)
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
