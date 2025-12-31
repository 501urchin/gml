package gmlx

import "github.com/501urchin/gml"

func If(cond bool, fn func() gml.GmlElement) gml.GmlElement {
	if !cond {
		return gml.Empty()
	}

	if fn == nil {
		return gml.Empty()
	}

	return fn()
}
