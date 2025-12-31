package gml

// <br { attr }>
func Br() *gmlElement {
	return newgmlElement("br", true)
}

// <hr { attr }>
func Hr() *gmlElement {
	return newgmlElement("hr", true)
}

// <base { attr }>
func Base() *gmlElement {
	return newgmlElement("base", true)
}

// <link { attr }>
func Link() *gmlElement {
	return newgmlElement("link", true)
}

// <meta { attr }>
func Meta() *gmlElement {
	return newgmlElement("meta", true)
}

// <input { attr }>
func Input() *gmlElement {
	return newgmlElement("input", true)
}

// <wbr { attr }>
func Wbr() *gmlElement {
	return newgmlElement("wbr", true)
}

// <col { attr }>
func Col() *gmlElement {
	return newgmlElement("col", true)
}

// <area { attr }>
func Area() *gmlElement {
	return newgmlElement("area", true)
}

// <embed { attr }>
func Embed() *gmlElement {
	return newgmlElement("embed", true)
}

// <img { attr }>
func Img() *gmlElement {
	return newgmlElement("img", true)
}

// <param { attr }>
func Param() *gmlElement {
	return newgmlElement("param", true)
}

// <source { attr }>
func Source() *gmlElement {
	return newgmlElement("source", true)
}

// <track { attr }>
func Track() *gmlElement {
	return newgmlElement("track", true)
}
