package internals

import (
	"fmt"
	"go-fe-fwk/pkgs/utils"
	"go-fe-fwk/types"
	"syscall/js"
)

func DestroyDom(vdom any) {
	switch v := any(vdom).(type) {
	case types.StringDom:
		removeTextNode(v)
		v.El = js.Null()
	case types.Vdom:
		removeElementNode(v)
		v.El = js.Null()
	default:
		fmt.Println("can not destroy node of this type type")
		return
	}
}

func removeTextNode(vdom types.StringDom) {
	fmt.Println("aaaaaaaaaaaaaaaaa destroying", vdom)
	// element := vdom.El
	// element.Call("remove")
}

func removeElementNode(vdom types.Vdom) {
	fmt.Println("aaaaaaaaaaaa destroying el", vdom)
	element := vdom.El
	children := vdom.Children
	listeners := vdom.Listeners
	element.Set("className", "awaaww")
	// element.Call("remove")
	for _, child := range children {
		DestroyDom(child)
	}
	fmt.Println("Removing this shit")
	utils.RemoveEventListener(listeners, element)
	vdom.Listeners = nil
}

// const removeElementNode = (vdom)=>{
//     const {el,children, listeners} = vdom
//     el.remove()
//     children.forEach((child)=>{
//         destroyDom(child)
//     })
//     if(listeners){
//         removeEventListeners(listeners,el)
//         delete vdom.listeners
//     }
// }

// import {removeEventListeners} from "./events.mjs"
// import { DOM_TYPES } from "./h.mjs"

// export const destroyDom = (vdom) =>{
//     switch(vdom.type){
//         case DOM_TYPES.TEXT:
//             removeTextNode(vdom)
//             break;
//         case DOM_TYPES.FRAGMENT:
//             removeFragmentNode(vdom)
//             break;
//         case DOM_TYPES.ELEMENT:
//             removeElementNode(vdom)
//             break;
//         default:
//         throw new Error("cant destro node of type"+vdom.type)
//     }
//     delete vdom.el
// }

// const removeTextNode = (vdom)=>{
//     const {el} = vdom
//     el.remove()
// }

// const removeFragmentNode = (vdom) => {
//     const {children} = vdom
//     children.forEach((child)=>{
//         destroyDom(child)
//     })
// }
