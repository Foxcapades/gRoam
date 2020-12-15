package chrome

type WindowCreateData interface {
	// If true, opens an active window. If false, opens an inactive window.
	Focused() OptionalBool
	SetFocused(bool) WindowCreateData

	// The height in pixels of the new window, including the frame.
	//
	// If not specified, defaults to a natural height.
	Height() OptionalHeight
	SetHeight(Height) WindowCreateData

	// Whether the new window should be an incognito window.
	Incognito() OptionalBool
	SetIncognito(bool) WindowCreateData

	// The number of pixels to position the new window from the left edge of the
	// screen.  If not specified, the new window is offset naturally from the last
	// focused window. This value is ignored for panels.
	Left() OptionalI32
	SetLeft(int32) WindowCreateData

	// If true, the newly-created window's 'window.opener' is set to the caller
	// and is in the same unit of related browsing contexts as the caller.
	SelfAsOpener() OptionalBool
	SetSelfAsOpener(bool) WindowCreateData

	// The initial state of the window. The minimized, maximized, and fullscreen
	// states cannot be combined with left, top, width, or height.
	State() OptionalWindowState
	SetState(WindowState) WindowCreateData

	// The ID of the tab to add to the new window.
	TabID() OptionalTabID
	SetTabID(TabID) WindowCreateData

	// The number of pixels to position the new window from the top edge of the
	// screen. If not specified, the new window is offset naturally from the last
	// focused window. This value is ignored for panels.
	Top() OptionalI32
	SetTop(int32)

	// Specifies what type of browser window to create.
	CreateType() OptionalWindowCreateType
	SetCreateType(WindowCreateType) WindowCreateData

	// A URL or array of URLs to open as tabs in the window.  Fully-qualified URLs
	// must include a scheme, e.g., 'http://www.google.com', not 'www.google.com'.
	//
	// Non-fully-qualified URLs are considered relative within the extension.
	//
	// Defaults to the New Tab Page.
	URL() OptionalString
	SetURL(string) WindowCreateData
	AddURL(string) WindowCreateData

	// The width in pixels of the new window, including the frame.
	//
	// If not specified, defaults to a natural width.
	Width() OptionalWidth
}
