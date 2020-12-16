package tabs

import (
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type MutedInfo struct {
	extID  chrome.OptionalString
	muted  bool
	reason chrome.OptionalMutedInfoReason
}

func (m *MutedInfo) ExtensionID() chrome.OptionalString {
	return m.extID
}

func (m *MutedInfo) Muted() bool {
	return m.muted
}

func (m *MutedInfo) Reason() chrome.OptionalMutedInfoReason {
	return m.reason
}
