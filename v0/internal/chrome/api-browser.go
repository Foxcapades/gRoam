// +build js,wasm
// +build chrome

package chrome

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

const (
	jsKeyChrome  = "chrome"
	jsKeyRuntime = "runtime"
)

type Browser struct {
	chrome js.Value
}

func (b *Browser) Runtime() chrome.RuntimeAPI {
	return &RuntimeAPI{runtime: b.chrome.Get(jsKeyRuntime)}
}

func (b *Browser) Storage() chrome.StorageAPI {
	panic("implement me")
}

func (b *Browser) Tabs() chrome.TabAPI {
	panic("implement me")
}

func (b *Browser) Windows() chrome.WindowAPI {
	panic("implement me")
}

