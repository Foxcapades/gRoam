package chrome

import "github.com/foxcapades/groam/v0/pkg/chrome"

type OptionalFrameID struct {
	v **chrome.FrameID
}

func (o *OptionalFrameID) IsPresent() bool {
	return *o.v != nil
}

func (o *OptionalFrameID) IsAbsent() bool {
	return *o.v == nil
}

func (o *OptionalFrameID) Clear() {
	*o.v = nil
}

func (o *OptionalFrameID) Get() chrome.FrameID {
	return **o.v
}

func (o *OptionalFrameID) Set(id chrome.FrameID) {
	*o.v = &id
}

func (o *OptionalFrameID) OrElse(id chrome.FrameID) chrome.FrameID {
	if *o.v != nil {
		return **o.v
	}

	return id
}

func (o *OptionalFrameID) With(f func(chrome.FrameID)) {
	if *o.v != nil {
		f(**o.v)
	}
}

