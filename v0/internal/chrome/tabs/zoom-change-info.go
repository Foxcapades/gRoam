package tabs

import (
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type ZoomChangeInfo struct {
	newFactor float32
	oldFactor float32
	tabID     chrome.TabID
	settings  chrome.ZoomSettings
}

func (z *ZoomChangeInfo) NewZoomFactor() float32 {
	return z.newFactor
}

func (z *ZoomChangeInfo) OldZoomFactor() float32 {
	return z.oldFactor
}

func (z *ZoomChangeInfo) TabID() chrome.TabID {
	panic("implement me")
}

func (z *ZoomChangeInfo) ZoomSettings() chrome.ZoomSettings {
	panic("implement me")
}
