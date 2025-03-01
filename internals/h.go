package internals

import (
	"go-fe-fwk/types"
	"reflect"
)

func H(tag string, props types.Props, children []any) types.Vdom {
	newChildren := []any{}
	for i := range children {
		if reflect.TypeOf(children[i]) == reflect.TypeOf("") {
			newChildren = append(newChildren, Hstring(children[i].(string)))
		} else {
			newChildren = append(newChildren, children[i])
		}
	}
	return types.Vdom{
		Tag:      tag,
		Props:    props,
		Children: newChildren,
		Type:     types.DOM_TYPES["ELEMENT"],
	}
}

func Hstring(value string) types.StringDom {
	return types.StringDom{
		Type:  types.DOM_TYPES["TEXT"],
		Value: value,
	}
}
