package tabs

import "github.com/foxcapades/groam/v0/pkg/chrome"

type OptionalZoomSettingsMode struct {
	value *chrome.ZoomSettingsMode
}

func (o *OptionalZoomSettingsMode) IsPresent() bool {
	return o.value != nil
}

func (o *OptionalZoomSettingsMode) IsAbsent() bool {
	return o.value == nil
}

func (o *OptionalZoomSettingsMode) Clear() {
	o.value = nil
}

func (o *OptionalZoomSettingsMode) Get() chrome.ZoomSettingsMode {
	return *o.value
}

func (o *OptionalZoomSettingsMode) Set(mode chrome.ZoomSettingsMode) {
	o.value = &mode
}

func (o *OptionalZoomSettingsMode) OrElse(mode chrome.ZoomSettingsMode) chrome.ZoomSettingsMode {
	if o.IsPresent() {
		return *o.value
	} else {
		return mode
	}
}

func (o *OptionalZoomSettingsMode) With(f func(chrome.ZoomSettingsMode)) {
	if o.IsPresent() {
		f(*o.value)
	}
}
