package windows

import (
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type OptionalWindowState struct {
	value *chrome.WindowState
}

func (o *OptionalWindowState) IsPresent() bool {
	return o.value != nil
}

func (o *OptionalWindowState) IsAbsent() bool {
	return o.value == nil
}

func (o *OptionalWindowState) Clear() {
	o.value = nil
}

func (o *OptionalWindowState) Get() chrome.WindowState {
	return *o.value
}

func (o *OptionalWindowState) Set(state chrome.WindowState) {
	o.value = &state
}

func (o *OptionalWindowState) OrElse(state chrome.WindowState) chrome.WindowState {
	if o.IsPresent() {
		return *o.value
	}

	return state
}

func (o *OptionalWindowState) With(f func(chrome.WindowState)) {
	if o.IsPresent() {
		f(o.Get())
	}
}
