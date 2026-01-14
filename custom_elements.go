package gml

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/501urchin/gml/pkg"
)

// text content
type content struct {
	html string
}

var null = ""

// Content creates a content struct representing the inner content of an HTML GmlElement.
// It can also be used to render raw HTML values directly.
// Content accepts a variadic list of values; all arguments are concatenated in order
// without adding any separators or whitespace.
//
// For example, Content("hello", 12) renders as: hello12
func Content(t ...any) content {
	var builder strings.Builder
	builder.Grow(8 * len(t)) // default of 8 bytes since int is 64 bits

	for _, t := range t {
		switch v := t.(type) {
		case string:
			builder.WriteString(v)
		case fmt.Stringer:
			builder.WriteString(v.String())
		case nil:
			builder.WriteString(null)
		case []byte:
			builder.WriteString(pkg.BytesToString(v))
		case bool:
			builder.WriteString(strconv.FormatBool(v))
		case int:
			builder.WriteString(strconv.FormatInt(int64(v), 10))
		case int8:
			builder.WriteString(strconv.FormatInt(int64(v), 10))
		case int16:
			builder.WriteString(strconv.FormatInt(int64(v), 10))
		case int32:
			builder.WriteString(strconv.FormatInt(int64(v), 10))
		case int64:
			builder.WriteString(strconv.FormatInt(v, 10))
		case uint:
			builder.WriteString(strconv.FormatUint(uint64(v), 10))
		case uint8:
			builder.WriteString(strconv.FormatUint(uint64(v), 10))
		case uint16:
			builder.WriteString(strconv.FormatUint(uint64(v), 10))
		case uint32:
			builder.WriteString(strconv.FormatUint(uint64(v), 10))
		case uint64:
			builder.WriteString(strconv.FormatUint(v, 10))
		case float32:
			builder.WriteString(strconv.FormatFloat(float64(v), 'g', -1, 32))
		case float64:
			builder.WriteString(strconv.FormatFloat(v, 'g', -1, 64))
		default:
			// fall back to rendering it as json

			dt, _ := json.Marshal(v)
			builder.WriteString(pkg.BytesToString(dt))
		}
	}

	return content{html: builder.String()}
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
	return pkg.StringToBytes(t.html), nil
}

func (t content) Render(_ context.Context, w io.Writer) error {
	if len(t.html) == 0 {
		return nil
	}
	_, err := w.Write(pkg.StringToBytes(t.html))
	return err
}

func (t content) RenderBestEffort(_ context.Context, w io.Writer) error {
	if len(t.html) == 0 {
		return nil
	}
	_, err := w.Write(pkg.StringToBytes(t.html))
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
