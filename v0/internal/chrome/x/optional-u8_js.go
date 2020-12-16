// +build js,wasm
// +build chrome

package x

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewOptionalU8(val js.Value) chrome.OptionalU8 {
	out := new(OptionalU8)

	if !val.IsUndefined() && !val.IsNull() {
		tmp := uint8(val.Int())
		out.value = &tmp
	}

	return out
}
