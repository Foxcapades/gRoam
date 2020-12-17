package tabs

import (
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type SelectInfo struct {
	windowID chrome.WindowID
}

func (s *SelectInfo) WindowID() chrome.WindowID {
	return s.windowID
}
