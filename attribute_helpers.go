package gml

var classKey = []byte("class")

func Class(value string) Attr {
	return Attr{Key: classKey, Value: []byte(value)}
}

var idKey = []byte("id")

func Id(value string) Attr {
	return Attr{Key: idKey, Value: []byte(value)}
}

var srcKey = []byte("src")

func Src(value string) Attr {
	return Attr{Key: srcKey, Value: []byte(value)}
}
