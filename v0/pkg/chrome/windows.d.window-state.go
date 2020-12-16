package chrome

type WindowState string

const (
	WindowStateNormal           WindowState = "normal"
	WindowStateMinimized        WindowState = "minimized"
	WindowStateMaximized        WindowState = "maximized"
	WindowStateFullscreen       WindowState = "fullscreen"
	WindowStateLockedFullscreen WindowState = "locked-fullscreen"
)

type OptionalWindowState interface {
	Optional

	Get() WindowState
	Set(WindowState)
	OrElse(WindowState) WindowState
	With(func(WindowState))
}
