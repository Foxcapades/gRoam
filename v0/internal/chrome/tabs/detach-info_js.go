// +build js,wasm
// +build chrome

package tabs

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewDetachInfo(v js.Value) chrome.TabDetachInfo {
	out := new(DetachInfo)

	out.position = uint16(v.Get(x.JsKeyOldPosition).Int())
	out.windowID = chrome.WindowID(v.Get(x.JsKeyWindowID).Int())

	return out
}

func NewAttachInfo(v js.Value) chrome.TabAttachInfo {
	out := new(DetachInfo)

	out.position = uint16(v.Get(x.JsKeyNewPosition).Int())
	out.windowID = chrome.WindowID(v.Get(x.JsKeyWindowID).Int())

	return out
}
