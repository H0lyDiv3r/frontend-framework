package internals

import (
	"fmt"
	"syscall/js"
)

type Reducers map[string]func(state any, payload ...any) any

type App struct {
	State    any                                                            `json:"state"`
	View     func(state any, emit func(eventName string, payload any)) Vdom `json:"view"`
	Reducers Reducers                                                       `json:"reducers"`
}

func (this *App) CreateApp() map[string]func(parentEl js.Value) {
	var parentEl js.Value
	var vdom Vdom = nil

	var dispatcher Dispatcher = Dispatcher{Subs: map[string][]func(payload ...any){}}
	emit := func(eventName string, payload any) {
		dispatcher.Dispatch(eventName, payload)
	}
	renderApp := func(payload ...any) {
		if vdom != nil {
			DestroyDom(vdom)
		}
		vdom = this.View(this.State, emit)
		MountDom(vdom, parentEl)
	}
	subscriptions := []func(){dispatcher.AfterEveryCommand(renderApp)}
	for actionName, reducer := range this.Reducers {
		subs := dispatcher.Subscribe(actionName, func(payload ...any) {
			this.State = reducer(this.State, payload[0])
		})
		subscriptions = append(subscriptions, subs)
	}

	return map[string]func(payload js.Value){
		"mount": func(_parentEl js.Value) {
			parentEl = _parentEl
			renderApp()
			fmt.Println("mouting")
		},
		"unmount": func(_parentEl js.Value) {
			fmt.Println("dismounting")
		},
	}

}

//     return {
//         mount(_parentEl){
//             parentEl=_parentEl
//             console.log("aaaaaaaaaaaaaaaaaaaaaaaa",_parentEl)
//             renderApp()
//         },
//     unmount(){
//         destroyDom(vdom),
//         vdom=null,
//         subscriptions.forEach((unsubscribe)=>unsubscribe())
//     }
//     }

// export const createApp = ({state,view,reducers={}})=>{
//     let parentEl = null
//     let vdom = null
//     const emit = (eventName,payload)=>{
//         dispatcher.dispatch(eventName, payload)
//     }
//     const renderApp = () => {
//             if (vdom) {
//                 destroyDom(vdom)
//             }
//             vdom = view(state,emit)
//         console.log("vdom", vdom);
//             mountDom(vdom, parentEl)
//         };
//     const dispatcher = new Dispatcher()
//     const subscriptions = [dispatcher.afterEveryCommand(renderApp)]

//     for(const actionName in reducers){
//         const reducer = reducers[actionName]
//         const subs=dispatcher.subscribe(actionName, (payload)=>{
//             state=reducer(state,payload)
//         })
//         subscriptions.push(subs)
//     }

//     return {
//         mount(_parentEl){
//             parentEl=_parentEl
//             console.log("aaaaaaaaaaaaaaaaaaaaaaaa",_parentEl)
//             renderApp()
//         },
//     unmount(){
//         destroyDom(vdom),
//         vdom=null,
//         subscriptions.forEach((unsubscribe)=>unsubscribe())
//     }
//     }
// }
