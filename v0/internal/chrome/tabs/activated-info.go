package tabs

import (
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type ActivatedInfo struct {
	tabID    chrome.TabID
	windowID chrome.WindowID
}

func (a *ActivatedInfo) TabID() chrome.TabID {
	return a.tabID
}

func (a *ActivatedInfo) WindowID() chrome.WindowID {
	return a.windowID
}
