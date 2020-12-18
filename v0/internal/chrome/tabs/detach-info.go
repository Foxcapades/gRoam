package tabs

import (
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type DetachInfo struct {
	position uint16
	windowID chrome.WindowID
}

func (d *DetachInfo) NewPosition() uint16 {
	return d.position
}

func (d *DetachInfo) NewWindowID() chrome.WindowID {
	return d.windowID
}

func (d *DetachInfo) OldPosition() uint16 {
	return d.position
}

func (d *DetachInfo) OldWindowID() chrome.WindowID {
	return d.windowID
}
