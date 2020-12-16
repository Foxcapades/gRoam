// +build js,wasm
// +build chrome

package tabs

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewMutedInfo(v js.Value) chrome.MutedInfo {
	if v.IsNull() || v.IsUndefined() {
		return nil
	}

	out := new(MutedInfo)

	out.extID = x.NewOptionalString(v.Get(x.JsKeyExtensionID))
	out.muted = v.Get(x.JsKeyMuted).Bool()
	out.reason = NewOptionalMutedInfoReason(v.Get(x.JsKeyReason))

	return out
}
