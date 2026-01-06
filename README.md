# gml

**gml** is a lightweight, composable DSL for building HTML interfaces directly in Go. It allows you to describe HTML structure using idiomatic Go code while preserving clarity, type safety, and explicit control over rendering.

The library is designed to be minimal, predictable, and runtime free.

### Go (gml)

```go
import "github.com/501urchin/gml"

func Example() gml.GmlElement {
	return gml.Div().
		Attributes(
			gml.Class("flex flex-col"),
			gml.Id("demo"),
			gml.Attribute("aria-hidden", "false"),
		).
		Children(
			gml.P().Children(
				gml.Content("hello from"),
			),

			gml.Raw("<p>GML !!!</p>"),
		)
}
```

### Equivalent HTML

```html
<div class="flex flex-col" id="demo" aria-hidden="false">
  <p>hello from</p>
  <p>GML !!!</p>
</div>
```

## Table of Contents

- [How It Works](#how-it-works)
  - [Element Model](#element-model)
  - [Rendering Strategy](#rendering-strategy)
- [Design Goals](#design-goals)
- [Elements, Attributes, and Children](#elements-attributes-and-children)
  - [Creating an Element](#creating-an-element)
  - [Adding Attributes](#adding-attributes)
  - [Adding Children](#adding-children)
  - [Text Content](#text-content)
- [gmlx Extensions](#gmlx-extensions)
  - [Conditional Rendering](#conditional-rendering)
  - [Iteration](#iteration)
  - [Grouping Elements](#grouping-elements)


## How It Works

At the core of **gml** is the `GmlElement` interface. Every element in the library such as `gml.Div()`, `gml.P()`, or `gml.Raw()` implements this interface.

```go
type GmlElement interface {
	// Attributes returns the receiver with the provided attributes appended.
	Attributes(attributes ...Attr) GmlElement

	// Children returns the receiver with the provided child elements appended.
	Children(children ...GmlElement) GmlElement

	// Render writes the rendered HTML representation of the element to w.
	// If rendering this element or any of its children returns an error,
	// rendering stops immediately and that error is returned.
	Render(ctx context.Context, w io.Writer) error

	// RenderBestEffort writes the rendered HTML representation of the element to w.
	// Unlike Render, this method continues rendering remaining children even if
	// one or more children return an error. The returned error, if non-nil,
	// represents at least one failure during rendering.
	RenderBestEffort(ctx context.Context, w io.Writer) error

	// RenderHtml renders the element and returns the resulting HTML as a byte slice.
	// This is intended for use cases that require raw HTML output rather than
	// streaming to an io.Writer.
	RenderHtml(ctx context.Context) ([]byte, error)
}
```

### Element Model

* Each HTML construct is represented as a Go value implementing `GmlElement`.
* Elements are composed by calling `Attributes(...)` and `Children(...)`.
* Rendering is explicit and opt-in via one of the `Render` methods.

### Rendering Strategy

**gml** has no runtime or virtual DOM. Rendering is performed via straightforward recursion:

* Each element knows how to render itself.
* Child elements are rendered by calling their `Render` method.
* The full HTML output is produced by recursively walking the element tree.

Because all elements implement the same interface, you can define custom elements by implementing `GmlElement` yourself and seamlessly integrate them into existing trees.


## Design Goals

* **Explicit rendering**: you control when and how HTML is produced
* **Composable primitives**: build higher-level components from simple elements
* **Idiomatic Go**: no templates, no DSL parser, just Go


## Elements, Attributes, and Children

**gml** implements most standard HTML elements (for example `div`, `p`, `span`, etc.). Each element is created by calling its corresponding constructor function.

### Creating an Element

To render a `div`, you first create the element:

```go
gml.Div()
```

This call returns an internal concrete type that implements the `GmlElement` interface. At this stage, the element has no attributes and no children it is just a value representing a `div` node.


### Adding Attributes

Attributes are added by calling the `Attributes(...)` method. Each attribute is represented by the `Attr` type. gml's internal type appends the attribute to the attributes array and returns the receiver

```go
gml.Div().Attributes(
	gml.Attribute("class", "myclass"),
)
```

The `gml.Attribute` function accepts a key and a value and returns an `Attr`.

### Behavior

The internal element type appends the provided attributes to its attribute slice and returns the receiver.

Repeated calls to `.Attributes()` add more attributes to the existing list rather than replacing them.


```go
gml.Div().Attributes(
    gml.Attribute("class", "myclass"),
).Attributes(
    gml.Id("demo"),
)
```
The resulting div has both class="myclass" and id="demo".

> **Note**: Attribute values are **not escaped automatically**. If you are inserting untrusted or structured data, you must escape it yourself.

```go
gml.Div().Attributes(
	gml.Attribute("data-tag", html.EscapeString(jsonString)),
)
```

Because repeatedly calling `gml.Attribute(key, value)` can be verbose, **gml** provides convenience helpers for commonly used attributes.

```go
gml.Div().Attributes(
	gml.Class("myclass"),
)
```

These helpers return the same `Attr` type, but with clearer intent and less boilerplate.

Here’s a rewritten version for **children** instead of attributes, keeping the same style and structure:

### Adding Children

Child elements are added using the `Children(...)` method. This method accepts any value that implements the `GmlElement` interface.

```go
gml.Div().Children(
    gml.Content("child1"),
    gml.P(),
)
```

### Behavior

The internal element type appends the provided children to its children slice and returns the receiver.

Repeated calls to `.Children()` add more children to the existing list rather than replacing them.

```go
gml.Div().Children(
    gml.Content("child1"),
).Children(
    gml.Content("child2"),
)
```

The resulting `<div>` contains both `"child1"` and `"child2"` as children.

## Content

`gml.Content` is a built-in element that implements `GmlElement` and renders arbitrary values as text. It accepts a value of type `any` and attempts to convert it to a string representation.

Examples:

```go
gml.Content("child")  // renders: "child"
gml.Content(1)        // renders: "1"
gml.Content(false)    // renders: "false"
gml.Content(nil)      // renders: "nil"
```

This makes it convenient to inject dynamic values without manually converting them beforehand.
> **Note**: Content values are **not escaped automatically**. If you are inserting untrusted or structured data, you must escape it yourself.

## Raw HTML

Sometimes, we want to render raw HTML directly.

`gml` provides a `gml.Raw()` method for this purpose. It accepts a string containing the HTML you want to render.


```go
gml.Raw("<p>hello</p>")
```

## Extensions

**gmlx** is an extension package built on top of **gml** that provides higher-level helpers and custom elements for common control-flow and composition patterns.

Key features include:

* **`gmlx.If`** – conditional rendering
* **`gmlx.Group`** – grouping multiple elements into a single `GmlElement`
* **`gmlx.Map`** – iterating over collections to produce element trees

### Conditional Rendering

```go
gmlx.If(LoggedIn, func() gml.GmlElement {
	return gml.Content("LoggedIn")
}).Else(func() gml.GmlElement {
	return gml.Content("not logged in")
})
```

### Iteration

```go
var items []customType

gmlx.Map(items, func(index int, item customType) gml.GmlElement {
	return gml.Div().Children(
		gml.Content(item.Field1),
		gml.Content(item.Field2),
	)
})
```

### Grouping Elements

```go
gmlx.Group(
	gml.DocTypeHtml(),
	gml.Html(),
)
```
