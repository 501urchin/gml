package gml

// <br { attr }>
func Br() *Element {
	return newElement("br", true)
}

// <hr { attr }>
func Hr() *Element {
	return newElement("hr", true)
}

// <base { attr }>
func Base() *Element {
	return newElement("base", true)
}

// <link { attr }>
func Link() *Element {
	return newElement("link", true)
}

// <meta { attr }>
func Meta() *Element {
	return newElement("meta", true)
}

// <input { attr }>
func Input() *Element {
	return newElement("input", true)
}

// <wbr { attr }>
func Wbr() *Element {
	return newElement("wbr", true)
}

// <col { attr }>
func Col() *Element {
	return newElement("col", true)
}

// <area { attr }>
func Area() *Element {
	return newElement("area", true)
}

// <embed { attr }>
func Embed() *Element {
	return newElement("embed", true)
}

// <img { attr }>
func Img() *Element {
	return newElement("img", true)
}

// <param { attr }>
func Param() *Element {
	return newElement("param", true)
}

// <source { attr }>
func Source() *Element {
	return newElement("source", true)
}

// <track { attr }>
func Track() *Element {
	return newElement("track", true)
}
