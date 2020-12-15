package chrome

import "github.com/foxcapades/groam/v0/pkg/chrome"

type InjectDetails struct {
	allFrames       *bool
	code            *string
	cssOrigin       *chrome.CSSOrigin
	file            *string
	frameId         *chrome.FrameID
	matchAboutBlank *bool
	runAt           *chrome.RunAt
}

func (i *InjectDetails) AllFrames() chrome.OptionalBool {
	return &OptionalBool{&i.allFrames}
}

func (i *InjectDetails) SetAllFrames(b bool) chrome.InjectDetails {
	i.allFrames = &b
	return i
}

func (i *InjectDetails) Code() chrome.OptionalString {
	return &OptionalString{&i.code}
}

func (i *InjectDetails) SetCode(s string) chrome.InjectDetails {
	i.code = &s
	return i
}

func (i *InjectDetails) CSSOrigin() chrome.OptionalCSSOrigin {
	return &OptionalCSSOrigin{&i.cssOrigin}
}

func (i *InjectDetails) SetCSSOrigin(origin chrome.CSSOrigin) chrome.InjectDetails {
	i.cssOrigin = &origin
	return i
}

func (i *InjectDetails) File() chrome.OptionalString {
	return &OptionalString{&i.file}
}

func (i *InjectDetails) SetFile(s string) chrome.InjectDetails {
	i.file = &s
	return i
}

func (i *InjectDetails) FrameID() chrome.OptionalFrameID {
	return &OptionalFrameID{&i.frameId}
}

func (i *InjectDetails) SetFrameID(id chrome.FrameID) chrome.InjectDetails {
	i.frameId = &id
	return i
}

func (i *InjectDetails) MatchAboutBlank() chrome.OptionalBool {
	return &OptionalBool{&i.matchAboutBlank}
}

func (i *InjectDetails) SetMatchAboutBlank(b bool) chrome.InjectDetails {
	i.matchAboutBlank = &b
	return i
}

func (i *InjectDetails) RunAt() chrome.OptionalRunAt {
	return &OptionalRunAt{&i.runAt}
}

func (i *InjectDetails) SetRunAt(at chrome.RunAt) chrome.InjectDetails {
	i.runAt = &at
	return i
}
