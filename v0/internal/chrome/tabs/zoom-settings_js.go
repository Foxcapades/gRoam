// +build js,wasm
// +build chrome

package tabs

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewZoomSettings(val js.Value) chrome.ZoomSettings {
	out := new(ZoomSettings)
	out.defZoomFactor = x.NewOptionalF32(val.Get(x.JsKeyDefaultZoomFactor))
	out.mode = NewOptionalZoomSettingsMode(val.Get(x.JsKeyMode))
	out.scope = NewOptionalZoomSettingsScope(val.Get(x.JsKeyScope))

	return out
}

func NewZoomChangeInfo(val js.Value) chrome.ZoomChangeInfo {
	out := new(ZoomChangeInfo)
	out.oldFactor = float32(val.Get(x.JsKeyOldZoomFactor).Float())
	out.newFactor = float32(val.Get(x.JsKeyNewZoomFactor).Float())
	out.tabID = chrome.TabID(val.Get(x.JsKeyTabID).Int())
	out.settings = NewZoomSettings(val.Get(x.JsKeyZoomSettings))

	return out
}

func ZoomSettingsToJS(zoom chrome.ZoomSettings) js.Value {
	out := x.JsObject.New()

	x.PutOptionalF32(out, x.JsKeyDefaultZoomFactor, zoom.DefaultZoomFactor())
	PutOptionalZoomSettingsMode(out, x.JsKeyMode, zoom.Mode())
	PutOptionalZoomSettingsScope(out, x.JsKeyScope, zoom.Scope())

	return out
}
