// +build js,wasm
// +build chrome

package tabs

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewActivatedInfo(v js.Value) chrome.TabActivatedInfo {
	out := new(ActivatedInfo)

	out.tabID = chrome.TabID(v.Get(x.JsKeyTabID).Int())
	out.windowID = chrome.WindowID(v.Get(x.JsKeyWindowID).Int())

	return out
}
