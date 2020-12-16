// +build js,wasm
// +build chrome

package tabs

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewOptionalMutedInfoReason(val js.Value) chrome.OptionalMutedInfoReason {
	out := new(OptionalMutedInfoReason)

	if !val.IsUndefined() && !val.IsNull() {
		tmp := chrome.MutedInfoReason(val.String())
		out.value = &tmp
	}

	return out
}
