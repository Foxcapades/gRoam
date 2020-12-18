package tabs

import (
	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type ZoomChangeEvent struct {
	x.Event
}

func (z *ZoomChangeEvent) AddListener(key interface{}, cb chrome.ZoomChangeListener) {
	z.Listeners[key] = cb
}

type UpdateEvent struct {
	x.Event
}

func (u *UpdateEvent) AddListener(key interface{}, cb chrome.TabUpdateListener) {
	u.Listeners[key] = cb
}

type SelectEvent struct {
	x.Event
}

func (s *SelectEvent) AddListener(key interface{}, cb chrome.TabSelectionListener) {
	s.Listeners[key] = cb
}

type ReplaceEvent struct {
	x.Event
}

func (r *ReplaceEvent) AddListener(key interface{}, cb chrome.TabReplaceListener) {
	r.Listeners[key] = cb
}

type RemovalEvent struct {
	x.Event
}

func (r *RemovalEvent) AddListener(key interface{}, cb chrome.TabRemoveListener) {
	r.Listeners[key] = cb
}

type MoveEvent struct {
	x.Event
}

func (m *MoveEvent) AddListener(key interface{}, cb chrome.TabMoveListener) {
	m.Listeners[key] = cb
}

type HighlightEvent struct {
	x.Event
}

func (h *HighlightEvent) AddListener(key interface{}, cb chrome.TabHighlightListener) {
	h.Listeners[key] = cb
}

type DetachEvent struct {
	x.Event
}

func (d *DetachEvent) AddListener(key interface{}, cb chrome.TabDetachListener) {
	d.Listeners[key] = cb
}

type CreateEvent struct {
	x.Event
}

func (c *CreateEvent) AddListener(key interface{}, cb chrome.TabCreateListener) {
	c.Listeners[key] = cb
}

type AttachEvent struct {
	x.Event
}

func (a *AttachEvent) AddListener(key interface{}, cb chrome.TabAttachListener) {
	a.Listeners[key] = cb
}

type ActivatedEvent struct {
	x.Event
}

func (a *ActivatedEvent) AddListener(key interface{}, cb chrome.TabActivatedListener) {
	a.Listeners[key] = cb
}
