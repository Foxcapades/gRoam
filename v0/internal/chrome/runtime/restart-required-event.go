// +build js,wasm
// +build chrome

package runtime

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewRestartRequiredEvent(val js.Value) chrome.RestartRequiredEvent {
	out := new(RestartRequiredEvent)
	out.Root = val
	out.Listeners = make(map[interface{}]interface{})
	out.Handler = js.FuncOf(out.listener)

	val.Call(x.JsFnAddListener, out.Handler)

	return out
}

type RestartRequiredEvent struct {
	x.Event
}

func (r *RestartRequiredEvent) AddListener(key interface{}, cb chrome.RestartRequiredEventListener) {
	r.Listeners[key] = cb
}

func (r *RestartRequiredEvent) listener(_ js.Value, args []js.Value) interface{} {
	reason := chrome.OnRestartReason(args[0].String())

	for _, fn := range r.Listeners {
		fn.(chrome.RestartRequiredEventListener)(reason)
	}

	return nil
}
