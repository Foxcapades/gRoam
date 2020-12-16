package sessions

import "github.com/foxcapades/groam/v0/pkg/chrome"

type OptionalSessionID struct {
	V *chrome.SessionID
}

func (o *OptionalSessionID) IsPresent() bool {
	return o.V != nil
}

func (o *OptionalSessionID) IsAbsent() bool {
	return o.V == nil
}

func (o *OptionalSessionID) Clear() {
	o.V = nil
}

func (o *OptionalSessionID) Get() chrome.SessionID {
	return *o.V
}

func (o *OptionalSessionID) Set(id chrome.SessionID) {
	o.V = &id
}

func (o *OptionalSessionID) OrElse(id chrome.SessionID) chrome.SessionID {
	if o.V == nil {
		return id
	}

	return *o.V
}

func (o *OptionalSessionID) With(f func(chrome.SessionID)) {
	if o.V != nil {
		f(*o.V)
	}
}
