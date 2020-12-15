package chrome

import (
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type DeleteInjectionDetails struct {
	allFrames       *bool
	code            *string
	cssOrigin       *chrome.CSSOrigin
	file            *string
	frameId         *chrome.FrameID
	matchAboutBlank *bool
}

func (d *DeleteInjectionDetails) AllFrames() chrome.OptionalBool {
	return &OptionalBool{&d.allFrames}
}

func (d *DeleteInjectionDetails) Code() chrome.OptionalString {
	return &OptionalString{&d.code}
}

func (d *DeleteInjectionDetails) CSSOrigin() chrome.OptionalCSSOrigin {
	return &OptionalCSSOrigin{&d.cssOrigin}
}

func (d *DeleteInjectionDetails) File() chrome.OptionalString {
	return &OptionalString{&d.file}
}

func (d *DeleteInjectionDetails) FrameID() chrome.OptionalFrameID {
	return &OptionalFrameID{&d.frameId}
}

func (d *DeleteInjectionDetails) MatchAboutBlank() chrome.OptionalBool {
	return &OptionalBool{&d.matchAboutBlank}
}
