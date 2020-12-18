// +build js,wasm
// +build chrome

package tabs

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewActivatedEvent(v js.Value) chrome.TabActivatedEvent {
	out := new(ActivatedEvent)
	out.Event = x.NewEvent(v, out.handle)

	return out
}

func (a *ActivatedEvent) handle(_ js.Value, args []js.Value) interface{} {
	tabID := chrome.TabID(args[0].Int())
	info := NewAcInfo(args[1])

	for _, fn := range a.Listeners {
		fn.(chrome.TabActivatedListener)(tabID, info)
	}

	return nil
}

func NewAttachEvent(v js.Value) chrome.TabAttachEvent {
	out := new(AttachEvent)
	out.Event = x.NewEvent(v, out.handle)

	return out
}

func (a *AttachEvent) handle(_ js.Value, args []js.Value) interface{} {
	tabID := chrome.TabID(args[0].Int())
	info := NewAttachInfo(args[1])

	for _, fn := range a.Listeners {
		fn.(chrome.TabAttachListener)(tabID, info)
	}

	return nil
}

func NewCreateEvent(v js.Value) chrome.TabCreateEvent {
	out := new(CreateEvent)
	out.Event = x.NewEvent(v, out.handle)

	return out
}

func (c *CreateEvent) handle(_ js.Value, args []js.Value) interface{} {
	tab := NewTab(args[0])

	for _, fn := range c.Listeners {
		fn.(chrome.TabCreateListener)(tab)
	}

	return nil
}

func NewDetachEvent(v js.Value) chrome.TabDetachEvent {
	out := new(DetachEvent)
	out.Event = x.NewEvent(v, out.handle)

	return out
}

func (d *DetachEvent) handle(_ js.Value, args []js.Value) interface{} {
	tabID := chrome.TabID(args[0].Int())
	info := NewDetachInfo(args[1])

	for _, fn := range d.Listeners {
		fn.(chrome.TabDetachListener)(tabID, info)
	}

	return nil
}

func NewHighlightEvent(v js.Value) chrome.TabHighlightEvent {
	out := new(HighlightEvent)
	out.Event = x.NewEvent(v, out.handle)

	return out
}

func (h *HighlightEvent) handle(_ js.Value, args []js.Value) interface{} {
	info := NewHighlightInfo(args[0])

	for _, fn := range h.Listeners {
		fn.(chrome.TabHighlightListener)(info)
	}

	return nil
}

func NewMoveEvent(v js.Value) chrome.TabMoveEvent {
	out := new(MoveEvent)
	out.Event = x.NewEvent(v, out.handle)

	return out
}

func (m *MoveEvent) handle(_ js.Value, args []js.Value) interface{} {
	tabID := chrome.TabID(args[0].Int())
	info := NewMoveInfo(args[1])

	for _, fn := range m.Listeners {
		fn.(chrome.TabMoveListener)(tabID, info)
	}

	return nil
}

func NewRemoveEvent(v js.Value) chrome.TabRemovalEvent {
	out := new(RemovalEvent)
	out.Event = x.NewEvent(v, out.handle)

	return out
}

func (r *RemovalEvent) handle(_ js.Value, args []js.Value) interface{} {
	tabID := chrome.TabID(args[0].Int())
	info := NewRemovalInfo(args[1])

	for _, fn := range r.Listeners {
		fn.(chrome.TabRemoveListener)(tabID, info)
	}

	return nil
}

func NewReplaceEvent(v js.Value) chrome.TabReplaceEvent {
	out := new(ReplaceEvent)
	out.Event = x.NewEvent(v, out.handle)

	return out
}

func (r *ReplaceEvent) handle(_ js.Value, args []js.Value) interface{} {
	added := chrome.TabID(args[0].Int())
	removed := chrome.TabID(args[1].Int())

	for _, fn := range r.Listeners {
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
