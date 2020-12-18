package tabs

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func SendMessageOptionsToJS(obj chrome.SendMessageOptions) js.Value {
	out := x.JsObject.New()

	if obj.FrameID().IsPresent() {
		out.Set(x.JsKeyFrameID, int(obj.FrameID().Get()))
	}

	return out
}
