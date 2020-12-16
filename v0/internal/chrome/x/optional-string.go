package x

type OptionalString struct {
	value *string
}

func (o *OptionalString) IsPresent() bool {
	return o.value != nil
}

func (o *OptionalString) IsAbsent() bool {
	return o.value == nil
}

func (o *OptionalString) Clear() {
	o.value = nil
}

func (o *OptionalString) Get() string {
	return *o.value
}

func (o *OptionalString) Set(b string) {
	o.value = &b
}

func (o *OptionalString) OrElse(b string) string {
	if o.value == nil {
		return b
	}

	return *o.value
}

func (o *OptionalString) With(f func(string)) {
	if o.value != nil {
		f(*o.value)
	}
}
