package chrome

type OptionalBool struct {
	value **bool
}

func (o *OptionalBool) IsPresent() bool {
	return *o.value != nil
}

func (o *OptionalBool) IsAbsent() bool {
	return *o.value == nil
}

func (o *OptionalBool) Clear() {
	*o.value = nil
}

func (o *OptionalBool) Get() bool {
	return **o.value
}

func (o *OptionalBool) Set(b bool) {
	*o.value = &b
}

func (o *OptionalBool) OrElse(b bool) bool {
	if *o.value == nil {
		return b
	}

	return **o.value
}

func (o *OptionalBool) With(f func(bool)) {
	if o.value != nil {
		f(**o.value)
	}
}
