package chrome

import "github.com/foxcapades/groam/v0/pkg/chrome"

type OptionalImageFormat struct {
	v **chrome.ImageFormat
}

func (o *OptionalImageFormat) IsPresent() bool {
	return *o.v != nil
}

func (o *OptionalImageFormat) IsAbsent() bool {
	return *o.v == nil
}

func (o *OptionalImageFormat) Clear() {
	*o.v = nil
}

func (o *OptionalImageFormat) Get() chrome.ImageFormat {
	return **o.v
}

func (o *OptionalImageFormat) Set(format chrome.ImageFormat) {
	*o.v = &format
}

func (o *OptionalImageFormat) OrElse(format chrome.ImageFormat) chrome.ImageFormat {
	if *o.v == nil {
		return format
	}

	return **o.v
}

func (o *OptionalImageFormat) With(f func(chrome.ImageFormat)) {
	if *o.v != nil {
		f(**o.v)
	}
}

