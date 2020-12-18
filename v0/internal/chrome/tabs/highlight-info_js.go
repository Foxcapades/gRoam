// +build js,wasm
// +build chrome

package tabs

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewHighlightInfo(v js.Value) chrome.TabHighlightInfo {
	out := new(HighlightInfo)

	tmp := v.Get(x.JsKeyTabs)
	ln := tmp.Length()
	out.tabs = make([]chrome.TabID, ln)
	for i := 0; i < ln; i++ {
		out.tabs[i] = chrome.TabID(tmp.Index(i).Int())
	}

	out.windowID = chrome.WindowID(v.Get(x.JsKeyWindowID).Int())

	return out
}
