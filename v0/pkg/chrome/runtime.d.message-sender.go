package chrome

// MessageSender
//
// An object containing information about the script context that sent a message
// or request.
type MessageSender interface {
	// FrameId
	//
	// The frame that opened the connection.
	//
	// 0 for top-level frames, positive for child frames.
	//
	// This will only be set when tab is set.
	FrameId() OptionalFrameID

	// GuestProcessId
	//
	// The guest process id of the requesting webview, if available.
	//
	// Only available for component extensions.
	GuestProcessId() OptionalI32

	// GuestRenderFrameRoutingId
	//
	// The guest render frame routing id of the requesting webview, if available.
	//
	// Only available for component extensions.
	GuestRenderFrameRoutingId() OptionalI32

	// Id
	//
	// The ID of the extension or app that opened the connection, if any.
	Id() OptionalString

	// NativeApplication
	//
	// The name of the native application that opened the connection, if any.
	NativeApplication() OptionalString

	// Origin
	//
	// The origin of the page or frame that opened the connection.
	//
	// It can vary from the url property (e.g., about:blank) or can be opaque
	// (e.g., sandboxed iframes). This is useful for identifying if the origin can
	// be trusted if we can't immediately tell from the URL.
	Origin() OptionalString

	// Tab
	//
	// The tabs.Tab which opened the connection, if any.
	//
	// This property will only be present when the connection was opened from a
	// tab (including content scripts), and only if the receiver is an extension,
	// not an app.
	Tab() OptionalTab

	// TlsChannelId
	//
	// The TLS channel ID of the page or frame that opened the connection, if
	// requested by the extension or app, and if available.
	TlsChannelId() OptionalString

	// Url
	//
	// The URL of the page or frame that opened the connection.
	//
	// If the sender is in an iframe, it will be iframe's URL not the URL of the
	// page which hosts it.
	Url() OptionalString
}

type OptionalMessageSender interface {
	Optional

	Get() MessageSender
	Set(MessageSender)
	OrElse(MessageSender) MessageSender
	With(func(MessageSender))
}