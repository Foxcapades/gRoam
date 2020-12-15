// +build js,wasm
// +build chrome

package chrome

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

const (
	jsKeyArch = "arch"
	jsKeyNACL = "nacl_arch"
	jsKeyOS   = "os"
)

func NewPlatformInfo(raw js.Value) (out *PlatformInfo) {
	out = new(PlatformInfo)

	out.arch = chrome.PlatformArch(raw.Get(jsKeyArch).String())
	out.nacl = chrome.PlatformNACLArch(raw.Get(jsKeyNACL).String())
	out.os   = chrome.PlatformOS(raw.Get(jsKeyOS).String())

	return
}

type PlatformInfo struct {
	arch chrome.PlatformArch
	nacl chrome.PlatformNACLArch
	os   chrome.PlatformOS
}

func (p *PlatformInfo) Arch() chrome.PlatformArch {
	panic("implement me")
}

func (p *PlatformInfo) NaclArch() chrome.PlatformNACLArch {
	panic("implement me")
}

func (p *PlatformInfo) OS() chrome.PlatformOS {
	panic("implement me")
}
