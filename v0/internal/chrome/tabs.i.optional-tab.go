package chrome

import "github.com/foxcapades/groam/v0/pkg/chrome"

type OptionalTab struct {
	value **chrome.Tab
}

func (o *OptionalTab) IsPresent() bool {
	return *o.value != nil
}

func (o *OptionalTab) IsAbsent() bool {
	return *o.value == nil
}

func (o *OptionalTab) Clear() {
	*o.value = nil
}

func (o *OptionalTab) Get() chrome.Tab {
	return **o.value
}

func (o *OptionalTab) Set(b chrome.Tab) {
	*o.value = &b
}

func (o *OptionalTab) OrElse(b chrome.Tab) chrome.Tab {
	if *o.value == nil {
		return b
	}

	return **o.value
}

func (o *OptionalTab) With(f func(chrome.Tab)) {
	if *o.value != nil {
		f(**o.value)
	}
}
