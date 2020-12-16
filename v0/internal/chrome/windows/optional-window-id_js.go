// +build js,wasm
// +build chrome

package windows

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewOptionalWindowID(val js.Value) chrome.OptionalWindowID {
	out := new(OptionalWindowID)

	if !val.IsUndefined() && !val.IsNull() {
		tmp := chrome.WindowType(val.String())
		out.value = &tmp
	}

	return out
}
