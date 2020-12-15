// +build js,wasm
// +build chrome

package chrome

import "syscall/js"

type Event struct {
	root    js.Value
	handler js.Func
}

func (e *Event) Release() {
	e.root.Call(jsFnRemoveListener, e.handler)
	e.handler.Release()
}
