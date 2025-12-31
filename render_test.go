package gml

import (
	"bytes"
	"os"
	"testing"
)

func TestElement(t *testing.T) {
	expected := `<h1 class="text-white">hello</h1>`

	attributes := []Attr{{Key: "class", Value: "text-white"}}
	childElm := Text("hello")
	void := false
	elementTag := "h1"

	elm := newElement(elementTag, void, attributes, childElm)

	if elm.tag != elementTag {
		t.Errorf("tag mismatch: expected %q, got %q", elementTag, elm.tag)
	}

	if len(elm.children) != 1 {
		t.Errorf("children length mismatch: expected 1, got %d", len(elm.children))
	} else if elm.children[0] != childElm {
		t.Errorf("child mismatch: expected %#v, got %#v", childElm, elm.children[0])
	}

	if len(elm.attributes) != 1 {
		t.Errorf("attributes length mismatch: expected 1, got %d", len(elm.attributes))
	} else if elm.attributes[0] != attributes[0] {
		t.Errorf("attribute mismatch: expected %#v, got %#v", attributes[0], elm.attributes[0])
	}

	if elm.void != void {
		t.Errorf("void flag mismatch: expected %t, got %t", void, elm.void)
	}

	t.Run("render empty tag", func(t *testing.T) {
		html := newElement("", void, attributes, childElm).RenderHtml()
		if html != "" {
			t.Errorf("render empty tag mismatch: expected empty string, got %q", html)
		}
	})

	t.Run("render", func(t *testing.T) {
		var buf bytes.Buffer

		err := elm.Render(&buf)
		if err != nil {
			t.Errorf("render returned error: %v", err)
		}

		if got := buf.String(); got != expected {
			t.Errorf("render output mismatch:\nexpected:\n%s\ngot:\n%s", expected, got)
		}
	})

	t.Run("render raw html", func(t *testing.T) {
		if got := elm.RenderHtml(); got != expected {
			t.Errorf("renderHtml output mismatch:\nexpected:\n%s\ngot:\n%s", expected, got)
		}
	})

}

func BenchmarkRender(b *testing.B) {
	template := Html(nil,
		Head(nil,
			Title(nil, Text("gsx!!!!")),
			Script([]Attr{
				Attribute("src", "https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4"),
			}),
			Script([]Attr{
				Attribute("src", "https://unpkg.com/alpinejs@3.x.x/dist/cdn.min.js"),
				Attribute("defer", ""),
			}),
		),
		Body(
			[]Attr{
				Attribute("class", "flex bg-black min-h-screen w-full items-center justify-center"),
			},
			Script(nil, Text("console.log('Hello from GML !!!')")),
			Div([]Attr{
				Attribute("x-data", "{'count': 0}"),
				Attribute("class", "grid grid-cols-2 w-50 gap-2"),
			},
				P([]Attr{
					Attribute("x-text", "count"),
					Attribute("class", "text-white bg-neutral-800 rounded-md p-1 col-span-2 text-center"),
				}),

				Button([]Attr{
					Attribute("@click", "count += 1"),
					Attribute("class", "col-span-1 p-1 w-full cursor-pointer text-sm text-white select-none bg-neutral-900 rounded-md active:scale-98  will-change-transform ease-in-out"),
				}, Text("+")),

				Button([]Attr{
					Attribute("@click", "count > 0  ? count -= 1  : null"),
					Attribute(":class", "count === 0 ? 'cursor-not-allowed grayscale' : 'cursor-pointer active:scale-98  will-change-transform ease-in-out '"),
					Attribute("class", "col-span-1 p-1 w-full  text-sm text-white bg-neutral-900 rounded-md"),
				}, Text("-")),

				Button([]Attr{
					Attribute("@click", "count = 0"),
					Attribute(":class", "count === 0 ? 'cursor-not-allowed grayscale' : 'cursor-pointer active:scale-98  will-change-transform ease-in-out '"),
					Attribute("class", "col-span-2 p-1 w-full  text-sm text-white bg-neutral-900 rounded-md"),
				}, Text("⟳")),
			),
		),
	)

	b.Run("stream", func(b *testing.B) {
		file, err := os.Open(os.DevNull)
		if err != nil {
			b.Fatal(err)
		}
		defer file.Close()

		b.ResetTimer()
		for b.Loop() {
			template.Render(file)
		}

	})

	b.Run("render normal", func(b *testing.B) {

		for b.Loop() {
			template.RenderHtml()
		}
	})
}
