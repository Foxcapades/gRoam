package windows

import "github.com/foxcapades/groam/v0/pkg/chrome"

type OptionalWindowID struct {
	value *chrome.WindowID
}

func (o *OptionalWindowID) IsPresent() bool {
	return o.value != nil
}

func (o *OptionalWindowID) IsAbsent() bool {
	return o.value == nil
}

func (o *OptionalWindowID) Clear() {
	o.value = nil
}

func (o *OptionalWindowID) Get() chrome.WindowID {
	return *o.value
}

func (o *OptionalWindowID) Set(id chrome.WindowID) {
	o.value = &id
}

func (o *OptionalWindowID) OrElse(id chrome.WindowID) chrome.WindowID {
	if o.IsAbsent() {
		return id
	}

	return *o.value
}

func (o *OptionalWindowID) With(f func(chrome.WindowID)) {
	if o.IsPresent() {
		f(*o.value)
	}
}
