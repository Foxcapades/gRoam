package tabs

import "github.com/foxcapades/groam/v0/pkg/chrome"

type OptionalGroupID struct {
	value *chrome.GroupID
}

func (o *OptionalGroupID) IsPresent() bool {
	return o.value != nil
}

func (o *OptionalGroupID) IsAbsent() bool {
	return o.value == nil
}

func (o *OptionalGroupID) Clear() {
	o.value = nil
}

func (o *OptionalGroupID) Get() chrome.GroupID {
	return *o.value
}

func (o *OptionalGroupID) Set(id chrome.GroupID) {
	o.value = &id
}

func (o *OptionalGroupID) OrElse(id chrome.GroupID) chrome.GroupID {
	if o.value == nil {
		return id
	}

	return *o.value
}

func (o *OptionalGroupID) With(f func(chrome.GroupID)) {
	if o.value != nil {
		f(*o.value)
	}
}
