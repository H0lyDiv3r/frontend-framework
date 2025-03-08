package internals

import (
	"fmt"
	"go-fe-fwk/types"
	"syscall/js"
)

type Reducers map[string]func(this js.Value, args []js.Value) any

type App struct {
	State    any                                                            `json:"state"`
	View     func(state any, emit func(eventName string, payload any)) Vdom `json:"view"`
	Reducers Reducers                                                       `json:"reducers"`
}

func (app *App) CreateApp() map[string]func(parentEl js.Value) {
	fmt.Println("CREATING THIS APP")
	var parentEl js.Value
	var vdom Vdom = nil

	var dispatcher Dispatcher = Dispatcher{Subs: map[string][]types.JsFunc{}, AfterHandlers: []func(payload ...any){}}
	emit := func(eventName string, payload any) {
		dispatcher.Dispatch(eventName, payload)
	}
	renderApp := func(payload ...any) {
		if vdom != nil {
			DestroyDom(vdom)
		}
		fmt.Println("RE RENDERING THE APP")
		vdom = app.View(app.State, emit)
		MountDom(vdom, parentEl)
	}
	afterHandler := dispatcher.AfterEveryCommand(renderApp)
	subscriptions := []func(){afterHandler}
	for actionName, reducer := range app.Reducers {
		subs := dispatcher.Subscribe(actionName, func(this js.Value, args []js.Value) any {
			args = append(args, js.ValueOf(app.State))
			app.State = reducer(js.Undefined(), args)
			return nil
		})
		subscriptions = append(subscriptions, subs)
		fmt.Println("subscribing the functions.", dispatcher.Subs)
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
