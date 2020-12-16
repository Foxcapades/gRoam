package windows

import (
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type OptionalWindowType struct {
	value *chrome.WindowType
}

func (o *OptionalWindowType) IsPresent() bool {
	return o.value != nil
}

func (o *OptionalWindowType) IsAbsent() bool {
	return o.value == nil
}

func (o *OptionalWindowType) Clear() {
	o.value = nil
}

func (o *OptionalWindowType) Get() chrome.WindowType {
	return *o.value
}

func (o *OptionalWindowType) Set(windowType chrome.WindowType) {
	o.value = &windowType
}

func (o *OptionalWindowType) OrElse(windowType chrome.WindowType) chrome.WindowType {
	if o.IsPresent() {
		return o.Get()
	}

	return windowType
}

func (o *OptionalWindowType) With(f func(chrome.WindowType)) {
	if o.IsPresent() {
		f(o.Get())
	}
}
