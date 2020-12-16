// +build js,wasm
// +build chrome

package tabs

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewOptionalMutedInfo(val js.Value) chrome.OptionalMutedInfo {
	out := new(OptionalMutedInfo)
	out.value = NewMutedInfo(val)
	return out
}
