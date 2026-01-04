# gml

```go
	gml.Div().
		Attributes(
			gml.Class("flex flex-col"),
			gml.Id("demo"),
			gml.Attribute("aria-hidden", "false"),
		).Children(

		gml.P().
			Children(
				gmlx.If(45),
				gml.Content("hello from"),
			),

		gml.Raw("<p>GML !!!</p>"),
	)
```
