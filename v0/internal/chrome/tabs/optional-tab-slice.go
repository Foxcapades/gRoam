package tabs

import "github.com/foxcapades/groam/v0/pkg/chrome"

type OptionalTabSlice struct {
	Value []chrome.Tab
}

func (o *OptionalTabSlice) IsPresent() bool {
	return o.Value != nil
}

func (o *OptionalTabSlice) IsAbsent() bool {
	return o.Value == nil
}

func (o *OptionalTabSlice) Clear() {
	o.Value = nil
}

func (o *OptionalTabSlice) Get() []chrome.Tab {
	return o.Value
}

func (o *OptionalTabSlice) Set(tabs []chrome.Tab) {
	o.Value = tabs
}

func (o *OptionalTabSlice) Add(tab chrome.Tab) {
	o.Value = append(o.Value, tab)
}

func (o *OptionalTabSlice) OrElse(tabs []chrome.Tab) []chrome.Tab {
	if o.IsAbsent() {
		return tabs
	}

	return o.Value
}

func (o *OptionalTabSlice) With(f func([]chrome.Tab)) {
	if o.IsPresent() {
		f(o.Value)
	}
}
