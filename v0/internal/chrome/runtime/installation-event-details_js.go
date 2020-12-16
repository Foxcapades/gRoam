// +build js,wasm
// +build chrome

package runtime

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewInstallationEventDetails(val js.Value) *InstallationEventDetails {
	out := new(InstallationEventDetails)
	out.reason = chrome.OnInstalledReason(val.Get(x.JsKeyReason).String())

	if v := val.Get(x.JsKeyID); !v.IsUndefined() && !v.IsNull() {
		tmp := v.String()
		out.id = &tmp
	}

	if v := val.Get(x.JsKeyPreviousVersion); !v.IsUndefined() && !v.IsNull() {
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
	return &x.OptionalString{V: &i.id}
}

func (i *InstallationEventDetails) PreviousVersion() chrome.OptionalString {
	return &x.OptionalString{V: &i.previousVersion}
}

func (i *InstallationEventDetails) Reason() chrome.OnInstalledReason {
	return i.reason
}
