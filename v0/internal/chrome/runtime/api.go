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

	conEvent     *ConnectEvent
	conExtEvent  *ConnectEvent
	conNatEvent  *ConnectEvent
	installEvent *InstallationEvent
	eUpdateEvent *UpdateAvailableEvent

	eStartup         chrome.NotificationEvent
	eSuspend         chrome.NotificationEvent
	eSuspendCanceled chrome.NotificationEvent
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
	if r.conEvent == nil {
		r.conEvent = NewConnectEvent(r.runtime.Get(x.JsKeyOnConnect))
	}

	return r.conEvent
}

func (r *RuntimeAPI) OnConnectExternal() chrome.ConnectEvent {
	if r.conExtEvent == nil {
		r.conExtEvent = NewConnectEvent(r.runtime.Get(x.JsKeyOnConnectExternal))
	}

	return r.conExtEvent
}

func (r *RuntimeAPI) OnConnectNative() chrome.ConnectEvent {
	if r.conNatEvent == nil {
		r.conNatEvent = NewConnectEvent(r.runtime.Get(x.JsKeyOnConnectNative))
	}

	return r.conNatEvent
}

func (r *RuntimeAPI) OnInstalled() chrome.InstallationEvent {
	if r.installEvent == nil {
		r.installEvent = NewInstallationEvent(r.runtime.Get(x.JsKeyOnInstalled))
	}

	return r.installEvent
}

func (r *RuntimeAPI) OnMessage() chrome.RuntimeMessageEvent {
}

func (r *RuntimeAPI) OnMessageExternal() chrome.RuntimeMessageEvent {
}

func (r *RuntimeAPI) OnRestartRequired() chrome.RestartRequiredEvent {
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
	if r.eUpdateEvent == nil {
		r.eUpdateEvent = NewUpdateAvailableEvent(r.runtime.Get(x.JsKeyOnUpdateAvailable))
	}

	return r.eUpdateEvent
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

	if r.installEvent != nil {
		r.installEvent.Release()
	}

	if r.eUpdateEvent != nil {
		r.eUpdateEvent.Release()
	}
}
