package gml

const (
	elBr uint8 = iota + 100 // offset to avoid accidental overlap
	elHr
	elBase
	elLink
	elMeta
	elInput
	elWbr
	elCol
	elArea
	elEmbed
	elImg
	elParam
	elSource
	elTrack
	elDoctypeHTML
)

var voidElementNames = map[uint8]string{
	elBr:          "br",
	elHr:          "hr",
	elBase:        "base",
	elLink:        "link",
	elMeta:        "meta",
	elInput:       "input",
	elWbr:         "wbr",
	elCol:         "col",
	elArea:        "area",
	elEmbed:       "embed",
	elImg:         "img",
	elParam:       "param",
	elSource:      "source",
	elTrack:       "track",
	elDoctypeHTML: "!DOCTYPE html",
}

// <br { attr }>
func Br() *gmlElement { return newgmlElement(elBr, true) }

// <hr { attr }>
func Hr() *gmlElement { return newgmlElement(elHr, true) }

// <base { attr }>
func Base() *gmlElement { return newgmlElement(elBase, true) }

// <link { attr }>
func Link() *gmlElement { return newgmlElement(elLink, true) }

// <meta { attr }>
func Meta() *gmlElement { return newgmlElement(elMeta, true) }

// <input { attr }>
func Input() *gmlElement { return newgmlElement(elInput, true) }

// <wbr { attr }>
func Wbr() *gmlElement { return newgmlElement(elWbr, true) }

// <col { attr }>
func Col() *gmlElement { return newgmlElement(elCol, true) }

// <area { attr }>
func Area() *gmlElement { return newgmlElement(elArea, true) }

// <embed { attr }>
func Embed() *gmlElement { return newgmlElement(elEmbed, true) }

// <img { attr }>
func Img() *gmlElement { return newgmlElement(elImg, true) }

// <param { attr }>
func Param() *gmlElement { return newgmlElement(elParam, true) }

// <source { attr }>
func Source() *gmlElement { return newgmlElement(elSource, true) }

// <track { attr }>
func Track() *gmlElement { return newgmlElement(elTrack, true) }

// <!DOCTYPE html>
func DocTypeHtml() *gmlElement {
	return newgmlElement(elDoctypeHTML, true)
}
