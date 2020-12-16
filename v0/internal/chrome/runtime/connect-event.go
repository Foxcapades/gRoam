// +build js,wasm
// +build chrome

package runtime

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewConnectEvent(val js.Value) (out *ConnectEvent) {
	out = new(ConnectEvent)
	out.Event = x.NewEvent(val, out.listener)

	return out
}

type ConnectEvent struct {
	x.Event
}

func (c *ConnectEvent) AddListener(key interface{}, cb chrome.PortEventListener) {
	c.Listeners[key] = cb
}

func (c *ConnectEvent) listener(_ js.Value, args []js.Value) interface{} {
	port := NewPort(args[0])

	for _, fn := range c.Listeners {
		fn.(chrome.PortEventListener)(port)
	}

	return nil
}
