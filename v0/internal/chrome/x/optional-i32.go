package x

type OptionalI32 struct {
	value *int32
}

func (o *OptionalI32) IsPresent() bool {
	return o.value != nil
}

func (o *OptionalI32) IsAbsent() bool {
	return o.value == nil
}

func (o *OptionalI32) Clear() {
	o.value = nil
}

func (o *OptionalI32) Get() int32 {
	return *o.value
}

func (o *OptionalI32) Set(b int32) {
	o.value = &b
}

func (o *OptionalI32) OrElse(b int32) int32 {
	if o.value == nil {
		return b
	}

	return *o.value
}

func (o *OptionalI32) With(f func(int32)) {
	if o.value != nil {
		f(*o.value)
	}
}
