package chrome

import (
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type OptionalMutedInfoReason struct {
	v **chrome.MutedInfoReason
}

func (o *OptionalMutedInfoReason) IsPresent() bool {
	return *o.v != nil
}

func (o *OptionalMutedInfoReason) IsAbsent() bool {
	return *o.v == nil
}

func (o *OptionalMutedInfoReason) Clear() {
	*o.v = nil
}

func (o *OptionalMutedInfoReason) Get() chrome.MutedInfoReason {
	return **o.v
}

func (o *OptionalMutedInfoReason) Set(info chrome.MutedInfoReason) {
	*o.v = &info
}

func (o *OptionalMutedInfoReason) OrElse(info chrome.MutedInfoReason) chrome.MutedInfoReason {
	if *o.v == nil {
		return info
	}

	return **o.v
}

func (o *OptionalMutedInfoReason) With(f func(chrome.MutedInfoReason)) {
	if *o.v != nil {
		f(**o.v)
	}
}
