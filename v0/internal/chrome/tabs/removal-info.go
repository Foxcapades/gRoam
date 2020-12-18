package tabs

import (
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type RemovalInfo struct {
	closing  bool
	windowID chrome.WindowID
}

func (r *RemovalInfo) IsWindowClosing() bool {
	return r.closing
}

func (r *RemovalInfo) WindowID() chrome.WindowID {
	return r.windowID
}
