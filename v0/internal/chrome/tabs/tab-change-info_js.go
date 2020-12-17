package tabs

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewTabChangeInfo(v js.Value) chrome.TabChangeInfo {
	out := new(TabChangeInfo)

	out.audible = x.NewOptionalBool(v.Get(x.JsKeyAudible))
	out.autoDiscardable = x.NewOptionalBool(v.Get(x.JsKeyAutoDiscardable))
	out.discarded = x.NewOptionalBool(v.Get(x.JsKeyDiscarded))
	out.favIconURL = x.NewOptionalString(v.Get(x.JsKeyFavIconURL))
	out.groupID = NewOptionalGroupID(v.Get(x.JsKeyGroupID))
	out.mutedInfo = NewOptionalMutedInfo(v.Get(x.JsKeyMutedInfo))
	out.pinned = x.NewOptionalBool(v.Get(x.JsKeyPinned))
	out.status = NewOptionalTabStatus(v.Get(x.JsKeyStatus))
	out.title = x.NewOptionalString(v.Get(x.JsKeyTitle))
	out.url = x.NewOptionalString(v.Get(x.JsKeyURL))

	return out
}
