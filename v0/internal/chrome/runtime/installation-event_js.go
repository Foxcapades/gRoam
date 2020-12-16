// +build js,wasm
// +build chrome

package runtime

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

const (
	jsFnAddListener    = "addListener"
	jsFnRemoveListener = "removeListener"
)

func NewInstallationEvent(val js.Value) *InstallationEvent {
	out := new(InstallationEvent)
	out.root = val
	out.listeners = make(map[interface{}]chrome.InstallationEventListener)
	out.listener = js.FuncOf(out.jsFunc)

	val.Call(jsFnAddListener, out.listener)

	return out
}

type InstallationEvent struct {
	root      js.Value
	listeners map[interface{}]chrome.InstallationEventListener
	listener  js.Func
}

func (i *InstallationEvent) HasListener(key interface{}) bool {
	_, ok := i.listeners[key]
	return ok
}

func (i *InstallationEvent) HasListeners(keys []interface{}) bool {
	for j := range keys {
		if i.HasListener(keys[j]) {
			return true
		}
	}

	return false
}

func (i *InstallationEvent) RemoveListener(key interface{}) {
	delete(i.listeners, key)
}

func (i *InstallationEvent) Release() {
	i.listener.Release()
	i.root.Call(jsFnRemoveListener, i.listener)
}

func (i *InstallationEvent) AddListener(key interface{}, cb chrome.InstallationEventListener) {
	i.listeners[key] = cb
}

func (i InstallationEvent) jsFunc(this js.Value, args []js.Value) interface{} {
	details := NewInstallationEventDetails(args[0])

	for _, fn := range i.listeners {
		fn(details)
	}

	return nil
}
