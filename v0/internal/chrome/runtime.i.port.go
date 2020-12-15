// +build js,wasm
// +build chrome

package chrome

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

const (
	jsFnDisconnect  = "disconnect"
	jsFnPostMessage = "postMessage"
)

func NewPort(port js.Value) *Port {
	return &Port{port: port}
}

type Port struct {
	port js.Value

	pde chrome.PortDisconnectEvent
	pme chrome.PortMessageEvent

	sender chrome.MessageSender
}

func (p *Port) Disconnect() {
	p.port.Call(jsFnDisconnect)
}

func (p *Port) Name() string {
	return p.port.Get(jsKeyName).String()
}

func (p *Port) OnDisconnect() chrome.ConnectEvent {
	if p.pde == nil {
		p.pde = NewPortEvent()
	}

	return p.pde
}

func (p *Port) OnMessage() chrome.PortMessageEvent {
	if p.pme == nil {
		p.pme = NewPortMessageEvent()
	}

	return p.pme
}

// TODO: Check by hand that this returns nothing, unconvinced by the docs.
func (p *Port) PostMessage(any interface{}) {
	p.port.Call(jsFnPostMessage, any)
}

func (p *Port) Sender() chrome.OptionalMessageSender {

}

func (p *Port) Release() {
	if p.pde != nil {
		p.pde.Release()
	}
	if p.pme != nil {
		p.pme.Release()
	}
}
