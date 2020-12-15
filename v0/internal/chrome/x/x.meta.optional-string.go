package x

type OptionalString struct {
	V **string
}

func (o *OptionalString) IsPresent() bool {
	return *o.V != nil
}

func (o *OptionalString) IsAbsent() bool {
	return *o.V == nil
}

func (o *OptionalString) Clear() {
	*o.V = nil
}

func (o *OptionalString) Get() string {
	return **o.V
}

func (o *OptionalString) Set(b string) {
	*o.V = &b
}

func (o *OptionalString) OrElse(b string) string {
	if *o.V == nil {
		return b
	}

	return **o.V
}

func (o *OptionalString) With(f func(string)) {
	if *o.V != nil {
		f(**o.V)
	}
}
