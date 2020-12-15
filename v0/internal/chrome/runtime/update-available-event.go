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
	out.Root = val
	out.Handler = js.FuncOf(out.handler)
	out.Listeners = make(map[interface{}]interface{})

	val.Call(jsFnAddListener, out.Event.Handler)

	return out
}

type UpdateAvailableEvent struct {
	x.Event
}

func (u *UpdateAvailableEvent) AddListener(key interface{}, cb chrome.UpdateAvailableEventListener) {
	u.Listeners[key] = cb
}

func (u *UpdateAvailableEvent) handler(_ js.Value, args []js.Value) interface{} {
	details := NewUpdateAvailableDetails(args[0])

	for _, fn := range u.Listeners {
		fn.(chrome.UpdateAvailableEventListener)(details)
	}

	return nil
}
