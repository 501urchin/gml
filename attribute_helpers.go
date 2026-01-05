package gml

import "github.com/501urchin/gopt"

var classKey = []byte("class")

func Class(value string) Attr {
	return Attr{Key: classKey, Value: gopt.StringToBytes(value)}
}

var idKey = []byte("id")

func Id(value string) Attr {
	return Attr{Key: idKey, Value: gopt.StringToBytes(value)}
}

var srcKey = []byte("src")

func Src(value string) Attr {
	return Attr{Key: srcKey, Value: gopt.StringToBytes(value)}
}
