package x

type OptionalBool struct {
	V **bool
}

func (o *OptionalBool) IsPresent() bool {
	return *o.V != nil
}

func (o *OptionalBool) IsAbsent() bool {
	return *o.V == nil
}

func (o *OptionalBool) Clear() {
	*o.V = nil
}

func (o *OptionalBool) Get() bool {
	return **o.V
}

func (o *OptionalBool) Set(b bool) {
	*o.V = &b
}

func (o *OptionalBool) OrElse(b bool) bool {
	if *o.V == nil {
		return b
	}

	return **o.V
}

func (o *OptionalBool) With(f func(bool)) {
	if o.V != nil {
		f(**o.V)
	}
}
