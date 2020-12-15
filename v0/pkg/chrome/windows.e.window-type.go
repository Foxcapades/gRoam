package chrome

type WindowType string

const (
	WindowTypeNormal   WindowType = "normal"
	WindowTypePopup    WindowType = "popup"
	WindowTypePanel    WindowType = "panel"
	WindowTypeApp      WindowType = "app"
	WindowTypeDevTools WindowType = "devtools"
)

type OptionalWindowType interface {
	Optional

	Get() WindowType
	Set(WindowType)
	Add(WindowType)
	OrElse(WindowType)
	With(func(WindowType))
}

type OptionalWindowTypeSlice interface {
	Optional

	Get() []WindowType
	Set([]WindowType)
	Add(WindowType)
	OrElse([]WindowType)
	With(func([]WindowType))
}
