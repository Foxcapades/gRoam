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
	jsKeyPreviousVersion = "previousVersion"
)

func NewInstallationEventDetails(val js.Value) *InstallationEventDetails {
	out := new(InstallationEventDetails)
	out.reason = chrome.OnInstalledReason(val.Get(chrome2.jsKeyReason).String())

	if v := val.Get(chrome2.jsKeyID); !v.IsUndefined() && !v.IsNull() {
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
	return &x.OptionalString{&i.id}
}

func (i *InstallationEventDetails) PreviousVersion() chrome.OptionalString {
	return &x.OptionalString{&i.previousVersion}
}

func (i *InstallationEventDetails) Reason() chrome.OnInstalledReason {
	return i.reason
}
