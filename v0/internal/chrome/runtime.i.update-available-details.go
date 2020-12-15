// +build js,wasm
// +build chrome

package chrome

import "syscall/js"

func NewUpdateAvailableDetails(val js.Value) (out *UpdateAvailableDetails) {
	out = new(UpdateAvailableDetails)
	out.version = val.Get(jsKeyVersion).String()

	return out
}

type UpdateAvailableDetails struct {
	version string
}

func (u *UpdateAvailableDetails) Version() string {
	return u.version
}
