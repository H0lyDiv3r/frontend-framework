
 const DOM_TYPES={
    TEXT:'text',
    ELEMENT:'element',
    FRAGMENT: 'fragment'
}

 const h = (tag,props={},children=[])=>{
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
 const hString = (str) =>{
  return { type:DOM_TYPES.TEXT, value:str }
}

 const hFragment = (vNodes)=>{
  return {
    type: DOM_TYPES.FRAGMENT,
    children: mapTextNodes(withoutNulls(vNodes))
  }
}

  const withoutNulls = (arr)=>{
     return arr.filter((item)=>item!=null)
 }
   const mountDom = (vdom,parentEl)=>{
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
      console.log(vdom)
      const element = document.createElement(tag)
      addProps(element,props,vdom)
      vdom.el = element

      children.forEach((child)=>{
          mountDom(child,element)
      })
      parentEl.append(element)
      console.log(parentEl)
  }
  const addProps = (el,props,vdom)=>{
      const {on:events,...attrs} = props
      vdom.listeners = addEventListeners(events,el)
      setAttributes(attrs,el)
  }



  const addEventListener = (eventName,handler,el) => {
      el.addEventListener(eventName,handler)
      return handler
  }

   const addEventListeners = (listeners={},el)=>{
      const addedListeners = {}
      Object.entries(listeners).forEach(([eventName,handler])=>{
          const listener = addEventListener(eventName,handler,el)
          addedListeners[eventName] = listener
      })
      return addedListeners
  }

   const setAttributes = (el,attrs) =>{
      const { class:className,style,...otherAttrs} = attrs
      if(className){
          setClass(el,className)
      }
      if(style){
          Object.entries(style).forEach(([prop,value])=>{
              setStyle(el,prop,value)
          })
      }
      for(const [name,value] of Object.entries(otherAttrs)){
          setAttribute(el,name,value)
      }
  }

  const setClass = (el,className) => {
      el.className = ''
      if(typeof className === 'string') {
          el.className = className
      }
      if(Array.isArray(className)){
          el.classList.add(...className)
      }
  }

  const setStyle = (el,name,value)=>{
      el.style[name] = value
  }
  const removeStyle = (el,name)=>{
      el.style[name]=null
  }
  const setAttribute = (el,name,value)=>{
      if(value==null){
          removeAttribute(el,name)
      }else if(name.startsWith('data-')){
          el.setAttribute(name,value)
      }else{
          el[name]=value
      }
  }

  const removeAttribute = (el,name) =>{
      el[name]=null
      el.removeAttribute(name)
  }

  const vdom = h('section', {} ,[
  h('h1', {}, ['My Blog']),
  h('p', {}, ['Welcome to my blog!'])
  ])

  document.body.style.backgroundColor="red"
  console.log(document.getElementsByTagName('div'))
 mountDom(vdom,document.body)

 mountDom(h('div',{},[h('ul',{},[h('li',{},['one']),h('li',{},['two'])])]),document.body)
