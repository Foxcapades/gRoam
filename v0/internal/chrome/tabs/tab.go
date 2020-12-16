// +build js,wasm
// +build chrome

package tabs

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/sessions"
	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewTab(v js.Value) (out *Tab) {
	if v.IsUndefined() || v.IsNull() {
		return
	}

	out = new(Tab)

	out.active = v.Get(x.JsKeyActive).Bool()
	out.autoDiscardable = v.Get(x.JsKeyAutoDiscardable).Bool()
	out.discarded = v.Get(x.JsKeyDiscarded).Bool()
	out.groupID = chrome.GroupID(v.Get(x.JsKeyGroupID).Int())
	out.highlighted = v.Get(x.JsKeyHighlighted).Bool()
	out.index = uint16(v.Get(x.JsKeyIndex).Int())
	out.pinned = v.Get(x.JsKeyPinned).Bool()
	out.selected = v.Get(x.JsKeySelected).Bool()
	out.windowID = chrome.WindowID(v.Get(x.JsKeyWindowID).Int())

	mi := chrome.MutedInfo(NewMutedInfo(v.Get(x.JsKeyMutedInfo)))
	out.mutedInfo = &mi

	if val := v.Get(x.JsKeyAudible); !val.IsUndefined() && !val.IsNull() {
		tmp := val.Bool()
		out.audible = &tmp
	}

	if val := v.Get(x.JsKeyFavIconURL); !val.IsUndefined() && !val.IsNull() {
		tmp := val.String()
		out.favIconURL = &tmp
	}

	if val := v.Get(x.JsKeyHeight); !val.IsUndefined() && !val.IsNull() {
		tmp := uint16(val.Int())
		out.height = &tmp
	}

	if val := v.Get(x.JsKeyOpenerTabID); !val.IsUndefined() && !val.IsNull() {
		tmp := int32(val.Int())
		out.openerTabID = (*chrome.TabID)(&tmp)
	}

	if val := v.Get(x.JsKeyPendingURL); !val.IsUndefined() && !val.IsNull() {
		tmp := val.String()
		out.pendingURL = &tmp
	}

	if val := v.Get(x.JsKeySessionID); !val.IsUndefined() && !val.IsNull() {
		tmp := val.String()
		out.sessionID = (*chrome.SessionID)(&tmp)
	}

	if val := v.Get(x.JsKeyStatus); !val.IsUndefined() && !val.IsNull() {
		tmp := val.String()
		out.status = (*chrome.TabStatus)(&tmp)
	}

	if val := v.Get(x.JsKeyTitle); !val.IsUndefined() && !val.IsNull() {
		tmp := val.String()
		out.title = &tmp
	}

	if val := v.Get(x.JsKeyURL); !val.IsUndefined() && !val.IsNull() {
		tmp := val.String()
		out.url = &tmp
	}

	if val := v.Get(x.JsKeyWidth); !val.IsUndefined() && !val.IsNull() {
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
	return &x.OptionalBool{V: &t.audible}
}

func (t *Tab) AutoDiscardable() bool {
	return t.autoDiscardable
}

func (t *Tab) Discarded() bool {
	return t.discarded
}

func (t *Tab) FavIconURL() chrome.OptionalString {
	return &x.OptionalString{V: &t.favIconURL}
}

func (t *Tab) GroupID() chrome.GroupID {
	return t.groupID
}

func (t *Tab) Height() chrome.OptionalHeight {
	return &x.OptionalHeight{Value: &t.height}
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
	return &x.OptionalString{V: &t.pendingURL}
}

func (t *Tab) Pinned() bool {
	return t.pinned
}

func (t *Tab) Selected() bool {
	return t.selected
}

func (t *Tab) SessionID() chrome.OptionalSessionID {
	return &sessions.OptionalSessionID{V: &t.sessionID}
}

func (t *Tab) Status() chrome.OptionalTabStatus {
	return &OptionalTabStatus{&t.status}
}

func (t *Tab) Title() chrome.OptionalString {
	return &x.OptionalString{V: &t.title}
}

func (t *Tab) URL() chrome.OptionalString {
	return &x.OptionalString{V: &t.url}
}

func (t *Tab) Width() chrome.OptionalWidth {
	return &x.OptionalWidth{V: &t.width}
}

func (t *Tab) WindowID() chrome.WindowID {
	return t.windowID
}
