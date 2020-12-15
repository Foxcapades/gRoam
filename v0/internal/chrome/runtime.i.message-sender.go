package chrome

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

const (
	jsKeyFrameID        = "frameId"
	jsKeyGuestProcessID = "guestProcessId"
	jsKeyGuestRoutingID = "guestRoutingId"
	jsKeyNativeApp      = "nativeApplication"
	jsKeyOrigin         = "origin"
	jsKeyTab            = "tab"
	jsKeyTLSChanID      = "tlsChannelId"
	jsKeyURL            = "url"
)

func NewMessageSender(v js.Value) (out *MessageSender) {
	if v.IsUndefined() || v.IsNull() {
		return nil
	}

	out = new(MessageSender)

	if fid := v.Get(jsKeyFrameID); !fid.IsUndefined() && !fid.IsNull() {
		val := int32(fid.Int())
		out.frameID = (*chrome.FrameID)(&val)
	}

	if gid := v.Get(jsKeyGuestProcessID); !gid.IsUndefined() && !gid.IsNull() {
		val := int32(gid.Int())
		out.guestProcessID = &val
	}

	if gid := v.Get(jsKeyGuestRoutingID); !gid.IsUndefined() && !gid.IsNull() {
		val := int32(gid.Int())
		out.guestRoutingID = &val
	}

	if str := v.Get(jsKeyID); !str.IsUndefined() && !str.IsNull() {
		val := str.String()
		out.id = &val
	}

	if str := v.Get(jsKeyNativeApp); !str.IsUndefined() && !str.IsNull() {
		val := str.String()
		out.nativeApp = &val
	}

	if str := v.Get(jsKeyOrigin); !str.IsUndefined() && !str.IsNull() {
		val := str.String()
		out.origin = &val
	}

	if tab := v.Get(jsKeyTab); !tab.IsUndefined() && !tab.IsNull() {
		tmp := chrome.Tab(NewTab(tab))
		out.tab = &tmp
	}

	if str := v.Get(jsKeyTLSChanID); !str.IsUndefined() && !str.IsNull() {
		val := str.String()
		out.tlsChannelID = &val
	}

	if str := v.Get(jsKeyURL); !str.IsUndefined() && !str.IsNull() {
		val := str.String()
		out.url = &val
	}

	return
}

type MessageSender struct {
	frameID        *chrome.FrameID
	guestProcessID *int32
	guestRoutingID *int32
	id             *string
	nativeApp      *string
	origin         *string
	tab            *chrome.Tab
	tlsChannelID   *string
	url            *string
}

func (m *MessageSender) FrameId() chrome.OptionalFrameID {
	return &OptionalFrameID{&m.frameID}
}

func (m *MessageSender) GuestProcessId() chrome.OptionalI32 {
	return &OptionalI32{&m.guestProcessID}
}

func (m *MessageSender) GuestRenderFrameRoutingId() chrome.OptionalI32 {
	return &OptionalI32{&m.guestRoutingID}
}

func (m *MessageSender) Id() chrome.OptionalString {
	return &OptionalString{&m.id}
}

func (m *MessageSender) NativeApplication() chrome.OptionalString {
	return &OptionalString{&m.nativeApp}
}

func (m *MessageSender) Origin() chrome.OptionalString {
	return &OptionalString{&m.origin}
}

func (m *MessageSender) Tab() chrome.OptionalTab {
	return &OptionalTab{&m.tab}
}

func (m *MessageSender) TlsChannelId() chrome.OptionalString {
	return &OptionalString{&m.tlsChannelID}
}

func (m *MessageSender) Url() chrome.OptionalString {
	return &OptionalString{&m.url}
}
