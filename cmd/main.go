package main

import (
	"fmt"
	"go-fe-fwk/internals"
	"go-fe-fwk/types"
	"syscall/js"
)

func mountDom(this js.Value, args []js.Value) interface{} {
	props := types.Props{
		Attributes: types.Attributes[string]{
			"class": "colored",
		},
		On: types.EventHandlers{},
	}

	// l := []any{
	// 	internals.H("ul", types.Props{}, []any{
	// 		internals.H("li", types.Props{}, []any{"one"}),
	// 		internals.H("li", types.Props{}, []any{"one"}),
	// 	}),
	// }
	vdom := internals.H("p", props, []any{"aaaaaaaaaaaaa"})

	internals.MountDom(vdom, args[0])
	// time.Sleep(2 * time.Second)
	internals.DestroyDom(vdom)
	fmt.Println(args[0])
	return nil
}

func main() {

	fmt.Println("wasm connected")
	js.Global().Set("mountDom", js.FuncOf(mountDom))
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
