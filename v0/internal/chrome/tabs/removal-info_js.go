// +build js,wasm
// +build chrome

package tabs

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewRemovalInfo(v js.Value) chrome.TabRemovalInfo {
	out := new(RemovalInfo)

	out.closing = v.Get(x.JsKeyIsWindowClosing).Bool()
	out.windowID = chrome.WindowID(v.Get(x.JsKeyWindowID).Int())

	return out
}
