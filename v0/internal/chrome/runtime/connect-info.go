package runtime

import (
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type ConnectInfo struct {
	includeTLS chrome.OptionalBool
	name       chrome.OptionalString
}

func (c *ConnectInfo) IncludeTLSChannelID() chrome.OptionalBool {
	return c.includeTLS
}

func (c *ConnectInfo) Name() chrome.OptionalString {
	return c.name
}
