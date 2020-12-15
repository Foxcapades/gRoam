package chrome

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewTab(v js.Value) *Tab {

}

type Tab struct {
	active          bool
	autoDiscardable bool
	discarded       bool
	highlighted     bool
	incognito       bool
	pinned          bool
	selected        bool
	index           uint16
	groupID         chrome.GroupID
	windowID        chrome.WindowID
	audible         *bool
	favIconURL      *string
	pendingURL      *string
	title           *string
	url             *string
	height          *chrome.Height
	id              *chrome.TabID
	openerTabID     *chrome.TabID
	mutedInfo       *chrome.MutedInfo
	sessionID       *chrome.SessionID
	status          *chrome.TabStatus
	width           *chrome.Width
}

func (t *Tab) Active() bool {
	return t.active
}

func (t *Tab) Audible() chrome.OptionalBool {
	return &OptionalBool{&t.audible}
}

func (t *Tab) AutoDiscardable() bool {
	return t.autoDiscardable
}

func (t *Tab) Discarded() bool {
	return t.discarded
}

func (t *Tab) FavIconURL() chrome.OptionalString {
	return &OptionalString{&t.favIconURL}
}

func (t *Tab) GroupID() chrome.GroupID {
	return t.groupID
}

func (t *Tab) Height() chrome.OptionalHeight {
}

func (t *Tab) Highlighted() bool {
	return t.highlighted
}

func (t *Tab) TabID() chrome.OptionalTabID {
}

func (t *Tab) Incognito() bool {
	return t.incognito
}

func (t *Tab) Index() uint16 {
	return t.index
}

func (t *Tab) MutedInfo() chrome.OptionalMutedInfo {
}

func (t *Tab) OpenerTabID() chrome.OptionalTabID {
}

func (t *Tab) PendingURL() chrome.OptionalString {
	return &OptionalString{&t.pendingURL}
}

func (t *Tab) Pinned() bool {
	return t.pinned
}

func (t *Tab) Selected() bool {
	return t.selected
}

func (t *Tab) SessionID() chrome.OptionalSessionID {
}

func (t *Tab) Status() chrome.OptionalTabStatus {
}

func (t *Tab) Title() chrome.OptionalString {
	return &OptionalString{&t.title}
}

func (t *Tab) URL() chrome.OptionalString {
	return &OptionalString{&t.url}
}

func (t *Tab) Width() chrome.OptionalWidth {
}

func (t *Tab) WindowID() chrome.WindowID {
	return t.windowID
}
