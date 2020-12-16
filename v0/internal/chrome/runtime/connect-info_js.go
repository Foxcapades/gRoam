// +build js,wasm
// +build chrome

package runtime

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/x"
)

func (c *ConnectInfo) JS() (out js.Value) {
	out = x.JsObject.New()

	if c.includeTLS != nil {
		out.Set(x.JsKeyIncludeTLSChannelID, *c.includeTLS)
	}

	if c.name != nil {
		out.Set(x.JsKeyName, *c.name)
	}

	return
}
