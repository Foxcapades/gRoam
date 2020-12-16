// +build js,wasm
// +build chrome

package runtime

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func MessageOption2JS(in chrome.MessageOptions) js.Value {
	if in == nil {
		return js.Undefined()
	}

	out := x.JsObject.New()

	if in.IncludeTLSChannelID().IsPresent() {
		out.Set(x.JsKeyIncludeTLSChannelID, in.IncludeTLSChannelID().Get())
	}

	return out
}
