// +build js,wasm
// +build chrome

package runtime

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewUpdateAvailableEvent(val js.Value) *UpdateAvailableEvent {
	out := new(UpdateAvailableEvent)
	out.root = val
	out.Event.handler = js.FuncOf(out.handler)
	out.list = make(map[interface{}]chrome.UpdateAvailableEventListener)

	val.Call(jsFnAddListener, out.Event.handler)

	return out
}

type UpdateAvailableEvent struct {
	x.Event

	list map[interface{}]chrome.UpdateAvailableEventListener
}

func (u *UpdateAvailableEvent) HasListener(key interface{}) bool {
	_, ok := u.list[key]
	return ok
}

func (u *UpdateAvailableEvent) HasListeners(keys []interface{}) bool {
	for i := range keys {
		if u.HasListener(keys[i]) {
			return true
		}
	}

	return false
}

func (u *UpdateAvailableEvent) RemoveListener(key interface{}) {
	delete(u.list, key)
}

func (u *UpdateAvailableEvent) AddListener(key interface{}, cb chrome.UpdateAvailableEventListener) {
	u.list[key] = cb
}

func (u *UpdateAvailableEvent) handler(_ js.Value, args []js.Value) interface{} {
	details := NewUpdateAvailableDetails(args[0])

	for _, fn := range u.list {
		fn(details)
	}

	return nil
}
