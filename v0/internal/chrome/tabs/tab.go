package tabs

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/sessions"
	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

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
	audible         chrome.OptionalBool
	favIconURL      chrome.OptionalString
	pendingURL      chrome.OptionalString
	title           chrome.OptionalString
	url             chrome.OptionalString
	height          chrome.OptionalHeight
	id              chrome.OptionalTabID
	openerTabID     chrome.OptionalTabID
	mutedInfo       chrome.OptionalMutedInfo
	sessionID       chrome.OptionalSessionID
	status          chrome.OptionalTabStatus
	width           chrome.OptionalWidth
}

func (t *Tab) Active() bool {
	return t.active
}

func (t *Tab) Audible() chrome.OptionalBool {
	return &x.OptionalBool{Value: &t.audible}
}

func (t *Tab) AutoDiscardable() bool {
	return t.autoDiscardable
}

func (t *Tab) Discarded() bool {
	return t.discarded
}

func (t *Tab) FavIconURL() chrome.OptionalString {
	return &x.OptionalString{value: &t.favIconURL}
}

func (t *Tab) GroupID() chrome.GroupID {
	return t.groupID
}

func (t *Tab) Height() chrome.OptionalHeight {
	return &x.OptionalHeight{value: &t.height}
}

func (t *Tab) Highlighted() bool {
	return t.highlighted
}

func (t *Tab) TabID() chrome.OptionalTabID {
	return &OptionalTabID{&t.id}
}

func (t *Tab) Incognito() bool {
	return t.incognito
}

func (t *Tab) Index() uint16 {
	return t.index
}

func (t *Tab) MutedInfo() chrome.OptionalMutedInfo {
	return &OptionalMutedInfo{&t.mutedInfo}
}

func (t *Tab) OpenerTabID() chrome.OptionalTabID {
	return &OptionalTabID{&t.openerTabID}
}

func (t *Tab) PendingURL() chrome.OptionalString {
	return &x.OptionalString{value: &t.pendingURL}
}

func (t *Tab) Pinned() bool {
	return t.pinned
}

func (t *Tab) Selected() bool {
	return t.selected
}

func (t *Tab) SessionID() chrome.OptionalSessionID {
	return &sessions.OptionalSessionID{Value: &t.sessionID}
}

func (t *Tab) Status() chrome.OptionalTabStatus {
	return &OptionalTabStatus{&t.status}
}

func (t *Tab) Title() chrome.OptionalString {
	return &x.OptionalString{value: &t.title}
}

func (t *Tab) URL() chrome.OptionalString {
	return &x.OptionalString{value: &t.url}
}

func (t *Tab) Width() chrome.OptionalWidth {
	return &x.OptionalWidth{value: &t.width}
}

func (t *Tab) WindowID() chrome.WindowID {
	return t.windowID
}
