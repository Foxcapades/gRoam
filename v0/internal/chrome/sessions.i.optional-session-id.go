package chrome

import "github.com/foxcapades/groam/v0/pkg/chrome"

type OptionalSessionID struct {
	v **chrome.SessionID
}

func (o *OptionalSessionID) IsPresent() bool {
	return *o.v != nil
}

func (o *OptionalSessionID) IsAbsent() bool {
	return *o.v == nil
}

func (o *OptionalSessionID) Clear() {
	*o.v = nil
}

func (o *OptionalSessionID) Get() chrome.SessionID {
	return **o.v
}

func (o *OptionalSessionID) Set(id chrome.SessionID) {
	*o.v = &id
}

func (o *OptionalSessionID) OrElse(id chrome.SessionID) chrome.SessionID {
	if *o.v == nil {
		return id
	}

	return **o.v
}

func (o *OptionalSessionID) With(f func(chrome.SessionID)) {
	if *o.v != nil {
		f(**o.v)
	}
}
