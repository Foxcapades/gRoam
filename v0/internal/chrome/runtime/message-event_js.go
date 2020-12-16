// +build js,wasm
// +build chrome

package runtime

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewMessageEvent(val js.Value) chrome.RuntimeMessageEvent {
	out := new(MessageEvent)
	out.Event = x.NewEvent(val, js.FuncOf(out.listener))

	return out
}

type MessageEvent struct {
	x.Event
}

func (r *MessageEvent) AddListener(key interface{}, cb chrome.RuntimeMessageEventListener) {
	r.Listeners[key] = cb
}

func (r *MessageEvent) listener(_ js.Value, args []js.Value) interface{} {
	sender := NewMessageSender(args[1])
	callback := func() { args[2].Invoke() }

	for _, fn := range r.Listeners {
		fn.(chrome.RuntimeMessageEventListener)(args[0], sender, callback)
	}

	return nil
}
