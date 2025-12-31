package gml

func Html(attributes []Attr, children ...HtmlElement) Element {
	return newElement("html", false, attributes, children...)
}

func Head(attributes []Attr, children ...HtmlElement) Element {
	return newElement("head", false, attributes, children...)
}

// <body { attr }> { children } </body>
func Body(attributes []Attr, children ...HtmlElement) Element {
	return newElement("body", false, attributes, children...)
}

func Main(attributes []Attr, children ...HtmlElement) Element {
	return newElement("main", false, attributes, children...)
}

func Title(attributes []Attr, children ...HtmlElement) Element {
	return newElement("title", false, attributes, children...)
}

func Script(attributes []Attr, children ...HtmlElement) Element {
	return newElement("script", false, attributes, children...)
}

func Style(attributes []Attr, children ...HtmlElement) Element {
	return newElement("style", false, attributes, children...)
}

func Div(attributes []Attr, children ...HtmlElement) Element {
	return newElement("div", false, attributes, children...)
}

func P(attributes []Attr, children ...HtmlElement) Element {
	return newElement("p", false, attributes, children...)
}
func Button(attributes []Attr, children ...HtmlElement) Element {
	return newElement("button", false, attributes, children...)
}

func H1(attributes []Attr, children ...HtmlElement) Element {
	return newElement("h1", false, attributes, children...)
}

func H2(attributes []Attr, children ...HtmlElement) Element {
	return newElement("h2", false, attributes, children...)
}

func H3(attributes []Attr, children ...HtmlElement) Element {
	return newElement("h3", false, attributes, children...)
}

func Span(attributes []Attr, children ...HtmlElement) Element {
	return newElement("span", false, attributes, children...)
}

func Ul(attributes []Attr, children ...HtmlElement) Element {
	return newElement("ul", false, attributes, children...)
}

func Li(attributes []Attr, children ...HtmlElement) Element {
	return newElement("li", false, attributes, children...)
}

func A(attributes []Attr, children ...HtmlElement) Element {
	return newElement("a", false, attributes, children...)
}
