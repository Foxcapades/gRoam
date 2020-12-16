// +build js,wasm
// +build chrome

package x

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewOptionalString(val js.Value) chrome.OptionalString {
	out := new(OptionalString)

	if !val.IsUndefined() && !val.IsNull() {
		tmp := val.String()
		out.value = &tmp
	}

	return out
}
