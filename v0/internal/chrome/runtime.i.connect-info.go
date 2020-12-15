// +build js,wasm
// +build chrome

package chrome

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

const (
	jsKeyIncludeTLS = "includeTlsChannelId"
	jsKeyName       = "name"
)

type ConnectInfo struct {
	includeTLS *bool
	name       *string
}

func (c *ConnectInfo) IncludeTLSChannelID() chrome.OptionalBool {
	return &OptionalBool{&c.includeTLS}
}

func (c *ConnectInfo) Name() chrome.OptionalString {
	return &OptionalString{&c.name}
}

func (c *ConnectInfo) JS() (out js.Value) {
	out = jsObject.New()

	if c.includeTLS != nil {
		out.Set(jsKeyIncludeTLS, *c.includeTLS)
	}
	if c.name != nil {
		out.Set(jsKeyName, *c.name)
	}

	return
}
