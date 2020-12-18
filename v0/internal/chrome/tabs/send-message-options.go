package tabs

import (
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type SendMessageOptions struct {
	frameID chrome.OptionalFrameID
}

func (s *SendMessageOptions) FrameID() chrome.OptionalFrameID {
	return s.frameID
}

func (s *SendMessageOptions) SetFrameID(id chrome.FrameID) chrome.SendMessageOptions {
	s.frameID.Set(id)
	return s
}
