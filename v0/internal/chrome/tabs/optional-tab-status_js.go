// +build js,wasm
// +build chrome

package tabs

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewOptionalTabStatus(val js.Value) chrome.OptionalTabStatus {
	out := new(OptionalTabStatus)

	if !val.IsUndefined() && !val.IsNull() {
		tmp := chrome.TabStatus(val.String())
		out.value = &tmp
	}

	return out
}
