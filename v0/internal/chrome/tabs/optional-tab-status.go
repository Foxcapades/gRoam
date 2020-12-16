package tabs

import "github.com/foxcapades/groam/v0/pkg/chrome"

type OptionalTabStatus struct {
	value *chrome.TabStatus
}

func (o *OptionalTabStatus) IsPresent() bool {
	return o.value != nil
}

func (o *OptionalTabStatus) IsAbsent() bool {
	return o.value == nil
}

func (o *OptionalTabStatus) Clear() {
	o.value = nil
}

func (o *OptionalTabStatus) Get() chrome.TabStatus {
	return *o.value
}

func (o *OptionalTabStatus) Set(id chrome.TabStatus) {
	o.value = &id
}

func (o *OptionalTabStatus) OrElse(id chrome.TabStatus) chrome.TabStatus {
	if o.value == nil {
		return id
	}

	return *o.value
}

func (o *OptionalTabStatus) With(f func(chrome.TabStatus)) {
	if o.value != nil {
		f(*o.value)
	}
}
