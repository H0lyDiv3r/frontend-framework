package types

import "syscall/js"

type Attributes[T string | []string] map[string]T
type EventHandlers map[string]js.Value
type Children[T StringDom | Vdom] []T
type Props struct {
	Attributes Attributes[string] `json:"attribute"`
	On         EventHandlers      `json:"on"`
}

var DOM_TYPES = map[string]string{
	"TEXT":     "text",
	"ELEMENT":  "element",
	"FRAGMENT": "fragment",
}

type Vdom struct {
	Tag       string              `json:"tag"`
	Props     Props               `json:"props"`
	Children  []any               `json:"children"`
	Type      string              `json:"type"`
	Listeners map[string]js.Value `json:"listeners"`
	El        js.Value            `json:"el"`
}

type StringDom struct {
	Type  string   `json:"type"`
	Value string   `json:"value"`
	El    js.Value `json:"el"`
}
