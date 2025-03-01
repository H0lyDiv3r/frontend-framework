package internals

import (
	"go-fe-fwk/pkgs/utils"
	"go-fe-fwk/types"
	"syscall/js"
)

func MountDom[T types.StringDom | types.Vdom | any](vdom T, parentEl js.Value) {
	switch v := any(vdom).(type) {
	case types.StringDom:
		CreateTextNode(v, parentEl)
	case types.Vdom:
		CreateElementNode(v, parentEl)
	}
}

func CreateTextNode(vdom types.StringDom, parentEl js.Value) {
	document := js.Global().Get("document")
	value := vdom.Value
	textNode := document.Call("createTextNode", value)
	vdom.El = textNode
	parentEl.Call("appendChild", textNode)
}

func CreateElementNode(vdom types.Vdom, parentEl js.Value) {
	document := js.Global().Get("document")
	tag := vdom.Tag
	props := vdom.Props
	children := vdom.Children
	element := document.Call("createElement", tag)
	addProps(element, props, vdom)
	vdom.El = element
	for _, child := range children {
		MountDom(child, element)
	}
	parentEl.Call("appendChild", element)
}

func addProps(element js.Value, props types.Props, vdom types.Vdom) {
	attributes := props.Attributes
	eventHandlers := props.On
	vdom.Listeners = utils.AddEventListeners(eventHandlers, element)
	utils.SetAttributes(element, attributes)
}
