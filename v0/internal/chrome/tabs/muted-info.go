// +build js,wasm
// +build chrome

package tabs

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewMutedInfo(v js.Value) (out *MutedInfo) {
	if v.IsNull() || v.IsUndefined() {
		return
	}

	out = new(MutedInfo)

	out.muted = v.Get(x.JsKeyMuted).Bool()

	if val := v.Get(x.JsKeyExtensionID); !val.IsUndefined() && !val.IsNull() {
		tmp := val.String()
		out.extID = &tmp
	}

	if val := v.Get(x.JsKeyReason); !val.IsUndefined() && !val.IsNull() {
		tmp := val.String()
		out.reason = (*chrome.MutedInfoReason)(&tmp)
	}

	return
}

type MutedInfo struct {
	extID  *string
	muted  bool
	reason *chrome.MutedInfoReason
}

func (m *MutedInfo) ExtensionID() chrome.OptionalString {
	return &x.OptionalString{V: &m.extID}
}

func (m *MutedInfo) Muted() bool {
	return m.muted
}

func (m *MutedInfo) Reason() chrome.OptionalMutedInfoReason {
	return &OptionalMutedInfoReason{&m.reason}
}
