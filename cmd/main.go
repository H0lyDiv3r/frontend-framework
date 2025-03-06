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
	On: types.EventHandlers{
		"click": func(this js.Value, args []js.Value) interface{} {
			fmt.Println("aaaaaaaaaaaaaaaaaaaaaaauuuuuuuu")
			return nil
		},
	},
}
var vdom = internals.H("div", props, []any{
	"aaaaaaaaa",
	internals.H("p", types.Props{}, []any{"awawaw"}),
})

func mountDom(this js.Value, args []js.Value) interface{} {
	vdm := internals.H("button", props, []any{"aaa"})
	// l := []any{
	// 	internals.H("ul", types.Props{}, []any{
	// 		internals.H("li", types.Props{}, []any{"one"}),
	// 		internals.H("li", types.Props{}, []any{"one"}),
	// 	}),
	// }

	// vdom.CreateNode(args[0])
	internals.MountDom(vdm, args[0])
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

func mountWithApp(this js.Value, args []js.Value) interface{} {
	var app internals.App
	app.State = "0"
	app.Reducers = internals.Reducers{
		"add": func(state any, value ...any) any {
			newState, _ := state.(int)
			val, _ := value[0].(int)
			return newState + val
		},
	}
	app.View = func(state any, emit func(eventName string, payload any)) internals.Vdom {
		return internals.H("button", types.Props{}, []any{app.State})
	}
	// emit := func(eventName string, payload any) {
	// 	fmt.Printf("Event emitted: %s with payload: %v\n", eventName, payload)
	// }
	// vdm := app.View(app.State, emit)
	// internals.MountDom(vdm, args[0])
	p := app.CreateApp()
	p["mount"](args[0])
	fmt.Println("wwwwwwwaaaaaaaawwwwwwwwaaaaa", app)
	return nil
}

func main() {

	fmt.Println("wasm connected")
	js.Global().Set("mountDom", js.FuncOf(mountDom))
	// js.Global().Set("removeDom", js.FuncOf(remDom))

	// js.Global().Set("mountDom", js.FuncOf(mountWithApp))
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
