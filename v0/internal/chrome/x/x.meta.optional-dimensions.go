package x

import "github.com/foxcapades/groam/v0/pkg/chrome"

type OptionalHeight struct {
	V **chrome.Height
}

func (o *OptionalHeight) IsPresent() bool {
	return *o.V != nil
}

func (o *OptionalHeight) IsAbsent() bool {
	return *o.V == nil
}

func (o *OptionalHeight) Clear() {
	*o.V = nil
}

func (o *OptionalHeight) Get() chrome.Height {
	return **o.V
}

func (o *OptionalHeight) Set(b chrome.Height) {
	*o.V = &b
}

func (o *OptionalHeight) OrElse(b chrome.Height) chrome.Height {
	if *o.V == nil {
		return b
	}

	return **o.V
}

func (o *OptionalHeight) With(f func(chrome.Height)) {
	if *o.V != nil {
		f(**o.V)
	}
}

type OptionalWidth struct {
	V **chrome.Width
}

func (o *OptionalWidth) IsPresent() bool {
	return *o.V != nil
}

func (o *OptionalWidth) IsAbsent() bool {
	return *o.V == nil
}

func (o *OptionalWidth) Clear() {
	*o.V = nil
}

func (o *OptionalWidth) Get() chrome.Width {
	return **o.V
}

func (o *OptionalWidth) Set(b chrome.Width) {
	*o.V = &b
}

func (o *OptionalWidth) OrElse(b chrome.Width) chrome.Width {
	if *o.V == nil {
		return b
	}

	return **o.V
}

func (o *OptionalWidth) With(f func(chrome.Width)) {
	if *o.V != nil {
		f(**o.V)
	}
}
