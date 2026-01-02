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

func IfAttr(cond bool, fn func() gml.AttributeIface) gml.AttributeIface {
	if !cond {
		return gml.Attr{}
	}

	if fn == nil {
		return gml.Attr{}
	}

	return fn()
}
