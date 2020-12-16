package runtime

import (
	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type ConnectInfo struct {
	includeTLS *bool
	name       *string
}

func (c *ConnectInfo) IncludeTLSChannelID() chrome.OptionalBool {
	return &x.OptionalBool{V: &c.includeTLS}
}

func (c *ConnectInfo) Name() chrome.OptionalString {
	return &x.OptionalString{V: &c.name}
}
