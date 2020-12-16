// +build js,wasm
// +build chrome

package windows

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewOptionalWindowState(val js.Value) chrome.OptionalWindowState {
	out := new(OptionalWindowState)

	if !val.IsUndefined() && !val.IsNull() {
		tmp := chrome.WindowState(val.String())
		out.value = &tmp
	}

	return out
}
