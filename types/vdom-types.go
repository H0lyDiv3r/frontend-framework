package types

import "syscall/js"

type Attributes[T string | []string] map[string]T
type EventHandlers map[string]func(this js.Value, args []js.Value) interface{}
type Props struct {
	Attributes Attributes[string] `json:"attribute"`
	On         EventHandlers      `json:"on"`
}

var DOM_TYPES = map[string]string{
	"TEXT":     "text",
	"ELEMENT":  "element",
	"FRAGMENT": "fragment",
}
