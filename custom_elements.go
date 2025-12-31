package gml

// text element
type text struct {
	htmlString string
}

func (t text) Render() string {
	return t.htmlString
}

func Text(t string) text {
	return text{htmlString: t}
}

func Meta(attributes []Attr) Element {
	return newElement("meta", true, attributes)
}

func Link(attributes []Attr) Element {
	return newElement("link", true, attributes)
}
