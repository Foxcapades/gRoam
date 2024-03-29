package chrome

type (
	Height = uint16
	Width  = uint16
)

type OptionalHeight interface {
	Optional

	Get() Height
	Set(Height)
	OrElse(Height) Height
	With(func(Height))
}

type OptionalWidth interface {
	Optional

	Get() Width
	Set(Width)
	OrElse(Width) Width
	With(func(Width))
}
