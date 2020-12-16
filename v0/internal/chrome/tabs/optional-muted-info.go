package tabs

import (
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type OptionalMutedInfo struct {
	value chrome.MutedInfo
}

func (o *OptionalMutedInfo) IsPresent() bool {
	return o.value != nil
}

func (o *OptionalMutedInfo) IsAbsent() bool {
	return o.value == nil
}

func (o *OptionalMutedInfo) Clear() {
	o.value = nil
}

func (o *OptionalMutedInfo) Get() chrome.MutedInfo {
	return o.value
}

func (o *OptionalMutedInfo) Set(info chrome.MutedInfo) {
	o.value = info
}

func (o *OptionalMutedInfo) OrElse(info chrome.MutedInfo) chrome.MutedInfo {
	if o.value == nil {
		return info
	}

	return o.value
}

func (o *OptionalMutedInfo) With(f func(chrome.MutedInfo)) {
	if o.value != nil {
		f(o.value)
	}
}
