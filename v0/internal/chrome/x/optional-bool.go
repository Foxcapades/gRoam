package x

type OptionalBool struct {
	Value *bool
}

func (o *OptionalBool) IsPresent() bool {
	return o.Value != nil
}

func (o *OptionalBool) IsAbsent() bool {
	return o.Value == nil
}

func (o *OptionalBool) Clear() {
	o.Value = nil
}

func (o *OptionalBool) Get() bool {
	return *o.Value
}

func (o *OptionalBool) Set(b bool) {
	o.Value = &b
}

func (o *OptionalBool) OrElse(b bool) bool {
	if o.Value == nil {
		return b
	}

	return *o.Value
}

func (o *OptionalBool) With(f func(bool)) {
	if o.Value != nil {
		f(*o.Value)
	}
}
