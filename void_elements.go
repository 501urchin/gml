package gml

// <br { attr }>
func Br(attributes []Attr) Element {
	return newElement("br", true, attributes)
}

// <hr { attr }>
func Hr(attributes []Attr) Element {
	return newElement("hr", true, attributes)
}

// <base { attr }>
func Base(attributes []Attr) Element {
	return newElement("base", true, attributes)
}

// <link { attr }>
func Link(attributes []Attr) Element {
	return newElement("link", true, attributes)
}

// <meta { attr }>
func Meta(attributes []Attr) Element {
	return newElement("meta", true, attributes)
}

// <input { attr }>
func Input(attributes []Attr) Element {
	return newElement("input", true, attributes)
}

// <wbr { attr }>
func Wbr(attributes []Attr) Element {
	return newElement("wbr", true, attributes)
}

// <col { attr }>
func Col(attributes []Attr) Element {
	return newElement("col", true, attributes)
}

// <area { attr }>
func Area(attributes []Attr) Element {
	return newElement("area", true, attributes)
}

// <embed { attr }>
func Embed(attributes []Attr) Element {
	return newElement("embed", true, attributes)
}

// <img { attr }>
func Img(attributes []Attr) Element {
	return newElement("img", true, attributes)
}

// <param { attr }>
func Param(attributes []Attr) Element {
	return newElement("param", true, attributes)
}

// <source { attr }>
func Source(attributes []Attr) Element {
	return newElement("source", true, attributes)
}

// <track { attr }>
func Track(attributes []Attr) Element {
	return newElement("track", true, attributes)
}
