// +build js,wasm
// +build chrome

package x

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewOptionalBool(val js.Value) chrome.OptionalBool {
	out := new(OptionalBool)

	if !val.IsUndefined() && !val.IsNull() {
		tmp := val.Bool()
		out.Value = &tmp
	}

	return out
}
