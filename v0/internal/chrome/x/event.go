// +build js,wasm
// +build chrome

package x

import (
	"syscall/js"
)

type Event struct {
	Root    js.Value
	Handler js.Func

	Listeners map[interface{}]interface{}
}

func (e *Event) HasListener(key interface{}) bool {
	_, ok := e.Listeners[key]
	return ok
}

func (e *Event) HasListeners(keys []interface{}) bool {
	for i := range keys {
		if e.HasListener(keys[i]) {
			return true
		}
	}

	return false
}

func (e *Event) RemoveListener(key interface{}) {
	delete(e.Listeners, key)
}

func (e *Event) Release() {
	e.Root.Call(JsFnRemoveListener, e.Handler)
	e.Handler.Release()
}
