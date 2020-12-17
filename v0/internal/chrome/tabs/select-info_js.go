// +build js,wasm
// +build chrome

package tabs

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewSelectInfo(v js.Value) chrome.TabSelectInfo {
	out := new(SelectInfo)
	out.windowID = chrome.WindowID(v.Get(x.JsKeyWindowID).Int())

	return out
}
