package x

type OptionalU8 struct {
	value **uint8
}

func (o *OptionalU8) IsPresent() bool {
	return *o.value != nil
}

func (o *OptionalU8) IsAbsent() bool {
	return *o.value == nil
}

func (o *OptionalU8) Clear() {
	*o.value = nil
}

func (o *OptionalU8) Get() uint8 {
	return **o.value
}

func (o *OptionalU8) Set(b uint8) {
	*o.value = &b
}

func (o *OptionalU8) OrElse(b uint8) uint8 {
	if *o.value == nil {
		return b
	}

	return **o.value
}

func (o *OptionalU8) With(f func(uint8)) {
	if *o.value != nil {
		f(**o.value)
	}
}
