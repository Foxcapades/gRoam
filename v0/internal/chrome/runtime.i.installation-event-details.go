package chrome

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

const (
	jsKeyPreviousVersion = "previousVersion"
)

func NewInstallationEventDetails(val js.Value) *InstallationEventDetails {
	out := new(InstallationEventDetails)
	out.reason = chrome.OnInstalledReason(val.Get(jsKeyReason).String())

	if v := val.Get(jsKeyID); !v.IsUndefined() && !v.IsNull() {
		tmp := v.String()
		out.id = &tmp
	}

	if v := val.Get(jsKeyPreviousVersion); !v.IsUndefined() && !v.IsNull() {
		tmp := v.String()
		out.id = &tmp
	}

	return out
}

type InstallationEventDetails struct {
	id              *string
	previousVersion *string
	reason          chrome.OnInstalledReason
}

func (i *InstallationEventDetails) ID() chrome.OptionalString {
	return &OptionalString{&i.id}
}

func (i *InstallationEventDetails) PreviousVersion() chrome.OptionalString {
	return &OptionalString{&i.previousVersion}
}

func (i *InstallationEventDetails) Reason() chrome.OnInstalledReason {
	return i.reason
}
