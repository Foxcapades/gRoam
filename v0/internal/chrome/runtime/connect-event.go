// +build js,wasm
// +build chrome

package runtime

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewConnectEvent(val js.Value) (out *ConnectEvent) {
	out = new(ConnectEvent)

	out.listeners = make(map[interface{}]chrome.PortEventListener)
	out.callback = js.FuncOf(out.listener)

	return out
}

type ConnectEvent struct {
	listeners map[interface{}]chrome.PortEventListener
	callback  js.Func
}

func (c *ConnectEvent) HasListener(key interface{}) bool {
	_, ok := c.listeners[key]
	return ok
}

func (c *ConnectEvent) HasListeners(keys []interface{}) bool {
	for i := range keys {
		if _, ok := c.listeners[keys[i]]; ok {
			return true
		}
	}

	return false
}

func (c *ConnectEvent) RemoveListener(key interface{}) {
	delete(c.listeners, key)
}

func (c *ConnectEvent) Release() {
	c.callback.Release()
}

func (c *ConnectEvent) AddListener(key interface{}, cb chrome.PortEventListener) {
	c.listeners[key] = cb
}

func (c *ConnectEvent) listener(_ js.Value, args []js.Value) interface{} {
	port := NewPort(args[0])

	for _, fn := range c.listeners {
		fn(port)
	}

	return nil
}
