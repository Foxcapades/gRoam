// +build js,wasm
// +build chrome

package tabs

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/sessions"
	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewTab(v js.Value) chrome.Tab {
	if v.IsUndefined() || v.IsNull() {
		return nil
	}

	out := new(Tab)

	out.active = v.Get(x.JsKeyActive).Bool()
	out.audible = x.NewOptionalBool(v.Get(x.JsKeyAudible))
	out.autoDiscardable = v.Get(x.JsKeyAutoDiscardable).Bool()
	out.discarded = v.Get(x.JsKeyDiscarded).Bool()
	out.favIconURL = x.NewOptionalString(v.Get(x.JsKeyFavIconURL))
	out.groupID = chrome.GroupID(v.Get(x.JsKeyGroupID).Int())
	out.height = x.NewOptionalHeight(v.Get(x.JsKeyHeight))
	out.highlighted = v.Get(x.JsKeyHighlighted).Bool()
	out.id = NewOptionalTabID(v.Get(x.JsKeyID))
	out.incognito = v.Get(x.JsKeyIncognito).Bool()
	out.index = uint16(v.Get(x.JsKeyIndex).Int())
	out.mutedInfo = NewOptionalMutedInfo(v.Get(x.JsKeyMutedInfo))
	out.openerTabID = NewOptionalTabID(v.Get(x.JsKeyOpenerTabID))
	out.pendingURL = x.NewOptionalString(v.Get(x.JsKeyPendingURL))
	out.pinned = v.Get(x.JsKeyPinned).Bool()
	out.selected = v.Get(x.JsKeySelected).Bool()
	out.sessionID = sessions.NewOptionalSessionID(v.Get(x.JsKeySessionID))
	out.status = NewOptionalTabStatus(v.Get(x.JsKeyStatus))
	out.title = x.NewOptionalString(v.Get(x.JsKeyTitle))
	out.url = x.NewOptionalString(v.Get(x.JsKeyURL))
	out.windowID = chrome.WindowID(v.Get(x.JsKeyWindowID).Int())
	out.width = x.NewOptionalWidth(v.Get(x.JsKeyWidth))

	return out
}
