package chrome

import "github.com/foxcapades/groam/v0/pkg/chrome"

type OptionalTabStatus struct {
	v **chrome.TabStatus
}

func (o *OptionalTabStatus) IsPresent() bool {
	return *o.v != nil
}

func (o *OptionalTabStatus) IsAbsent() bool {
	return *o.v == nil
}

func (o *OptionalTabStatus) Clear() {
	*o.v = nil
}

func (o *OptionalTabStatus) Get() chrome.TabStatus {
	return **o.v
}

func (o *OptionalTabStatus) Set(id chrome.TabStatus) {
	*o.v = &id
}

func (o *OptionalTabStatus) OrElse(id chrome.TabStatus) chrome.TabStatus {
	if *o.v == nil {
		return id
	}

	return **o.v
}

func (o *OptionalTabStatus) With(f func(chrome.TabStatus)) {
	if *o.v != nil {
		f(**o.v)
	}
}
