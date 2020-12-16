// +build js,wasm
// +build chrome

package x

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewOptionalHeight(val js.Value) chrome.OptionalHeight {
	out := new(OptionalHeight)

	if !val.IsUndefined() && !val.IsNull() {
		tmp := uint16(val.Int())
		out.value = &tmp
	}

	return out
}

func NewOptionalWidth(val js.Value) chrome.OptionalWidth {
	out := new(OptionalWidth)

	if !val.IsUndefined() && !val.IsNull() {
		tmp := uint16(val.Int())
		out.value = &tmp
	}

	return out
}
