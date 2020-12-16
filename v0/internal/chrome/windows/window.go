package windows

import (
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type Window struct {
	alwaysOnTop bool
	focused     bool
	height      chrome.OptionalHeight
	id          chrome.OptionalWindowID
	incognito   bool
	left        chrome.OptionalI32
	sessionID   chrome.OptionalSessionID
	state       chrome.OptionalWindowState
	tabs        chrome.OptionalTabSlice
	top         chrome.OptionalI32
	kind        chrome.OptionalWindowType
	width       chrome.OptionalWidth
}

func (w *Window) AlwaysOnTop() bool {
	return w.alwaysOnTop
}

func (w *Window) Focused() bool {
	return w.focused
}

func (w *Window) Height() chrome.OptionalHeight {
	return w.height
}

func (w *Window) ID() chrome.OptionalWindowID {
	return w.id
}

func (w *Window) Incognito() bool {
	return w.incognito
}

func (w *Window) Left() chrome.OptionalI32 {
	return w.left
}

func (w *Window) SessionID() chrome.OptionalSessionID {
	return w.sessionID
}

func (w *Window) State() chrome.OptionalWindowState {
	return w.state
}

func (w *Window) Tabs() chrome.OptionalTabSlice {
	return w.tabs
}

func (w *Window) Top() chrome.OptionalI32 {
	return w.top
}

func (w *Window) Type() chrome.OptionalWindowType {
	return w.kind
}

func (w *Window) Width() chrome.OptionalWidth {
	return w.width
}
