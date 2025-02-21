

const addEventListener = (eventName,handler,el) => {
    el.addEventListener(eventName,handler)
    return handler
}

export const addEventListeners = (listeners={},el)=>{
    const addedListeners = {}
    Object.entries(listeners).forEach(([eventName,handler])=>{
        const listener = addEventListener(eventName,handler,el)
        addedListeners[eventName] = listener
    })
    return addedListeners
}
