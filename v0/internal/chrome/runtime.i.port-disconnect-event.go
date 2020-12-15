// +build js,wasm
// +build chrome

package chrome

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewPortEvent() (out *PortEvent) {
	out = new(PortEvent)
	out.listeners = make(map[interface{}]chrome.PortEventListener)
	out.root = js.FuncOf(out.listener)

	return
}

type PortEvent struct {
	listeners map[interface{}]chrome.PortEventListener
	root js.Func
}

func (p *PortEvent) HasListener(key interface{}) bool {
	_, ok := p.listeners[key]
	return ok
}

func (p *PortEvent) HasListeners(keys []interface{}) bool {
	for i := range keys {
		if _, ok := p.listeners[keys[i]]; ok {
			return true
		}
	}

	return false
}

func (p *PortEvent) RemoveListener(key interface{}) {
	delete(p.listeners, key)
}

func (p *PortEvent) AddListener(key interface{}, cb chrome.PortEventListener) {
	p.listeners[key] = cb
}

func (p *PortEvent) Release() {
	p.root.Release()
}

func (p *PortEvent) listener(_ js.Value, args []js.Value) interface{} {
	tmp := Port{port: args[0]}

	for _, fn := range p.listeners {
		fn(&tmp)
	}

	return nil
}
