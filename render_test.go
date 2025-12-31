package gml

import (
	"bytes"
	"os"
	"strconv"
	"testing"
)

func TestElement(t *testing.T) {
	expected := `<h1 class="text-white">hello</h1>`

	attributes := []Attr{{Key: "class", Value: "text-white"}}
	childElm := Text("hello")
	void := false
	elementTag := "h1"

	v := newElement(elementTag, void).Attributes(attributes...).Children(childElm)
	elm, ok := v.(*Element)
	if !ok {
		t.Fatal("failed to cast newElement as *Element")
	}

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
		html := newElement("", void).RenderHtml(t.Context())
		if html != "" {
			t.Errorf("render empty tag mismatch: expected empty string, got %q", html)
		}
	})

	t.Run("render", func(t *testing.T) {
		var buf bytes.Buffer

		err := elm.Render(t.Context(), &buf)
		if err != nil {
			t.Errorf("render returned error: %v", err)
		}

		if got := buf.String(); got != expected {
			t.Errorf("render output mismatch:\nexpected:\n%s\ngot:\n%s", expected, got)
		}
	})

	t.Run("render raw html", func(t *testing.T) {
		if got := elm.RenderHtml(t.Context()); got != expected {
			t.Errorf("renderHtml output mismatch:\nexpected:\n%s\ngot:\n%s", expected, got)
		}
	})

}

func BenchmarkRender(b *testing.B) {
	template := Html().
		Attributes(Attribute("lang", "en")).
		Children(
			Head().
				Children(
					Meta().Attributes(Attribute("charset", "UTF-8")),
					Meta().Attributes(
						Attribute("name", "viewport"),
						Attribute("content", "width=device-width, initial-scale=1.0"),
					),
					Meta().Attributes(
						Attribute("name", "description"),
						Attribute("content", "counter demo with gml"),
					),
					Title().Children(
						Text("gml!!!!"),
					),
					Script().Attributes(
						Attribute("src", "https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4"),
					),
					Script().Attributes(
						Attribute("src", "https://unpkg.com/alpinejs@3.x.x/dist/cdn.min.js"),
						Attribute("defer", ""),
					),
				),
			Body().
				Attributes(Attribute("class", "flex bg-black min-h-screen w-full items-center justify-center")).
				Children(
					Script().Children(
						Text("console.log('Hello from GML !!!')"),
					),
					Main().
						Attributes(
							Attribute("x-data", "{'count': 0}"),
							Attribute("class", "grid grid-cols-2 w-50 gap-2"),
						).
						Children(
							P().
								Attributes(
									Attribute("x-text", "count"),
									Attribute("class", "text-white bg-neutral-800 rounded-md p-1 col-span-2 text-center"),
								).
								Children(
									Text("0"),
								),

							Button().
								Attributes(
									Attribute("@click", "count += 1"),
									Attribute("class", "col-span-1 p-1 w-full cursor-pointer text-sm text-white select-none bg-neutral-900 rounded-md active:scale-98  will-change-transform ease-in-out"),
								).
								Children(
									Text("+"),
								),

							Button().
								Attributes(
									Attribute("@click", "count > 0  ? count -= 1  : null"),
									Attribute(":class", "count === 0 ? 'cursor-not-allowed grayscale' : 'cursor-pointer active:scale-98  will-change-transform ease-in-out '"),
									Attribute("class", "col-span-1 p-1 w-full  text-sm text-white bg-neutral-900 rounded-md"),
								).
								Children(
									Text("-"),
								),

							Button().
								Attributes(
									Attribute("@click", "count = 0"),
									Attribute(":class", "count === 0 ? 'cursor-not-allowed grayscale' : 'cursor-pointer active:scale-98  will-change-transform ease-in-out '"),
									Attribute("class", "col-span-2 p-1 w-full  text-sm text-white bg-neutral-900 rounded-md"),
								).
								Children(
									Text("⟳"),
								),

							For(func() []HtmlElement {
								var buf []HtmlElement
								for i := range 3 {
									buf = append(buf,
										P().
											Attributes(
												Attribute("class", "text-white col-span-full text-center text-sm"),
											).
											Children(
												Text(strconv.Itoa(i)),
											),
									)
								}
								return buf
							}),
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
			template.Render(b.Context(), file)
		}

	})

	b.Run("render normal", func(b *testing.B) {

		for b.Loop() {
			template.RenderHtml(b.Context())
		}
	})
}
