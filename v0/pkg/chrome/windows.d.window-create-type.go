package chrome

type WindowCreateType string

const (
	WindowCreateTypeNormal WindowCreateType = "normal"
	WindowCreateTypePopup  WindowCreateType = "popup"
	WindowCreateTypePanel  WindowCreateType = "panel"
)

type OptionalWindowCreateType interface {
	Optional

	Get() WindowCreateType
	Set(WindowCreateType)
	OrElse(WindowCreateType)
	With(func(WindowCreateType))
}
