package extension

import (
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type OptionalRunAt struct {
	v **chrome.RunAt
}

func (o *OptionalRunAt) IsPresent() bool {
	return *o.v != nil
}

func (o *OptionalRunAt) IsAbsent() bool {
	return *o.v == nil
}

func (o *OptionalRunAt) Clear() {
	*o.v = nil
}

func (o *OptionalRunAt) Get() chrome.RunAt {
	return **o.v
}

func (o *OptionalRunAt) Set(at chrome.RunAt) {
	*o.v = &at
}

func (o *OptionalRunAt) OrElse(at chrome.RunAt) chrome.RunAt {
	if *o.v == nil {
		return at
	}

	return **o.v
}

func (o *OptionalRunAt) With(f func(chrome.RunAt)) {
	if *o.v != nil {
		f(**o.v)
	}
}
