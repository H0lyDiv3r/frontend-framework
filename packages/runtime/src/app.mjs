import {mountDom} from "./mount-dom.mjs"
import {Dispatcher} from "./dispatcher.mjs"
import {destroyDom} from "./destroy-dom.mjs"

export const createApp = ({state,view,reducers={}})=>{
    let parentEl = null
    let vdom = null

    const emit = (eventName,payload)=>{
        dispatcher.dispatch(eventName, payload)
    }
    const renderApp = () => {
            if (vdom) {
                destroyDom(vdom)
            }

            vdom = view(state,emit)
        console.log("vdom", vdom);
            mountDom(vdom, parentEl)

        };

    const dispatcher = new Dispatcher()
    const subscriptions = [dispatcher.afterEveryCommand(renderApp)]

    for(const actionName in reducers){
        const reducer = reducers[actionName]
        const subs=dispatcher.subscribe(actionName, (payload)=>{
            state=reducer(state,payload)
        })
        subscriptions.push(subs)
    }


    return {
        mount(_parentEl){
            parentEl=_parentEl
            console.log("aaaaaaaaaaaaaaaaaaaaaaaa",_parentEl)
            renderApp()
        },
    unmount(){
        destroyDom(vdom),
        vdom=null,
        subscriptions.forEach((unsubscribe)=>unsubscribe())
    }
    }
}
