// +build js,wasm
// +build chrome

package sessions

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewOptionalSessionID(val js.Value) chrome.OptionalSessionID {
	out := new(OptionalSessionID)

	if !val.IsUndefined() && !val.IsNull() {
		tmp := chrome.SessionID(val.String())
		out.value = &tmp
	}

	return out
}
