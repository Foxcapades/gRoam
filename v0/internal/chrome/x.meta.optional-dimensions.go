package chrome

import "github.com/foxcapades/groam/v0/pkg/chrome"

type OptionalHeight struct {
	value **chrome.Height
}

func (o *OptionalHeight) IsPresent() bool {
	return *o.value != nil
}

func (o *OptionalHeight) IsAbsent() bool {
	return *o.value == nil
}

func (o *OptionalHeight) Clear() {
	*o.value = nil
}

func (o *OptionalHeight) Get() chrome.Height {
	return **o.value
}

func (o *OptionalHeight) Set(b chrome.Height) {
	*o.value = &b
}

func (o *OptionalHeight) OrElse(b chrome.Height) chrome.Height {
	if *o.value == nil {
		return b
	}

	return **o.value
}

func (o *OptionalHeight) With(f func(chrome.Height)) {
	if *o.value != nil {
		f(**o.value)
	}
}

type OptionalWidth struct {
	value **chrome.Width
}

func (o *OptionalWidth) IsPresent() bool {
	return *o.value != nil
}

func (o *OptionalWidth) IsAbsent() bool {
	return *o.value == nil
}

func (o *OptionalWidth) Clear() {
	*o.value = nil
}

func (o *OptionalWidth) Get() chrome.Width {
	return **o.value
}

func (o *OptionalWidth) Set(b chrome.Width) {
	*o.value = &b
}

func (o *OptionalWidth) OrElse(b chrome.Width) chrome.Width {
	if *o.value == nil {
		return b
	}

	return **o.value
}

func (o *OptionalWidth) With(f func(chrome.Width)) {
	if *o.value != nil {
		f(**o.value)
	}
}
