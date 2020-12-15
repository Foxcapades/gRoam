// +build js,wasm
// +build chrome

package runtime

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome"
)

func NewUpdateAvailableDetails(val js.Value) (out *UpdateAvailableDetails) {
	out = new(UpdateAvailableDetails)
	out.version = val.Get(chrome.jsKeyVersion).String()

	return out
}

type UpdateAvailableDetails struct {
	version string
}

func (u *UpdateAvailableDetails) Version() string {
	return u.version
}
