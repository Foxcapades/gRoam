package chrome

type WindowID = int32

type OptionalWindowID interface {
	Optional

	Get() WindowID
	Set(WindowID)
	OrElse(WindowID)
	With(func(WindowID))
}

type (
	WindowCallback      = func(Window)
	MultiWindowCallback = func([]Window)
)

type Window interface {
	// Whether the window is set to be always on top.
	AlwaysOnTop() bool

	// Whether the window is currently the focused window.
	Focused() bool

	// The height of the window, including the frame, in pixels.  In some
	// circumstances a window may not be assigned a height property; for example,
	// when querying closed windows from the sessions API.
	Height() OptionalHeight

	// The ID of the window.  Window IDs are unique within a browser session.
	// In some circumstances a window may not be assigned an ID property; for
	// example, when querying windows using the sessions API, in which case a
	// session ID may be present.
	ID() OptionalWindowID

	// Whether the window is incognito.
	Incognito() bool

	// The offset of the window from the left edge of the screen in pixels.  In
	// some circumstances a window may not be assigned a left property; for
	// example, when querying closed windows from the sessions API.
	Left() OptionalI32

	// The session ID used to uniquely identify a window, obtained from the
	// sessions API.
	SessionID() OptionalSessionID

	// The state of this browser window.
	State() OptionalWindowState

	// Array of tabs.Tab objects representing the current tabs in the window.
	Tabs() OptionalTabSlice

	// The offset of the window from the top edge of the screen in pixels.  In
	// some circumstances a window may not be assigned a top property; for
	// example, when querying closed windows from the sessions API.
	Top() OptionalI32

	// The type of browser window this is.
	Type() OptionalWindowType

	// The width of the window, including the frame, in pixels.  In some
	// circumstances a window may not be assigned a width property; for example,
	// when querying closed windows from the sessions API.
	Width() OptionalWidth
}
