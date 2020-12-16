// +build js,wasm
// +build chrome

package tabs

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewZoomChangeEvent(j js.Value) chrome.ZoomChangeEvent {
	out := new(ZoomChangeEvent)
	out.Event = x.NewEvent(j, out.handle)

	return out
}

func (z *ZoomChangeEvent) handle(_ js.Value, args []js.Value) interface{} {
	arg := NewZoomChangeInfo(args[0])

	for _, fn := range z.Listeners {
		fn.(chrome.ZoomChangeListener)(arg)
	}

	return nil
}
