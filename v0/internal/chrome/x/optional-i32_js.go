// +build js,wasm
// +build chrome

package x

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewOptionalI32(val js.Value) chrome.OptionalI32 {
	out := new(OptionalI32)

	if !val.IsUndefined() && !val.IsNull() {
		tmp := int32(val.Int())
		out.value = &tmp
	}

	return out
}
