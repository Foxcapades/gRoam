package chrome

import (
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type ImageDetails struct {
	format  *chrome.ImageFormat
	quality *uint8
}

func (i *ImageDetails) Format() chrome.OptionalImageFormat {
	return &OptionalImageFormat{&i.format}
}

func (i *ImageDetails) Quality() chrome.OptionalU8 {
	return &OptionalU8{&i.quality}
}
