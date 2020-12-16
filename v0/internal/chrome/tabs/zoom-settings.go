package tabs

import (
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type ZoomSettings struct {
	defZoomFactor chrome.OptionalF32
	mode          chrome.OptionalZoomSettingsMode
	scope         chrome.OptionalZoomSettingsScope
}

func (z *ZoomSettings) DefaultZoomFactor() chrome.OptionalF32 {
	panic("implement me")
}

func (z *ZoomSettings) Mode() chrome.OptionalZoomSettingsMode {
	panic("implement me")
}

func (z *ZoomSettings) Scope() chrome.OptionalZoomSettingsScope {
	panic("implement me")
}
