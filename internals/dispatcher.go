package internals

import (
	"fmt"
	"go-fe-fwk/pkgs/utils"
)

type Dispatcher struct {
	Subs          map[string][]func(payload ...any) `json:"subs"`
	AfterHandlers []func(payload ...any)            `json:"afterHandlers"`
}

func (dispatcher *Dispatcher) Subscribe(commandName string, handler func(payload ...any)) func() {
	if _, ok := dispatcher.Subs[commandName]; !ok {
		dispatcher.Subs[commandName] = []func(payload ...any){}
	}
	handlers := dispatcher.Subs[commandName]
	if utils.FuncExists(handlers, handler) {
		return func() {}
	}
	handlers = append(handlers, handler)
	return func() {
		idx := utils.IndexOfFunction(handlers, handler)
		handlers = append(handlers[:idx], handlers[idx+1:]...)
	}
}

func (dispatcher *Dispatcher) AfterEveryCommand(handler func(payload ...any)) func() {
	dispatcher.AfterHandlers = append(dispatcher.AfterHandlers, handler)
	return func() {
		idx := utils.IndexOfFunction(dispatcher.AfterHandlers, handler)
		dispatcher.AfterHandlers = append(dispatcher.AfterHandlers[:idx], dispatcher.AfterHandlers[idx+1:]...)
	}
}

func (dispatcher *Dispatcher) Dispatch(commandName string, payload any) {
	if _, ok := dispatcher.Subs[commandName]; ok {
		for _, handler := range dispatcher.Subs[commandName] {
			handler(payload)
		}
	} else {
		fmt.Println("there is no command by that name")
	}
	for _, handler := range dispatcher.AfterHandlers {
		handler()
	}
}

//     dispatch(commandName, payload){
//         if(this.#subs.has(commandName)){
//             this.#subs.get(commandName).forEach((handler)=>{
//                 handler(payload)
//             })
//         }else{
//             console.log(`there is no command by the name ${commandName}`)
//         }

//         this.#afterHandlers.forEach((handler)=>{handler()})
//     }
// export class Dispatcher {
//     #subs = new Map()
//     #afterHandlers =[]
//     subscribe(commandName, handler){
//         if(!this.#subs.has(commandName)){
//             this.#subs.set(commandName, [])
//         }
//         const handlers = this.#subs.get(commandName)
//         if(handlers.includes(handler)){
//             return ()=>{}
//         }
//         handlers.push(handler)
//         return ()=>{
//             const idx = handlers.indexOf(handler)
//             handlers.splice(idx,1)
//         }
//     }
//     afterEveryCommand(handler) {
//         this.#afterHandlers.push(handler)
//         return () =>{
//             const idx= this.#afterHandlers.indexOf(handler)
//             this.#afterHandlers.splice(idx,1)
//         }
//     }

//     dispatch(commandName, payload){
//         if(this.#subs.has(commandName)){
//             this.#subs.get(commandName).forEach((handler)=>{
//                 handler(payload)
//             })
//         }else{
//             console.log(`there is no command by the name ${commandName}`)
//         }

//         this.#afterHandlers.forEach((handler)=>{handler()})
//     }

// }
