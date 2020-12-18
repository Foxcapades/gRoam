package tabs

import (
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type HighlightInfo struct {
	tabs     []chrome.TabID
	windowID chrome.WindowID
}

func (h *HighlightInfo) TabIDs() []chrome.TabID {
	return h.tabs
}

func (h *HighlightInfo) WindowID() chrome.WindowID {
	return h.windowID
}
