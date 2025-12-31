package gml

import (
	"os"
	"testing"
)

func BenchmarkStream(b *testing.B) {
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
			template.RenderStream(file)
		}

	})

	b.Run("render normal", func(b *testing.B) {

		for b.Loop() {
			template.Render()
		}
	})
}
