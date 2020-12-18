// +build js,wasm
// +build chrome

package x

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewOptionalBool(val js.Value) chrome.OptionalBool {
	out := new(OptionalBool)

	if !val.IsUndefined() && !val.IsNull() {
		tmp := val.Bool()
		out.Value = &tmp
	}

	return out
}

func PutOptionalBool(obj js.Value, key string, opt chrome.OptionalBool) {
	if opt.IsPresent() {
		if opt.Get() {
			obj.Set(key, JsTrue)
		} else {
			obj.Set(key, JsFalse)
		}
	}
}
