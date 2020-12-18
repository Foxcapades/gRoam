// +build js,wasm
// +build chrome

package tabs

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func PutOptionalZoomSettingsScope(obj js.Value, key string, opt chrome.OptionalZoomSettingsScope) {
	if opt.IsPresent() {
		obj.Set(key, string(opt.Get()))
	}
}
