// +build js,wasm
// +build chrome

package chrome

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/pkg/chrome"
)

const (
	jsKeyID                = "id"
	jsKeyLastError         = "lastError"
	jsKeyOnConnect         = "onConnect"
	jsKeyOnConnectExternal = "onConnectExternal"
	jsKeyOnConnectNative   = "onConnectExternal"

	jsFnConnect         = "connect"
	jsFnConnectNative   = "connectNative"
	jsFnGetManifest     = "getManifest"
	jsFnGetPlatformInfo = "getPlatformInfo"
	jsFnGetURL          = "getURL"
	jsFnOpenOptionsPage = "openOptionsPage"
	jsFnReload          = "reload"
	jsFnRestart         = "restart"
	jsFnSetUninstallURL = "setUninstallURL"
)

type RuntimeAPI struct {
	runtime js.Value

	conEvent    *ConnectEvent
	conExtEvent *ConnectEvent
	conNatEvent *ConnectEvent
}

func (r *RuntimeAPI) ID() string {
	return r.runtime.Get(jsKeyID).String()
}

func (r *RuntimeAPI) LastError() interface{} {
	return r.runtime.Get(jsKeyLastError)
}

func (r *RuntimeAPI) Connect(extID string, info chrome.ConnectInfo) chrome.Port {
	return &Port{port: r.runtime.Call(jsFnConnect, extID, info.(*ConnectInfo).JS())}
}

func (r *RuntimeAPI) ConnectNative(app string) chrome.Port {
	return &Port{port: r.runtime.Call(jsFnConnectNative, app)}
}

func (r *RuntimeAPI) GetBackgroundPage() chrome.Window {
}

func (r *RuntimeAPI) GetBackgroundPageAsync(cb chrome.WindowCallback) {
}

func (r *RuntimeAPI) GetManifest() interface{} {
	return r.runtime.Call(jsFnGetManifest)
}

func (r *RuntimeAPI) GetPackageDirectoryEntry() chrome.DirectoryEntry {
}

func (r *RuntimeAPI) GetPackageDirectoryEntryAsync(cb chrome.DirectoryEntryCallback) {
}

func (r *RuntimeAPI) GetPlatformInfo() (out chrome.PlatformInfo) {
	done := make(chan bool, 1)
	defer close(done)

	fn := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		out = NewPlatformInfo(args[0])
		return nil
	})
	defer fn.Release()

	r.runtime.Call(jsFnGetPlatformInfo)

	<-done

	return
}

func (r *RuntimeAPI) GetPlatformInfoAsync(cb chrome.PlatformInfoCallback) {
	fn := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		cb(NewPlatformInfo(args[0]))
		return nil
	})
	defer fn.Release()

	r.runtime.Call(jsFnGetPlatformInfo)
}

func (r *RuntimeAPI) GetURL(path string) string {
	return r.runtime.Call(jsFnGetURL, path).String()
}

func (r *RuntimeAPI) OpenOptionsPage() {
	fn, done := emptySyncCB()
	defer close(done)
	defer fn.Release()

	r.runtime.Call(jsFnOpenOptionsPage, fn)

	<-done
}

func (r *RuntimeAPI) OpenOptionsPageAsync(cb chrome.EmptyCallback) {
	fn := emptyAsyncCB(cb)
	defer fn.Release()
}

func (r *RuntimeAPI) Reload() {
	r.runtime.Call(jsFnReload)
}

func (r *RuntimeAPI) RequestUpdateCheck() (chrome.RequestUpdateCheckStatus, chrome.UpdateResponse) {
}

func (r *RuntimeAPI) RequestUpdateCheckAsync(cb chrome.RequestUpdateCheckCallback) {
}

func (r *RuntimeAPI) Restart() {
	r.runtime.Call(jsFnRestart)
}

func (r *RuntimeAPI) RestartAfterDelayAsync(seconds int, cb chrome.EmptyCallback) {
}

func (r *RuntimeAPI) SendMessageAsync(extID *string, msg interface{}, opts chrome.MessageOptions, cb chrome.AnyCallback) {
}

func (r *RuntimeAPI) SendNativeMessageAsync(app string, msg interface{}, cb chrome.AnyCallback) {
}

func (r *RuntimeAPI) SetUninstallURL(url string) {
	fn, done := emptySyncCB()
	defer close(done)
	defer fn.Release()

	r.runtime.Call(jsFnSetUninstallURL, url, fn)

	<-done
}

func (r *RuntimeAPI) SetUninstallURLAsync(url string, cb chrome.EmptyCallback) {
	fn := emptyAsyncCB(cb)
	defer fn.Release()

	r.runtime.Call(jsFnSetUninstallURL, url, fn)
}

func (r *RuntimeAPI) OnBrowserUpdateAvailable() chrome.NotificationEvent {
}

func (r *RuntimeAPI) OnConnect() chrome.ConnectEvent {
	if r.conEvent == nil {
		r.conEvent = NewConnectEvent(r.runtime.Get(jsKeyOnConnect))
	}

	return r.conEvent
}

func (r *RuntimeAPI) OnConnectExternal() chrome.ConnectEvent {
	if r.conExtEvent == nil {
		r.conExtEvent = NewConnectEvent(r.runtime.Get(jsKeyOnConnectExternal))
	}

	return r.conExtEvent
}

func (r *RuntimeAPI) OnConnectNative() chrome.ConnectEvent {
	if r.conNatEvent == nil {
		r.conNatEvent = NewConnectEvent(r.runtime.Get(jsKeyOnConnectNative))
	}

	return r.conNatEvent
}

func (r *RuntimeAPI) OnInstalled() chrome.InstallationEvent {
}

func (r *RuntimeAPI) OnMessage() chrome.RuntimeMessageEvent {
}

func (r *RuntimeAPI) OnMessageExternal() chrome.RuntimeMessageEvent {
}

func (r *RuntimeAPI) OnRestartRequired() chrome.RestartRequiredEvent {
}

func (r *RuntimeAPI) OnStartup() chrome.NotificationEvent {
}

func (r *RuntimeAPI) OnSuspend() chrome.NotificationEvent {
}

func (r *RuntimeAPI) OnSuspendCancelled() chrome.NotificationEvent {
}

func (r *RuntimeAPI) OnUpdateAvailable() chrome.UpdateAvailableEvent {
}

func (r *RuntimeAPI) Release() {
	if r.conEvent != nil {
		r.conEvent.Release()
	}
	if r.conExtEvent != nil {
		r.conExtEvent.Release()
	}
	if r.conNatEvent != nil {
		r.conNatEvent.Release()
	}
}
