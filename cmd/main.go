package main

import (
	"fmt"
	"go-fe-fwk/internals"
	"go-fe-fwk/types"
	"syscall/js"
)

var props = types.Props{
	Attributes: types.Attributes[string]{
		"class": "colored",
	},
	On: types.EventHandlers{},
}
var vdom = internals.H("div", props, []any{
	"aaaaaaaaa",
	internals.H("p", types.Props{}, []any{"awawaw"}),
})

func mountDom(this js.Value, args []js.Value) interface{} {

	// l := []any{
	// 	internals.H("ul", types.Props{}, []any{
	// 		internals.H("li", types.Props{}, []any{"one"}),
	// 		internals.H("li", types.Props{}, []any{"one"}),
	// 	}),
	// }

	// vdom.CreateNode(args[0])
	internals.MountDom(vdom, args[0])
	fmt.Println(vdom)

	// time.Sleep(2 * time.Second)
	// // fmt.Println(args[0], vdom)
	// internals.DestroyDom(vdom)
	return nil
}

func remDom(this js.Value, args []js.Value) interface{} {
	vdom.RemoveNode()
	return nil
}

func main() {

	fmt.Println("wasm connected")
	js.Global().Set("mountDom", js.FuncOf(mountDom))
	js.Global().Set("removeDom", js.FuncOf(remDom))
	<-make(chan struct{})
	// props := types.Props{
	// 	Attributes: types.Attributes[string]{
	// 		"aa": "aa",
	// 	},
	// 	On: types.EventHandlers{},
	// }

	// vdom := []any{
	// 	internals.H("p", props, []any{"paragraph"}),
	// }
	// jsonBytes, err := json.Marshal(internals.H("div", props, vdom))

	// if err != nil {
	// 	fmt.Println("there has been an error mate")
	// }

	// fmt.Println(string(jsonBytes))

}
