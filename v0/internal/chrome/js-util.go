// +build js,wasm
// +build chrome

package chrome

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

var jsObject = js.Global().Get("Object")
var jsArray = js.Global().Get("Array")


func emptyAsyncCB(cb chrome.EmptyCallback) js.Func {
	return js.FuncOf(func(js.Value, []js.Value) interface{} {
		cb()
		return nil
	})
}

func emptySyncCB() (js.Func, chan bool) {
	done := make(chan bool, 1)

	return js.FuncOf(func(js.Value, []js.Value) interface{} {
		done <- true
		return nil
	}), done
}
