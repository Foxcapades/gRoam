package windows

import "github.com/foxcapades/groam/v0/pkg/chrome"

type OptionalWindowID struct {
	Value *chrome.WindowID
}

func (o *OptionalWindowID) IsPresent() bool {
	return o.Value != nil
}

func (o *OptionalWindowID) IsAbsent() bool {
	return o.Value == nil
}

func (o *OptionalWindowID) Clear() {
	o.Value = nil
}

func (o *OptionalWindowID) Get() chrome.WindowID {
	return *o.Value
}

func (o *OptionalWindowID) Set(id chrome.WindowID) {
	o.Value = &id
}

func (o *OptionalWindowID) OrElse(id chrome.WindowID) chrome.WindowID {
	if o.IsAbsent() {
		return id
	}

	return *o.Value
}

func (o *OptionalWindowID) With(f func(chrome.WindowID)) {
	if o.IsPresent() {
		f(*o.Value)
	}
}
