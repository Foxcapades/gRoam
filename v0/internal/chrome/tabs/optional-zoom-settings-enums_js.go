// +build js,wasm
// +build chrome

package tabs

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewOptionalZoomSettingsScope(val js.Value) chrome.OptionalZoomSettingsScope {
	out := new(OptionalZoomSettingsScope)

	if !val.IsUndefined() && !val.IsNull() {
		tmp := chrome.ZoomSettingsScope(val.String())
		out.value = &tmp
	}

	return out
}

// ZoomSettingsScopeToJSProp appends the value of the given option to the given
// JS object if the option contains a value.
func ZoomSettingsScopeToJSProp(obj js.Value, opt chrome.OptionalZoomSettingsScope) {
	if opt.IsPresent() {
		obj.Set(x.JsKeyScope, string(opt.Get()))
	}
}

func NewOptionalZoomSettingsMode(val js.Value) chrome.OptionalZoomSettingsMode {
	out := new(OptionalZoomSettingsMode)

	if !val.IsUndefined() && !val.IsNull() {
		tmp := chrome.ZoomSettingsMode(val.String())
		out.value = &tmp
	}

	return out
}

// ZoomSettingsModeToJSProp appends the value of the given option to the given
// JS object if the option contains a value.
func ZoomSettingsModeToJSProp(obj js.Value, opt chrome.OptionalZoomSettingsMode) {
	if opt.IsPresent() {
		obj.Set(x.JsKeyMode, string(opt.Get()))
	}
}
