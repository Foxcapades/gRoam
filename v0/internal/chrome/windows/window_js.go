// +build js,wasm
// +build chrome

package windows

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/sessions"
	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewWindow(val js.Value) chrome.Window {
	out := new(Window)

	out.alwaysOnTop = val.Get(x.JsKeyAlwaysOnTop).Bool()
	out.focused = val.Get(x.JsKeyFocused).Bool()
	out.height = x.NewOptionalHeight(val.Get(x.JsKeyHeight))
	out.id = NewOptionalWindowID(val.Get(x.JsKeyWindowID))
	out.incognito = val.Get(x.JsKeyIncognito).Bool()
	out.left = x.NewOptionalI32(val.Get(x.JsKeyLeft))
	out.sessionID = sessions.NewOptionalSessionID(val.Get(x.JsKeySessionID))
	out.state = NewOptionalWindowState(val.Get(x.JsKeyState))
	out.top = x.NewOptionalI32(val.Get(x.JsKeyTop))
	out.kind = NewOptionalWindowType(val.Get(x.JsKeyType))
	out.width = x.NewOptionalWidth(val.Get(x.JsKeyWidth))

	return out
}
