package tabs

import "github.com/foxcapades/groam/v0/pkg/chrome"

type OptionalZoomSettingsScope struct {
	value *chrome.ZoomSettingsScope
}

func (o *OptionalZoomSettingsScope) IsPresent() bool {
	return o.value != nil
}

func (o *OptionalZoomSettingsScope) IsAbsent() bool {
	return o.value == nil
}

func (o *OptionalZoomSettingsScope) Clear() {
	o.value = nil
}

func (o *OptionalZoomSettingsScope) Get() chrome.ZoomSettingsScope {
	return *o.value
}

func (o *OptionalZoomSettingsScope) Set(mode chrome.ZoomSettingsScope) {
	o.value = &mode
}

func (o *OptionalZoomSettingsScope) OrElse(mode chrome.ZoomSettingsScope) chrome.ZoomSettingsScope {
	if o.IsPresent() {
		return *o.value
	} else {
		return mode
	}
}

func (o *OptionalZoomSettingsScope) With(f func(chrome.ZoomSettingsScope)) {
	if o.IsPresent() {
		f(*o.value)
	}
}
