package tabs

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewOptionalGroupID(v js.Value) chrome.OptionalGroupID {
	out := new(OptionalGroupID)

	if !v.IsUndefined() && !v.IsNull() {
		tmp := chrome.GroupID(v.Int())
		out.value = &tmp
	}

	return out
}
