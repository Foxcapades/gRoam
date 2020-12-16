package runtime

import (
	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type MessageOptions struct {
	includeTLS *bool
}

func (m *MessageOptions) IncludeTLSChannelID() chrome.OptionalBool {
	return &x.OptionalBool{Value: &m.includeTLS}
}
