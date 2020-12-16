package sessions

import "github.com/foxcapades/groam/v0/pkg/chrome"

type OptionalSessionID struct {
	value *chrome.SessionID
}

func (o *OptionalSessionID) IsPresent() bool {
	return o.value != nil
}

func (o *OptionalSessionID) IsAbsent() bool {
	return o.value == nil
}

func (o *OptionalSessionID) Clear() {
	o.value = nil
}

func (o *OptionalSessionID) Get() chrome.SessionID {
	return *o.value
}

func (o *OptionalSessionID) Set(id chrome.SessionID) {
	o.value = &id
}

func (o *OptionalSessionID) OrElse(id chrome.SessionID) chrome.SessionID {
	if o.value == nil {
		return id
	}

	return *o.value
}

func (o *OptionalSessionID) With(f func(chrome.SessionID)) {
	if o.value != nil {
		f(*o.value)
	}
}
