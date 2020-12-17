package tabs

import (
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type TabChangeInfo struct {
	audible         chrome.OptionalBool
	autoDiscardable chrome.OptionalBool
	discarded       chrome.OptionalBool
	favIconURL      chrome.OptionalString
	groupID         chrome.OptionalGroupID
	mutedInfo       chrome.OptionalMutedInfo
	pinned          chrome.OptionalBool
	status          chrome.OptionalTabStatus
	title           chrome.OptionalString
	url             chrome.OptionalString
}

func (t *TabChangeInfo) Audible() chrome.OptionalBool {
	return t.audible
}

func (t *TabChangeInfo) AutoDiscardable() chrome.OptionalBool {
	return t.autoDiscardable
}

func (t *TabChangeInfo) Discarded() chrome.OptionalBool {
	return t.discarded
}

func (t *TabChangeInfo) FavIconURL() chrome.OptionalString {
	return t.favIconURL
}

func (t *TabChangeInfo) GroupID() chrome.OptionalGroupID {
	return t.groupID
}

func (t *TabChangeInfo) MutedInfo() chrome.OptionalMutedInfo {
	return t.mutedInfo
}

func (t *TabChangeInfo) Pinned() chrome.OptionalBool {
	return t.pinned
}

func (t *TabChangeInfo) Status() chrome.OptionalTabStatus {
	return t.status
}

func (t *TabChangeInfo) Title() chrome.OptionalString {
	return t.title
}

func (t *TabChangeInfo) URL() chrome.OptionalString {
	return t.url
}
