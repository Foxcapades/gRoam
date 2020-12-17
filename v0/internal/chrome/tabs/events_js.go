// +build js,wasm
// +build chrome

package tabs

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewReplaceEvent(v js.Value) chrome.TabReplaceEvent {
	out := new(ReplaceEvent)
	out.Event = x.NewEvent(v, out.handle)

	return out
}

func (r *ReplaceEvent) handle(_ js.Value, args []js.Value) interface{} {
	added := chrome.TabID(args[0].Int())
	removed := chrome.TabID(args[1].Int())

	for _, fn := range s.Listeners {
		fn.(chrome.TabReplaceListener)(added, removed)
	}

	return nil
}

func NewSelectionEvent(v js.Value) chrome.TabSelectionEvent {
	out := new(SelectEvent)
	out.Event = x.NewEvent(v, out.handle)

	return out
}

func (s *SelectEvent) handle(_ js.Value, args []js.Value) interface{} {
	tabID := chrome.TabID(args[0].Int())
	info := NewSelectInfo(args[1])

	for _, fn := range s.Listeners {
		fn.(chrome.TabSelectionListener)(tabID, info)
	}

	return nil
}

func NewUpdateEvent(v js.Value) chrome.TabUpdateEvent {
	out := new(UpdateEvent)
	out.Event = x.NewEvent(v, out.handle)

	return out
}

func (u *UpdateEvent) handle(_ js.Value, args []js.Value) interface{} {
	tabID := chrome.TabID(args[0].Int())
	info := NewTabChangeInfo(args[1])

	for _, fn := range u.Listeners {
		fn.(chrome.TabUpdateListener)(tabID, info)
	}

	return nil
}

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
