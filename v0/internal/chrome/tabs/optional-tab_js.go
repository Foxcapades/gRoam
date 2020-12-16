// +build js,wasm
// +build chrome

package tabs

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewOptionalTab(val js.Value) chrome.OptionalTab {
	out := new(OptionalTab)
	out.value = NewTab(val)
	return out
}
