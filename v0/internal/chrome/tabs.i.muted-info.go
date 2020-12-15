package chrome

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

const (
	jsKeyExtensionID = "extensionId"
	jsKeyMuted       = "muted"
	jsKeyReason      = "reason"
)

func NewMutedInfo(v js.Value) (out *MutedInfo) {
	if v.IsNull() || v.IsUndefined() {
		return
	}

	out = new(MutedInfo)

	out.muted = v.Get(jsKeyMuted).Bool()

	if val := v.Get(jsKeyExtensionID); !val.IsUndefined() && !val.IsNull() {
		tmp := val.String()
		out.extID = &tmp
	}

	if val := v.Get(jsKeyReason); !val.IsUndefined() && !val.IsNull() {
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
	return &OptionalString{&m.extID}
}

func (m *MutedInfo) Muted() bool {
	return m.muted
}

func (m *MutedInfo) Reason() chrome.OptionalMutedInfoReason {
	return &OptionalMutedInfoReason{&m.reason}
}
