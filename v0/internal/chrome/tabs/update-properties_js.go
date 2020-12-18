// +build js,wasm
// +build chrome

package tabs

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func UpdatePropertiesToJS(props chrome.TabUpdateProperties) js.Value {
	out := x.JsObject.New()

	x.PutOptionalBool(out, x.JsKeyActive, props.Active())
	x.PutOptionalBool(out, x.JsKeyAutoDiscardable, props.AutoDiscardable())
	x.PutOptionalBool(out, x.JsKeyHighlighted, props.Highlighted())
	x.PutOptionalBool(out, x.JsKeyMuted, props.Muted())
	PutOptionalTabID(out, x.JsKeyOpenerTabID, props.OpenerTabID())
	x.PutOptionalBool(out, x.JsKeyPinned, props.Pinned())
	x.PutOptionalBool(out, x.JsKeySelected, props.Selected())
	x.PutOptionalString(out, x.JsKeyURL, props.URL())

	return out
}
