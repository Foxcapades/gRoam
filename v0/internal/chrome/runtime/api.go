// +build js,wasm
// +build chrome

package runtime

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type RuntimeAPI struct {
	runtime js.Value

	eConnect         chrome.ConnectEvent
	eConnectExternal chrome.ConnectEvent
	eConnectNative   chrome.ConnectEvent
	eInstall         chrome.InstallationEvent
	eMessage         chrome.RuntimeMessageEvent
	eMessageExternal chrome.RuntimeMessageEvent
	eRestartRequired chrome.RestartRequiredEvent
	eStartup         chrome.NotificationEvent
	eSuspend         chrome.NotificationEvent
	eSuspendCanceled chrome.NotificationEvent
	eUpdateAvail     chrome.UpdateAvailableEvent
}

func (r *RuntimeAPI) ID() string {
	return r.runtime.Get(x.JsKeyID).String()
}

func (r *RuntimeAPI) LastError() interface{} {
	return r.runtime.Get(x.JsKeyLastError)
}

func (r *RuntimeAPI) Connect(extID string, info chrome.ConnectInfo) chrome.Port {
	return &Port{port: r.runtime.Call(x.JsFnConnect, extID, info.(*ConnectInfo).JS())}
}

func (r *RuntimeAPI) ConnectNative(app string) chrome.Port {
	return &Port{port: r.runtime.Call(x.JsFnConnectNative, app)}
}

func (r *RuntimeAPI) GetBackgroundPage() chrome.Window {
}

func (r *RuntimeAPI) GetBackgroundPageAsync(cb chrome.WindowCallback) {
}

func (r *RuntimeAPI) GetManifest() interface{} {
	return r.runtime.Call(x.JsFnGetManifest)
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

	r.runtime.Call(x.JsFnGetPlatformInfo)

	<-done

	return
}

func (r *RuntimeAPI) GetPlatformInfoAsync(cb chrome.PlatformInfoCallback) {
	fn := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		cb(NewPlatformInfo(args[0]))
		return nil
	})
	defer fn.Release()

	r.runtime.Call(x.JsFnGetPlatformInfo)
}

func (r *RuntimeAPI) GetURL(path string) string {
	return r.runtime.Call(x.JsFnGetURL, path).String()
}

func (r *RuntimeAPI) OpenOptionsPage() {
	fn, done := x.EmptySyncCB()
	defer close(done)
	defer fn.Release()

	r.runtime.Call(x.JsFnOpenOptionsPage, fn)

	<-done
}

func (r *RuntimeAPI) OpenOptionsPageAsync(cb chrome.EmptyCallback) {
	fn := x.EmptyAsyncCB(cb)
	defer fn.Release()
}

func (r *RuntimeAPI) Reload() {
	r.runtime.Call(x.JsFnReload)
}

func (r *RuntimeAPI) RequestUpdateCheck() (chrome.RequestUpdateCheckStatus, chrome.UpdateResponse) {
}

func (r *RuntimeAPI) RequestUpdateCheckAsync(cb chrome.RequestUpdateCheckCallback) {
}

func (r *RuntimeAPI) Restart() {
	r.runtime.Call(x.JsFnRestart)
}

func (r *RuntimeAPI) RestartAfterDelayAsync(seconds int, cb chrome.EmptyCallback) {
}

func (r *RuntimeAPI) SendMessageAsync(extID *string, msg interface{}, opts chrome.MessageOptions, cb chrome.AnyCallback) {
}

func (r *RuntimeAPI) SendNativeMessageAsync(app string, msg interface{}, cb chrome.AnyCallback) {
}

func (r *RuntimeAPI) SetUninstallURL(url string) {
	fn, done := x.EmptySyncCB()
	defer close(done)
	defer fn.Release()

	r.runtime.Call(x.JsFnSetUninstallURL, url, fn)

	<-done
}

func (r *RuntimeAPI) SetUninstallURLAsync(url string, cb chrome.EmptyCallback) {
	fn := x.EmptyAsyncCB(cb)
	defer fn.Release()

	r.runtime.Call(x.JsFnSetUninstallURL, url, fn)
}

func (r *RuntimeAPI) OnBrowserUpdateAvailable() chrome.NotificationEvent {
}

func (r *RuntimeAPI) OnConnect() chrome.ConnectEvent {
	if r.eConnect == nil {
		r.eConnect = NewConnectEvent(r.runtime.Get(x.JsKeyOnConnect))
	}

	return r.eConnect
}

func (r *RuntimeAPI) OnConnectExternal() chrome.ConnectEvent {
	if r.eConnectExternal == nil {
		r.eConnectExternal = NewConnectEvent(r.runtime.Get(x.JsKeyOnConnectExternal))
	}

	return r.eConnectExternal
}

func (r *RuntimeAPI) OnConnectNative() chrome.ConnectEvent {
	if r.eConnectNative == nil {
		r.eConnectNative = NewConnectEvent(r.runtime.Get(x.JsKeyOnConnectNative))
	}

	return r.eConnectNative
}

func (r *RuntimeAPI) OnInstalled() chrome.InstallationEvent {
	if r.eInstall == nil {
		r.eInstall = NewInstallationEvent(r.runtime.Get(x.JsKeyOnInstalled))
	}

	return r.eInstall
}

func (r *RuntimeAPI) OnMessage() chrome.RuntimeMessageEvent {
	if r.eMessage == nil {
		r.eMessage = NewMessageEvent(r.runtime.Get(x.JsKeyOnMessage))
	}

	return r.eMessage
}

func (r *RuntimeAPI) OnMessageExternal() chrome.RuntimeMessageEvent {
	if r.eMessageExternal == nil {
		r.eMessageExternal = NewMessageEvent(r.runtime.Get(x.JsKeyOnMessageExternal))
	}

	return r.eMessageExternal
}

func (r *RuntimeAPI) OnRestartRequired() chrome.RestartRequiredEvent {
	if r.eRestartRequired == nil {
		r.eRestartRequired = NewRestartRequiredEvent(r.runtime.Get(x.JsKeyOnRestartRequired))
	}

	return r.eRestartRequired
}

func (r *RuntimeAPI) OnStartup() chrome.NotificationEvent {
	if r.eStartup == nil {
		r.eStartup = NewNotificationEvent(r.runtime.Get(x.JsKeyOnStartup))
	}

	return r.eStartup
}

func (r *RuntimeAPI) OnSuspend() chrome.NotificationEvent {
	if r.eSuspend == nil {
		r.eSuspend = NewNotificationEvent(r.runtime.Get(x.JsKeyOnSuspend))
	}

	return r.eSuspend
}

func (r *RuntimeAPI) OnSuspendCancelled() chrome.NotificationEvent {
	if r.eSuspendCanceled == nil {
		r.eSuspendCanceled = NewNotificationEvent(r.runtime.Get(x.JsKeyOnSuspendCanceled))
	}

	return r.eSuspendCanceled
}

func (r *RuntimeAPI) OnUpdateAvailable() chrome.UpdateAvailableEvent {
	if r.eUpdateAvail == nil {
		r.eUpdateAvail = NewUpdateAvailableEvent(r.runtime.Get(x.JsKeyOnUpdateAvailable))
	}

	return r.eUpdateAvail
}

func (r *RuntimeAPI) Release() {
	if r.eConnect != nil {
		r.eConnect.Release()
	}

	if r.eConnectExternal != nil {
		r.eConnectExternal.Release()
	}

	if r.eConnectNative != nil {
		r.eConnectNative.Release()
	}

	if r.eInstall != nil {
		r.eInstall.Release()
	}

	if r.eUpdateAvail != nil {
		r.eUpdateAvail.Release()
	}
}
