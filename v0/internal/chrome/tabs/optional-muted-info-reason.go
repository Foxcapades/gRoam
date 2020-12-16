package tabs

import (
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type OptionalMutedInfoReason struct {
	value *chrome.MutedInfoReason
}

func (o *OptionalMutedInfoReason) IsPresent() bool {
	return o.value != nil
}

func (o *OptionalMutedInfoReason) IsAbsent() bool {
	return o.value == nil
}

func (o *OptionalMutedInfoReason) Clear() {
	o.value = nil
}

func (o *OptionalMutedInfoReason) Get() chrome.MutedInfoReason {
	return *o.value
}

func (o *OptionalMutedInfoReason) Set(info chrome.MutedInfoReason) {
	o.value = &info
}

func (o *OptionalMutedInfoReason) OrElse(info chrome.MutedInfoReason) chrome.MutedInfoReason {
	if o.value == nil {
		return info
	}

	return *o.value
}

func (o *OptionalMutedInfoReason) With(f func(chrome.MutedInfoReason)) {
	if o.value != nil {
		f(*o.value)
	}
}
