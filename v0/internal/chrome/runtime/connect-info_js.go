// +build js,wasm
// +build chrome

package runtime

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewConnectInfo(js js.Value) chrome.ConnectInfo {
	if js.IsUndefined() || js.IsNull() {
		return nil
	}

	out := new(ConnectInfo)
	out.includeTLS = x.NewOptionalBool(js.Get(x.JsKeyIncludeTLSChannelID))
	out.name = x.NewOptionalString(js.Get(x.JsKeyName))

	return out
}

func SerializeConnectInfo(ci chrome.ConnectInfo) (out js.Value) {
	out = x.JsObject.New()

	if v := ci.IncludeTLSChannelID(); v.IsPresent() {
		out.Set(x.JsKeyIncludeTLSChannelID, v.Get())
	}

	if v := ci.Name(); v.IsPresent() {
		out.Set(x.JsKeyName, v.Get())
	}

	return

}
