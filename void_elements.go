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
func Br() *gmlElement { return newGmlElement(elBr, true) }

// <hr { attr }>
func Hr() *gmlElement { return newGmlElement(elHr, true) }

// <base { attr }>
func Base() *gmlElement { return newGmlElement(elBase, true) }

// <link { attr }>
func Link() *gmlElement { return newGmlElement(elLink, true) }

// <meta { attr }>
func Meta() *gmlElement { return newGmlElement(elMeta, true) }

// <input { attr }>
func Input() *gmlElement { return newGmlElement(elInput, true) }

// <wbr { attr }>
func Wbr() *gmlElement { return newGmlElement(elWbr, true) }

// <col { attr }>
func Col() *gmlElement { return newGmlElement(elCol, true) }

// <area { attr }>
func Area() *gmlElement { return newGmlElement(elArea, true) }

// <embed { attr }>
func Embed() *gmlElement { return newGmlElement(elEmbed, true) }

// <img { attr }>
func Img() *gmlElement { return newGmlElement(elImg, true) }

// <param { attr }>
func Param() *gmlElement { return newGmlElement(elParam, true) }

// <source { attr }>
func Source() *gmlElement { return newGmlElement(elSource, true) }

// <track { attr }>
func Track() *gmlElement { return newGmlElement(elTrack, true) }

// <!DOCTYPE html>
func DocTypeHtml() *gmlElement {
	return newGmlElement(elDoctypeHTML, true)
}
