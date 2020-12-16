// +build js,wasm
// +build chrome

package x

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewOptionalF32(val js.Value) chrome.OptionalF32 {
	out := new(OptionalF32)

	if !val.IsUndefined() && !val.IsNull() {
		tmp := float32(val.Float())
		out.value = &tmp
	}

	return out
}
