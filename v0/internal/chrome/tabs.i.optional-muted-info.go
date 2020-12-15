package chrome

import (
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type OptionalMutedInfo struct {
	v **chrome.MutedInfo
}

func (o *OptionalMutedInfo) IsPresent() bool {
	return *o.v != nil
}

func (o *OptionalMutedInfo) IsAbsent() bool {
	return *o.v == nil
}

func (o *OptionalMutedInfo) Clear() {
	*o.v = nil
}

func (o *OptionalMutedInfo) Get() chrome.MutedInfo {
	return **o.v
}

func (o *OptionalMutedInfo) Set(info chrome.MutedInfo) {
	*o.v = &info
}

func (o *OptionalMutedInfo) OrElse(info chrome.MutedInfo) chrome.MutedInfo {
	if *o.v == nil {
		return info
	}

	return **o.v
}

func (o *OptionalMutedInfo) With(f func(chrome.MutedInfo)) {
	if *o.v != nil {
		f(**o.v)
	}
}
