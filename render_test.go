package gml

import (
	"testing"
)

// func TestGmlElement(t *testing.T) {
// 	expected := `<h1 class="text-white">hello</h1>`

// 	attributes := []Attr{{Key: []byte("class"), Value: []byte("text-white")}}
// 	childElm := Content("hello")
// 	void := false
// 	gmlElementTag := elH1

// 	v := newgmlElement(gmlElementTag, void).Attributes(attributes...).Children(childElm)
// 	elm, ok := v.(*gmlElement)
// 	if !ok {
// 		t.Fatal("failed to cast newgmlElement as *gmlElement")
// 	}

// 	if elm.tag != gmlElementTag {
// 		t.Errorf("tag mismatch: expected %q, got %q", gmlElementTag, elm.tag)
// 	}

// 	if len(elm.children) != 1 {
// 		t.Errorf("children length mismatch: expected 1, got %d", len(elm.children))
// 	} else if html, _ := elm.children[0].RenderHtml(t.Context()); string(html) != string(childElm.html) {
// 		t.Errorf("child mismatch: expected %#v, got %#v", childElm, elm.children[0])
// 	}

// 	if len(elm.attributes) != 1 {
// 		t.Errorf("attributes length mismatch: expected 1, got %d", len(elm.attributes))
// 	} else if !bytes.Equal(elm.attributes[0].Key, attributes[0].Key) {
// 		t.Errorf("attribute mismatch: expected %#v, got %#v", attributes[0], elm.attributes[0])
// 	}

// 	if elm.void != void {
// 		t.Errorf("void flag mismatch: expected %t, got %t", void, elm.void)
// 	}

// 	// t.Run("render empty tag", func(t *testing.T) {
// 	// 	html, err := newgmlElement(0, void).RenderHtml(t.Context())
// 	// 	if err != nil {
// 	// 		t.Fatal(err)
// 	// 	}
// 	// 	if len(html) != 0 {
// 	// 		t.Errorf("render empty tag mismatch: expected empty string, got %q", html)
// 	// 	}
// 	// })

// 	t.Run("render", func(t *testing.T) {
// 		var buf bytes.Buffer

// 		err := elm.Render(t.Context(), &buf)
// 		if err != nil {
// 			t.Errorf("render returned error: %v", err)
// 		}

// 		if got := buf.String(); got != expected {
// 			t.Errorf("render output mismatch:\nexpected:\n%s\ngot:\n%s", expected, got)
// 		}
// 	})

// 	t.Run("render raw html", func(t *testing.T) {
// 		if got, _ := elm.RenderHtml(t.Context()); string(got) != expected {
// 			t.Errorf("renderHtml output mismatch:\nexpected:\n%s\ngot:\n%s", expected, got)
// 		}
// 	})

// 	t.Run("render nil as empty", func(t *testing.T) {
// 		elm := Div().Children(
// 			childElm,
// 			nil,
// 		)

// 		html, err := elm.RenderHtml(t.Context())
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		chtml, err := childElm.RenderHtml(t.Context())
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		if string(html) != "<div>"+string(chtml)+"</div>" {
// 			t.Error("function rendered a nil value")
// 		}
// 	})

// 	t.Run("nil receiver", func(t *testing.T) {
// 		defer func() {
// 			err := recover()
// 			if err != nil {
// 				t.Error("function failed to do nil checks and caused a panic")
// 			}
// 		}()

// 		var elm *gmlElement = nil

// 		elm.RenderHtml(t.Context())
// 	})

// }

var template = Html().
	Attributes(Attribute("lang", "en")).
	Children(
		Empty(),
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
					Content("gml!!!!"),
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
					Content("console.log('Hello from GML !!!')"),
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
								Content("0"),
							),

						Button().
							Attributes(
								Attribute("@click", "count += 1"),
								Attribute("class", "col-span-1 p-1 w-full cursor-pointer text-sm text-white select-none bg-neutral-900 rounded-md active:scale-98  will-change-transform ease-in-out"),
							).
							Children(
								Content("+"),
							),

						Button().
							Attributes(
								Attribute("@click", "count > 0  ? count -= 1  : null"),
								Attribute(":class", "count === 0 ? 'cursor-not-allowed grayscale' : 'cursor-pointer active:scale-98  will-change-transform ease-in-out '"),
								Attribute("class", "col-span-1 p-1 w-full  text-sm text-white bg-neutral-900 rounded-md"),
							).
							Children(
								Content("-"),
							),

						Button().
							Attributes(
								Attribute("@click", "count = 0"),
								Attribute(":class", "count === 0 ? 'cursor-not-allowed grayscale' : 'cursor-pointer active:scale-98  will-change-transform ease-in-out '"),
								Attribute("class", "col-span-2 p-1 w-full  text-sm text-white bg-neutral-900 rounded-md"),
							).
							Children(
								Content("⟳"),
							),
					),
			),
	)

// bofore gopt: enchmarkRender-10        851790              1400 ns/op            4296 B/op         34 allocs/op
// after gopt: BenchmarkRender-10       1077978              1108 ns/op            4080 B/op          7 allocs/op
func BenchmarkRender(b *testing.B) {
	for b.Loop() {
		template.RenderHtml(b.Context())
	}
}

func BenchmarkNewElement(b *testing.B) {
	for b.Loop() {
		newgmlElement(elDiv, false)
	}
}
