package gml

import "github.com/501urchin/gml/internal"


var classKey = []byte("class")

func Class(value string) Attr {
	return Attr{Key: classKey, Value: internal.StringToBytes(value)}
}

var idKey = []byte("id")

func Id(value string) Attr {
	return Attr{Key: idKey, Value: internal.StringToBytes(value)}
}

var srcKey = []byte("src")

func Src(value string) Attr {
	return Attr{Key: srcKey, Value: internal.StringToBytes(value)}
}

var hrefKey = []byte("href")

func Href(value string) Attr {
	return Attr{Key: hrefKey, Value: internal.StringToBytes(value)}
}

var altKey = []byte("alt")

func Alt(value string) Attr {
	return Attr{Key: altKey, Value: internal.StringToBytes(value)}
}

var typeKey = []byte("type")

func Type(value string) Attr {
	return Attr{Key: typeKey, Value: internal.StringToBytes(value)}
}

var valueKey = []byte("value")

func Value(value string) Attr {
	return Attr{Key: valueKey, Value: internal.StringToBytes(value)}
}

var nameKey = []byte("name")

func Name(value string) Attr {
	return Attr{Key: nameKey, Value: internal.StringToBytes(value)}
}

var placeholderKey = []byte("placeholder")

func Placeholder(value string) Attr {
	return Attr{Key: placeholderKey, Value: internal.StringToBytes(value)}
}

var disabledKey = []byte("disabled")

func Disabled() Attr {
	return Attr{Key: disabledKey, Value: nil}
}

var checkedKey = []byte("checked")

func Checked() Attr {
	return Attr{Key: checkedKey, Value: nil}
}

var forKey = []byte("for")

func For(value string) Attr {
	return Attr{Key: forKey, Value: internal.StringToBytes(value)}
}

var targetKey = []byte("target")

func Target(value string) Attr {
	return Attr{Key: targetKey, Value: internal.StringToBytes(value)}
}

var relKey = []byte("rel")

func Rel(value string) Attr {
	return Attr{Key: relKey, Value: internal.StringToBytes(value)}
}

var widthKey = []byte("width")

func Width(value string) Attr {
	return Attr{Key: widthKey, Value: internal.StringToBytes(value)}
}

var heightKey = []byte("height")

func Height(value string) Attr {
	return Attr{Key: heightKey, Value: internal.StringToBytes(value)}
}
