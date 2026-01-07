package gml

var classKey = "class"

func Class(value string) Attr {
	return Attr{Key: classKey, Value: value}
}

var idKey = "id"

func Id(value string) Attr {
	return Attr{Key: idKey, Value: value}
}

var srcKey = "src"

func Src(value string) Attr {
	return Attr{Key: srcKey, Value: value}
}

var hrefKey = "href"

func Href(value string) Attr {
	return Attr{Key: hrefKey, Value: value}
}

var altKey = "alt"

func Alt(value string) Attr {
	return Attr{Key: altKey, Value: value}
}

var typeKey = "type"

func Type(value string) Attr {
	return Attr{Key: typeKey, Value: value}
}

var valueKey = "value"

func Value(value string) Attr {
	return Attr{Key: valueKey, Value: value}
}

var nameKey = "name"

func Name(value string) Attr {
	return Attr{Key: nameKey, Value: value}
}

var placeholderKey = "placeholder"

func Placeholder(value string) Attr {
	return Attr{Key: placeholderKey, Value: value}
}

var disabledKey = "disabled"

func Disabled() Attr {
	return Attr{Key: disabledKey, Value: ""}
}

var checkedKey = "checked"

func Checked() Attr {
	return Attr{Key: checkedKey, Value: ""}
}

var forKey = "for"

func For(value string) Attr {
	return Attr{Key: forKey, Value: value}
}

var targetKey = "target"

func Target(value string) Attr {
	return Attr{Key: targetKey, Value: value}
}

var relKey = "rel"

func Rel(value string) Attr {
	return Attr{Key: relKey, Value: value}
}

var widthKey = "width"

func Width(value string) Attr {
	return Attr{Key: widthKey, Value: value}
}

var heightKey = "height"

func Height(value string) Attr {
	return Attr{Key: heightKey, Value: value}
}
