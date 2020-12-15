package tabs

import "github.com/foxcapades/groam/v0/pkg/chrome"

type OptionalTabID struct {
	v **chrome.TabID
}

func (o *OptionalTabID) IsPresent() bool {
	return *o.v != nil
}

func (o *OptionalTabID) IsAbsent() bool {
	return *o.v == nil
}

func (o *OptionalTabID) Clear() {
	*o.v = nil
}

func (o *OptionalTabID) Get() chrome.TabID {
	return **o.v
}

func (o *OptionalTabID) Set(id chrome.TabID) {
	*o.v = &id
}

func (o *OptionalTabID) OrElse(id chrome.TabID) chrome.TabID {
	if *o.v == nil {
		return id
	}

	return **o.v
}

func (o *OptionalTabID) With(f func(chrome.TabID)) {
	if *o.v != nil {
		f(**o.v)
	}
}
