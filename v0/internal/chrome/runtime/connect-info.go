// +build js,wasm
// +build chrome

package runtime

import (
	"syscall/js"

	chrome2 "github.com/foxcapades/groam/v0/internal/chrome"
	"github.com/foxcapades/groam/v0/internal/chrome/x"
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
	return &x.OptionalBool{&c.includeTLS}
}

func (c *ConnectInfo) Name() chrome.OptionalString {
	return &x.OptionalString{&c.name}
}

func (c *ConnectInfo) JS() (out js.Value) {
	out = chrome2.jsObject.New()

	if c.includeTLS != nil {
		out.Set(jsKeyIncludeTLS, *c.includeTLS)
	}
	if c.name != nil {
		out.Set(jsKeyName, *c.name)
	}

	return
}
