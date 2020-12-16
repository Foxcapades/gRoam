package x

type OptionalF32 struct {
	value *float32
}

func (o *OptionalF32) IsPresent() bool {
	return o.value != nil
}

func (o *OptionalF32) IsAbsent() bool {
	return o.value == nil
}

func (o *OptionalF32) Clear() {
	o.value = nil
}

func (o *OptionalF32) Get() float32 {
	return *o.value
}

func (o *OptionalF32) Set(b float32) {
	o.value = &b
}

func (o *OptionalF32) OrElse(b float32) float32 {
	if o.value == nil {
		return b
	}

	return *o.value
}

func (o *OptionalF32) With(f func(float32)) {
	if o.value != nil {
		f(*o.value)
	}
}
