package chrome

type (
	Height = uint16
	Width  = uint16
)

type OptionalHeight interface {
	Optional

	Get() Height
	Set(Height)
	OrElse(Height)
	With(func(Height))
}

type OptionalWidth interface {
	Optional

	Get() Width
	Set(Width)
	OrElse(Width)
	With(func(Width))
}
