// +build js,wasm
// +build chrome

package x

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

func NewDirectoryEntry(v js.Value) chrome.DirectoryEntry {
	return &DirectoryEntry{
		isFile:     v.Get(JsKeyIsFile).Bool(),
		name:       v.Get(JsKeyName).String(),
		fullPath:   v.Get(JsKeyFullPath).String(),
		fileSystem: v.Get(JsKeyFilesystem),
	}
}
