package chrome

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

const (
	jsKeyActive          = "active"
	jsKeyAudible         = "audible"
	jsKeyAutoDiscardable = "autoDiscardable"
	jsKeyDiscarded       = "discarded"
	jsKeyFavIconURL      = "favIconUrl"
	jsKeyGroupID         = "groupId"
	jsKeyHeight          = "height"
	jsKeyHighlighted     = "highlighted"
	jsKeyIndex           = "index"
	jsKeyMutedInfo       = "mutedInfo"
	jsKeyOpenerTabID     = "openerTabId"
	jsKeyPendingURL      = "pendingUrl"
	jsKeyPinned          = "pinned"
	jsKeySelected        = "selected"
	jsKeySessionID       = "sessionID"
	jsKeyStatus          = "status"
	jsKeyTitle           = "title"
	jsKeyWidth           = "width"
	jsKeyWindowID        = "windowId"
)

func NewTab(v js.Value) (out *Tab) {
	if v.IsUndefined() || v.IsNull() {
		return
	}

	out = new(Tab)

	out.active = v.Get(jsKeyActive).Bool()
	out.autoDiscardable = v.Get(jsKeyAutoDiscardable).Bool()
	out.discarded = v.Get(jsKeyDiscarded).Bool()
	out.groupID = chrome.GroupID(v.Get(jsKeyGroupID).Int())
	out.highlighted = v.Get(jsKeyHighlighted).Bool()
	out.index = uint16(v.Get(jsKeyIndex).Int())
	out.mutedInfo = NewMutedInfo(v.Get(jsKeyMutedInfo))
	out.pinned = v.Get(jsKeyPinned).Bool()
	out.selected = v.Get(jsKeySelected).Bool()
	out.windowID = chrome.WindowID(v.Get(jsKeyWindowID).Int())

	if val := v.Get(jsKeyAudible); !val.IsUndefined() && !val.IsNull() {
		tmp := val.Bool()
		out.audible = &tmp
	}

	if val := v.Get(jsKeyFavIconURL); !val.IsUndefined() && !val.IsNull() {
		tmp := val.String()
		out.favIconURL = &tmp
	}

	if val := v.Get(jsKeyHeight); !val.IsUndefined() && !val.IsNull() {
		tmp := uint16(val.Int())
		out.height = &tmp
	}

	if val := v.Get(jsKeyOpenerTabID); !val.IsUndefined() && !val.IsNull() {
		tmp := int32(val.Int())
		out.openerTabID = (*chrome.TabID)(&tmp)
	}

	if val := v.Get(jsKeyPendingURL); !val.IsUndefined() && !val.IsNull() {
		tmp := val.String()
		out.pendingURL = &tmp
	}

	if val := v.Get(jsKeySessionID); !val.IsUndefined() && !val.IsNull() {
		tmp := val.String()
		out.sessionID = (*chrome.SessionID)(&tmp)
	}

	if val := v.Get(jsKeyStatus); !val.IsUndefined() && !val.IsNull() {
		tmp := val.String()
		out.status = (*chrome.TabStatus)(&tmp)
	}

	if val := v.Get(jsKeyTitle); !val.IsUndefined() && !val.IsNull() {
		tmp := val.String()
		out.title = &tmp
	}

	if val := v.Get(jsKeyURL); !val.IsUndefined() && !val.IsNull() {
		tmp := val.String()
		out.url = &tmp
	}

	if val := v.Get(jsKeyWidth); !val.IsUndefined() && !val.IsNull() {
		tmp := uint16(val.Int())
		out.width = &tmp
	}

	return
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
	return &OptionalHeight{&t.height}
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
	return &OptionalString{&t.pendingURL}
}

func (t *Tab) Pinned() bool {
	return t.pinned
}

func (t *Tab) Selected() bool {
	return t.selected
}

func (t *Tab) SessionID() chrome.OptionalSessionID {
	return &OptionalSessionID{&t.sessionID}
}

func (t *Tab) Status() chrome.OptionalTabStatus {
	return &OptionalTabStatus{&t.status}
}

func (t *Tab) Title() chrome.OptionalString {
	return &OptionalString{&t.title}
}

func (t *Tab) URL() chrome.OptionalString {
	return &OptionalString{&t.url}
}

func (t *Tab) Width() chrome.OptionalWidth {
	return &OptionalWidth{&t.width}
}

func (t *Tab) WindowID() chrome.WindowID {
	return t.windowID
}
