// +build js,wasm
// +build chrome

package runtime

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewNotificationEvent(val js.Value) chrome.NotificationEvent {
	out := new(NotificationEvent)
	out.Root = val
	out.Listeners = make(map[interface{}]interface{})
	out.Handler = js.FuncOf(out.listen)

	val.Call(x.JsFnAddListener, out.Handler)

	return out
}

type NotificationEvent struct {
	x.Event
}

func (n *NotificationEvent) AddListener(key interface{}, cb chrome.EmptyCallback) {
	n.Listeners[key] = cb
}

func (n *NotificationEvent) listen(js.Value, []js.Value) interface{} {
	for _, fn := range n.Listeners {
		fn.(chrome.EmptyCallback)()
	}

	return nil
}
