package x

type OptionalI32 struct {
	V *int32
}

func (o *OptionalI32) IsPresent() bool {
	return o.V != nil
}

func (o *OptionalI32) IsAbsent() bool {
	return o.V == nil
}

func (o *OptionalI32) Clear() {
	o.V = nil
}

func (o *OptionalI32) Get() int32 {
	return *o.V
}

func (o *OptionalI32) Set(b int32) {
	o.V = &b
}

func (o *OptionalI32) OrElse(b int32) int32 {
	if o.V == nil {
		return b
	}

	return *o.V
}

func (o *OptionalI32) With(f func(int32)) {
	if o.V != nil {
		f(*o.V)
	}
}
