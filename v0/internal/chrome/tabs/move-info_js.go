// +build js,wasm
// +build chrome

package tabs

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewMoveInfo(v js.Value) chrome.TabMoveInfo {
	out := new(MoveInfo)

	out.from = uint16(v.Get(x.JsKeyFrom).Int())
	out.to = uint16(v.Get(x.JsKeyTo).Int())
	out.windowID = chrome.WindowID(v.Get(x.JsKeyWindowID).Int())

	return out
}
