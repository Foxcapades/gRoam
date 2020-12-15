// +build js,wasm
// +build chrome

package runtime

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewPortMessageEvent() (out *PortMessageEvent) {
	out = new(PortMessageEvent)
	out.listeners = make(map[interface{}]chrome.PortMessageEventListener)
	out.root = js.FuncOf(out.listener)

	return
}

type PortMessageEvent struct {
	listeners map[interface{}]chrome.PortMessageEventListener
	root      js.Func
}

func (p *PortMessageEvent) AddListener(key interface{}, cb chrome.PortMessageEventListener) {
	p.listeners[key] = cb
}

func (p *PortMessageEvent) HasListener(key interface{}) bool {
	_, ok := p.listeners[key]
	return ok
}

func (p *PortMessageEvent) HasListeners(keys []interface{}) bool {
	for i := range keys {
		if _, ok := p.listeners[keys[i]]; ok {
			return true
		}
	}

	return false
}

func (p *PortMessageEvent) RemoveListener(key interface{}) {
	delete(p.listeners, key)
}

func (p *PortMessageEvent) Release() {
	p.root.Release()
}

func (p *PortMessageEvent) listener(_ js.Value, args []js.Value) interface{} {
	msg := Port{port: args[0]}
	tmp := Port{port: args[1]}

	for _, fn := range p.listeners {
		fn(msg, &tmp)
	}

	return nil
}
