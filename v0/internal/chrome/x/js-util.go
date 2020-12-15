// +build js,wasm
// +build chrome

package x

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

var JsObject = js.Global().Get("Object")
var JsArray = js.Global().Get("Array")

func EmptyAsyncCB(cb chrome.EmptyCallback) js.Func {
	return js.FuncOf(func(js.Value, []js.Value) interface{} {
		cb()
		return nil
	})
}

func EmptySyncCB() (js.Func, chan bool) {
	done := make(chan bool, 1)

	return js.FuncOf(func(js.Value, []js.Value) interface{} {
		done <- true
		return nil
	}), done
}
