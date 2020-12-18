// +build js,wasm
// +build chrome

package tabs

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func PutOptionalZoomSettingsMode(obj js.Value, key string, opt chrome.OptionalZoomSettingsMode) {
	if opt.IsPresent() {
		obj.Set(key, string(opt.Get()))
	}
}
