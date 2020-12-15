package chrome

type Optional interface {
	IsPresent() bool
	IsAbsent() bool
	Clear()
}

type OptionalAny interface {
	Optional

	Get() interface{}
	Set(interface{})
	OrElse(interface{}) interface{}
	With(func(interface{}))
}

type OptionalBool interface {
	Optional

	Get() bool
	Set(bool)
	OrElse(bool) bool
	With(func(bool))
}

type OptionalF32 interface {
	Optional

	Get() float32
	Set(float32)
	OrElse(float32)
	With(func(float32))
}

type OptionalI32 interface {
	Optional

	Get() int32
	Set(int32)
	OrElse(int32) int32
	With(func(int32))
}

type OptionalString interface {
	Optional

	Get() string
	Set(string)
	OrElse(string) string
	With(func(string))
}

type OptionalStringSlice interface {
	Optional

	Get() []string
	Set([]string)
	Add(string)
	OrElse([]string) []string
	With(func([]string))
}

type OptionalU8 interface {
	Optional

	Get() uint8
	Set(uint8)
	OrElse(uint8) uint8
	With(func(uint8))
}

type OptionalU16 interface {
	Optional

	Get() uint16
	Set(uint16)
	OrElse(uint16) uint16
	With(func(uint16))
}
