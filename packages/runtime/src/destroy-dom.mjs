import {removeEventListeners} from "./events.mjs"
import { DOM_TYPES } from "./h.mjs"

export const destroyDom = (vdom) =>{
    switch(vdom.type){
        case DOM_TYPES.TEXT:
            removeTextNode(vdom)
            break;
        case DOM_TYPES.FRAGMENT:
            removeFragmentNode(vdom)
            break;
        case DOM_TYPES.ELEMENT:
            removeElementNode(vdom)
            break;
        default:
        throw new Error("cant destro node of type"+vdom.type)
    }
    delete vdom.el
}


const removeTextNode = (vdom)=>{
    const {el} = vdom
    el.remove()
}

const removeElementNode = (vdom)=>{
    const {el,children, listeners} = vdom
    el.remove()
    children.forEach((child)=>{
        destroyDom(child)
    })
    if(listeners){
        removeEventListeners(listeners,el)
        delete vdom.listeners
    }
}

const removeFragmentNode = (vdom) => {
    const {children} = vdom
    children.forEach((child)=>{
        destroyDom(child)
    })
}
