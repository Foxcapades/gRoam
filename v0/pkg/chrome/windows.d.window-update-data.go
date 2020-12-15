package chrome

type WindowUpdateData interface {
	// If true, causes the window to be displayed in a manner that draws the
	// user's attention to the window, without changing the focused window.
	//
	// The effect lasts until the user changes focus to the window.
	//
	// This option has no effect if the window already has focus.
	//
	// Set to false to cancel a previous drawAttention request.
	DrawAttention() OptionalBool
	SetDrawAttention(bool) WindowUpdateData

	// If true, brings the window to the front; cannot be combined with the state
	// 'minimized'. If false, brings the next window in the z-order to the front;
	// cannot be combined with the state 'fullscreen' or 'maximized'.
	Focused() OptionalBool
	SetFocused(bool) WindowUpdateData

	// The height to resize the window to in pixels. This value is ignored for
	// panels.
	Height() OptionalHeight
	SetHeight(Height) WindowUpdateData

	// The offset from the left edge of the screen to move the window to in
	// pixels. This value is ignored for panels.
	Left() OptionalI32
	SetLeft(int32) WindowUpdateData

	// The new state of the window. The 'minimized', 'maximized', and 'fullscreen'
	// states cannot be combined with 'left', 'top', 'width', or 'height'.
	State() OptionalWindowState
	SetState(WindowState) WindowUpdateData

	// The offset from the top edge of the screen to move the window to in pixels.
	// This value is ignored for panels.
	Top() OptionalI32
	SetTop(int32) WindowUpdateData

	// The width to resize the window to in pixels. This value is ignored for
	// panels.
	Width() OptionalWidth
	SetWidth(Width) WindowUpdateData
}
