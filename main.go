package main

import (
	"fmt"
	"net/http"

	"github.com/501urchin/go-playground/html"
)

func Simple() string {
	return html.Html(nil,
		html.Head(nil,
			html.Title(nil, html.Text("gsx!!!!")),
			html.Script([]html.Attr{
				html.Attribute("src", "https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4"),
			}),
			html.Script([]html.Attr{
				html.Attribute("src", "https://unpkg.com/alpinejs@3.x.x/dist/cdn.min.js"),
				html.Attribute("defer", ""),
			}),
		),
		html.Body(
			[]html.Attr{
				html.Attribute("class", "flex bg-black min-h-screen w-full items-center justify-center"),
			},
			html.Script(nil, html.Text("console.log('Hello from GML !!!')")),
			html.Div([]html.Attr{
				html.Attribute("x-data", "{'count': 0}"),
				html.Attribute("class", "grid grid-cols-2 w-50 gap-2"),
			},
				html.P([]html.Attr{
					html.Attribute("x-text", "count"),
					html.Attribute("class", "text-white bg-neutral-800 rounded-md p-1 col-span-2 text-center"),
				}),

				html.Button([]html.Attr{
					html.Attribute("@click", "count += 1"),
					html.Attribute("class", "col-span-1 p-1 w-full cursor-pointer text-sm text-white select-none bg-neutral-900 rounded-md active:scale-98  will-change-transform ease-in-out"),
				}, html.Text("+")),

				html.Button([]html.Attr{
					html.Attribute("@click", "count > 0  ? count -= 1  : null"),
					html.Attribute(":class", "count === 0 ? 'cursor-not-allowed grayscale' : 'cursor-pointer active:scale-98  will-change-transform ease-in-out '"),
					html.Attribute("class", "col-span-1 p-1 w-full  text-sm text-white bg-neutral-900 rounded-md"),
				}, html.Text("-")),

				html.Button([]html.Attr{
					html.Attribute("@click", "count = 0"),
					html.Attribute(":class", "count === 0 ? 'cursor-not-allowed grayscale' : 'cursor-pointer active:scale-98  will-change-transform ease-in-out '"),
					html.Attribute("class", "col-span-2 p-1 w-full  text-sm text-white bg-neutral-900 rounded-md"),
				}, html.Text("⟳")),
			),
		),
	).Render()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<!DOCTYPE html>" + Simple()))
	})

	fmt.Println("listening on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
