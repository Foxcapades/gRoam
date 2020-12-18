package tabs

import (
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type MoveInfo struct {
	from, to uint16
	windowID chrome.WindowID
}

func (m *MoveInfo) FromIndex() uint16 {
	return m.from
}

func (m *MoveInfo) ToIndex() uint16 {
	return m.to
}

func (m *MoveInfo) WindowID() chrome.WindowID {
	return m.windowID
}
