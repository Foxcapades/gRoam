// +build js,wasm
// +build chrome

package x

import (
	"sync"
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

func EmptySyncCB() (out js.Func, wait sync.WaitGroup) {
	wait.Add(1)
	out = js.FuncOf(func(js.Value, []js.Value) interface{} {
		wait.Done()
		return nil
	})

	return
}
