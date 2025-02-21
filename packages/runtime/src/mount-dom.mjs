import { setAttributes } from './attributes.mjs';
import { addEventListeners } from './events.mjs';
import {DOM_TYPES} from './h'


export const mountDom = (vdom,parentEl)=>{
  switch (vdom.type) {
    case DOM_TYPES.TEXT:
      createTextNode(vdom,parentEl)
      break;

    case DOM_TYPES.ELEMENT:
      createElementNode(vdom,parentEl)
      break;
    case DOM_TYPES.FRAGMENT:
      createFragmentNode(vdom,parentEl)
      break

    default:
      break;
  }
}

const createTextNode = (vdom,parentEl)=>{
    const {value} = vdom
    const textNode = document.createTextNode(value)
    vdom.el = textNode
    parentEl.append(textNode)
}

const createFragmentNode = (vdom,parentEl)=>{
    const {children}=vdom
    vdom.el = parentEl
    children.forEach((child)=>mountDom(child,parentEl))
}

const createElementNode = (vdom,parentEl)=>{
    const {tag, props,children} =vdom
    const element = document.createElement(tag)
    addProps(element,props,vdom)
    vdom.el = element

    children.forEach((child)=>{
        mountDom(child,element)
    })
    parentEl.append(element)
}
const addProps = (el,props,vdom)=>{
    const {on:events,...attrs} = props
    vdom.listeners = addEventListeners(events,el)
    setAttributes(attrs,el)
}
