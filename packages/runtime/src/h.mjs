import { withoutNulls } from "./utils/array.mjs"

export const DOM_TYPES={
    TEXT:'text',
    ELEMENT:'element',
    FRAGMENT: 'fragment'
}

export const h = (tag,props={},children=[])=>{
    return {
        tag,
        props,
        children: mapTextNodes(withoutNulls(children)),
        type:DOM_TYPES.ELEMENT
    }
}

const mapTextNodes = (children) =>{
  return children.map((child) => typeof child === 'string' ? hString(child) : child)
}
export const hString = (str) =>{
  return { type:DOM_TYPES.TEXT, value:str }
}

export const hFragment = (vNodes)=>{
  return {
    type: DOM_TYPES.FRAGMENT,
    children: mapTextNodes(withoutNulls(vNodes))
  }
}
