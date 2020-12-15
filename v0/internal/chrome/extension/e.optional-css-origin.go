package extension

import "github.com/foxcapades/groam/v0/pkg/chrome"

type OptionalCSSOrigin struct {
	v **chrome.CSSOrigin
}

func (o *OptionalCSSOrigin) IsPresent() bool {
	return *o.v != nil
}

func (o *OptionalCSSOrigin) IsAbsent() bool {
	return *o.v == nil
}

func (o *OptionalCSSOrigin) Clear() {
	*o.v = nil
}

func (o *OptionalCSSOrigin) Get() chrome.CSSOrigin {
	return **o.v
}

func (o *OptionalCSSOrigin) Set(c chrome.CSSOrigin) {
	*o.v = &c
}

func (o *OptionalCSSOrigin) OrElse(c chrome.CSSOrigin) chrome.CSSOrigin {
	if *o.v == nil {
		return c
	}

	return **o.v
}

func (o *OptionalCSSOrigin) With(f func(chrome.CSSOrigin)) {
	if o.v != nil {
		f(**o.v)
	}
}
