package windows

import (
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type OptionalWindowState struct {
	Value *chrome.WindowState
}

func (o *OptionalWindowState) IsPresent() bool {
	return o.Value != nil
}

func (o *OptionalWindowState) IsAbsent() bool {
	return o.Value == nil
}

func (o *OptionalWindowState) Clear() {
	o.Value = nil
}

func (o *OptionalWindowState) Get() chrome.WindowState {
	return *o.Value
}

func (o *OptionalWindowState) Set(state chrome.WindowState) {
	o.Value = &state
}

func (o *OptionalWindowState) OrElse(state chrome.WindowState) chrome.WindowState {
	if o.IsPresent() {
		return *o.Value
	}

	return state
}

func (o *OptionalWindowState) With(f func(chrome.WindowState)) {
	if o.IsPresent() {
		f(o.Get())
	}
}
