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

	if v := val.Get(x.JsKeyHeight); !v.IsUndefined() && !v.IsNull() {
		tmp := uint16(v.Int())
		out.height = &x.OptionalHeight{Value: &tmp}
	}

	if v := val.Get(x.JsKeyWindowID); !v.IsUndefined() && !v.IsNull() {
		tmp := int16(v.Int())
		out.id = &OptionalWindowID{Value: (*chrome.WindowID)(&tmp)}
	}

	out.incognito = val.Get(x.JsKeyIncognito).Bool()

	if v := val.Get(x.JsKeyLeft); !v.IsUndefined() && !v.IsNull() {
		tmp := int32(v.Int())
		out.left = &x.OptionalI32{V: &tmp}
	}

	if v := val.Get(x.JsKeySessionID); !v.IsUndefined() && !v.IsNull() {
		tmp := v.String()
		out.sessionID = &sessions.OptionalSessionID{V: (*chrome.SessionID)(&tmp)}
	}

	if v := val.Get(x.JsKeyState); !v.IsUndefined() && !v.IsNull() {
		tmp := chrome.WindowState(v.String())
		out.state = &OptionalWindowState{Value: &tmp}
	}

	if v := val.Get(x.JsKeyTop); !v.IsUndefined() && !v.IsNull() {
		tmp := int32(v.Int())
		out.top = &x.OptionalI32{V: &tmp}
	}

	if v := val.Get(x.JsKeyType); !v.IsUndefined() && !v.IsNull() {
		tmp := chrome.WindowType(v.String())
		out.kind = &OptionalWindowType{Value: &tmp}
	}

	if v := val.Get(x.JsKeyWidth); !v.IsUndefined() && !v.IsNull() {
		tmp := uint16(v.Int())
		out.width = &x.OptionalWidth{V: &tmp}
	}

	return out
}
