package tabs

import "github.com/foxcapades/groam/v0/pkg/chrome"

type OptionalTabID struct {
	value *chrome.TabID
}

func (o *OptionalTabID) IsPresent() bool {
	return o.value != nil
}

func (o *OptionalTabID) IsAbsent() bool {
	return o.value == nil
}

func (o *OptionalTabID) Clear() {
	o.value = nil
}

func (o *OptionalTabID) Get() chrome.TabID {
	return *o.value
}

func (o *OptionalTabID) Set(id chrome.TabID) {
	o.value = &id
}

func (o *OptionalTabID) OrElse(id chrome.TabID) chrome.TabID {
	if o.value == nil {
		return id
	}

	return *o.value
}

func (o *OptionalTabID) With(f func(chrome.TabID)) {
	if o.value != nil {
		f(*o.value)
	}
}
