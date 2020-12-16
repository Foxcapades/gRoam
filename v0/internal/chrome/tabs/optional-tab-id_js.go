// +build js,wasm
// +build chrome

package tabs

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewOptionalTabID(val js.Value) chrome.OptionalTabID {
	out := new(OptionalTabID)

	if !val.IsUndefined() && !val.IsNull() {
		tmp := chrome.TabID(val.Int())
		out.value = &tmp
	}

	return out
}
