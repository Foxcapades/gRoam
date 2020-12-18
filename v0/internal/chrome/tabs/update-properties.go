package tabs

import "github.com/foxcapades/groam/v0/pkg/chrome"

type UpdateProperties struct {
	active          chrome.OptionalBool
	autoDiscardable chrome.OptionalBool
	highlighted     chrome.OptionalBool
	muted           chrome.OptionalBool
	openerID        chrome.OptionalTabID
	pinned          chrome.OptionalBool
	selected        chrome.OptionalBool
	url             chrome.OptionalString
}

func (u *UpdateProperties) Active() chrome.OptionalBool {
	return u.active
}

func (u *UpdateProperties) SetActive(b bool) chrome.TabUpdateProperties {
	u.active.Set(b)
	return u
}

func (u *UpdateProperties) AutoDiscardable() chrome.OptionalBool {
	return u.autoDiscardable
}

func (u *UpdateProperties) SetAutoDiscardable(b bool) chrome.TabUpdateProperties {
	u.autoDiscardable.Set(b)
	return u
}

func (u *UpdateProperties) Highlighted() chrome.OptionalBool {
	return u.highlighted
}

func (u *UpdateProperties) SetHighlighted(b bool) chrome.TabUpdateProperties {
	u.highlighted.Set(b)
	return u
}

func (u *UpdateProperties) Muted() chrome.OptionalBool {
	return u.muted
}

func (u *UpdateProperties) SetMuted(b bool) chrome.TabUpdateProperties {
	u.muted.Set(b)
	return u
}

func (u *UpdateProperties) OpenerTabID() chrome.OptionalTabID {
	return u.openerID
}

func (u *UpdateProperties) SetOpener(id chrome.TabID) chrome.TabUpdateProperties {
	u.openerID.Set(id)
	return u
}

func (u *UpdateProperties) Pinned() chrome.OptionalBool {
	return u.pinned
}

func (u *UpdateProperties) SetPinned(b bool) chrome.TabUpdateProperties {
	u.pinned.Set(b)
	return u
}

func (u *UpdateProperties) Selected() chrome.OptionalBool {
	return u.selected
}

func (u *UpdateProperties) SetSelected(b bool) chrome.TabUpdateProperties {
	u.selected.Set(b)
	return u
}

func (u *UpdateProperties) URL() chrome.OptionalString {
	return u.url
}

func (u *UpdateProperties) SetURL(s string) chrome.TabUpdateProperties {
	u.url.Set(s)
	return u
}
