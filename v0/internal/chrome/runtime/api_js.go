// +build js,wasm
// +build chrome

package runtime

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/windows"
	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type API struct {
	runtime js.Value

	eBrowserUpdate   chrome.NotificationEvent
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

func (r *API) ID() string {
	return r.runtime.Get(x.JsKeyID).String()
}

func (r *API) LastError() interface{} {
	return r.runtime.Get(x.JsKeyLastError)
}

func (r *API) Connect(extID string, info chrome.ConnectInfo) chrome.Port {
	return &Port{port: r.runtime.Call(x.JsFnConnect, extID, SerializeConnectInfo(info))}
}

func (r *API) ConnectNative(app string) chrome.Port {
	return &Port{port: r.runtime.Call(x.JsFnConnectNative, app)}
}

func (r *API) GetBackgroundPage() (out chrome.Window) {
	wg := x.NewWaitSingle()
	fn := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		out = windows.NewWindow(args[0])
		wg.Done()
		return nil
	})
	defer fn.Release()

	r.runtime.Call(x.JsFnGetBackgroundPage, fn)
	wg.Wait()

	return
}

func (r *API) GetBackgroundPageAsync(cb chrome.WindowCallback) {
	fn := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		cb(windows.NewWindow(args[0]))
		return nil
	})
	defer fn.Release()

	r.runtime.Call(x.JsFnGetBackgroundPage, fn)
}

func (r *API) GetManifest() interface{} {
	return r.runtime.Call(x.JsFnGetManifest)
}

func (r *API) GetPackageDirectoryEntry() (out chrome.DirectoryEntry) {
	wait := x.NewWaitSingle()

	fn := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		out = x.NewDirectoryEntry(args[0])
		wait.Done()
		return nil
	})
	defer fn.Release()

	r.runtime.Call(x.JsFnGetPackageDirectoryEntry, fn)
	wait.Wait()

	return
}

func (r *API) GetPackageDirectoryEntryAsync(cb chrome.DirectoryEntryCallback) {
	fn := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		cb(x.NewDirectoryEntry(args[0]))
		return nil
	})
	defer fn.Release()

	r.runtime.Call(x.JsFnGetPackageDirectoryEntry, fn)
}

func (r *API) GetPlatformInfo() (out chrome.PlatformInfo) {
	done := x.NewWaitSingle()

	fn := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		out = NewPlatformInfo(args[0])
		done.Done()

		return nil
	})
	defer fn.Release()

	r.runtime.Call(x.JsFnGetPlatformInfo)
	done.Wait()

	return
}

func (r *API) GetPlatformInfoAsync(cb chrome.PlatformInfoCallback) {
	fn := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		cb(NewPlatformInfo(args[0]))
		return nil
	})
	defer fn.Release()

	r.runtime.Call(x.JsFnGetPlatformInfo)
}

func (r *API) GetURL(path string) string {
	return r.runtime.Call(x.JsFnGetURL, path).String()
}

func (r *API) OpenOptionsPage() {
	fn, done := x.EmptySyncCB()
	defer fn.Release()

	r.runtime.Call(x.JsFnOpenOptionsPage, fn)
	done.Wait()
}

func (r *API) OpenOptionsPageAsync(cb chrome.EmptyCallback) {
	fn := x.EmptyAsyncCB(cb)
	defer fn.Release()
}

func (r *API) Reload() {
	r.runtime.Call(x.JsFnReload)
}

func (r *API) RequestUpdateCheck() (
	status chrome.RequestUpdateCheckStatus,
	info chrome.UpdateAvailableDetails,
) {
	done := x.NewWaitSingle()

	fn := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		status = chrome.RequestUpdateCheckStatus(args[0].String())
		info = NewUpdateAvailableDetails(args[1])

		done.Done()
		return nil
	})

	r.runtime.Call(x.JsFnRequestUpdateCheck, fn)
	done.Wait()

	return
}

func (r *API) RequestUpdateCheckAsync(cb chrome.RequestUpdateCheckCallback) {
	go func() {
		cb(r.RequestUpdateCheck())
	}()
}

func (r *API) Restart() {
	r.runtime.Call(x.JsFnRestart)
}

func (r *API) RestartAfterDelay(seconds int) {
	done := x.NewWaitSingle()

	fn := js.FuncOf(func(js.Value, []js.Value) interface{} {
		done.Done()
		return nil
	})
	defer fn.Release()

	r.runtime.Call(x.JsFnRestartAfterDelay, seconds, fn)
	done.Wait()

	return
}

func (r *API) RestartAfterDelayAsync(seconds int, cb chrome.EmptyCallback) {
	go func() {
		r.RestartAfterDelay(seconds)
		cb()
	}()
}

func (r *API) SendMessage(
	extID *string,
	msg interface{},
	opts chrome.MessageOptions,
) (response interface{}) {
	done := x.NewWaitSingle()

	fn := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		response = args[0]
		done.Done()

		return nil
	})
	defer fn.Release()

	r.runtime.Call(x.JsFnSendMessage, extID, msg, MessageOption2JS(opts), fn)
	done.Wait()

	return
}

func (r *API) SendMessageAsync(
	extID *string,
	msg interface{},
	opts chrome.MessageOptions,
	cb chrome.AnyCallback,
) {
	go func() {
		cb(r.SendMessage(extID, msg, opts))
	}()
}

func (r *API) SendNativeMessage(app string, msg interface{}) (response interface{}) {
	done := x.NewWaitSingle()

	fn := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		response = args[0]
		done.Done()

		return nil
	})
	defer fn.Release()

	r.runtime.Call(x.JsFnSendNativeMsg, app, msg, fn)
	done.Wait()

	return
}

func (r *API) SendNativeMessageAsync(app string, msg interface{}, cb chrome.AnyCallback) {
	go func() {
		cb(r.SendNativeMessage(app, msg))
	}()
}

func (r *API) SetUninstallURL(url string) {
	fn, done := x.EmptySyncCB()
	defer fn.Release()

	r.runtime.Call(x.JsFnSetUninstallURL, url, fn)
	done.Wait()
}

func (r *API) SetUninstallURLAsync(url string, cb chrome.EmptyCallback) {
	fn := x.EmptyAsyncCB(cb)
	defer fn.Release()

	r.runtime.Call(x.JsFnSetUninstallURL, url, fn)
}

func (r *API) OnBrowserUpdateAvailable() chrome.NotificationEvent {
	if r.eBrowserUpdate == nil {
		r.eBrowserUpdate = NewNotificationEvent(r.runtime.Get(x.JsKeyOnBrowserUpdateAvail))
	}

	return r.eBrowserUpdate
}

func (r *API) OnConnect() chrome.ConnectEvent {
	if r.eConnect == nil {
		r.eConnect = NewConnectEvent(r.runtime.Get(x.JsKeyOnConnect))
	}

	return r.eConnect
}

func (r *API) OnConnectExternal() chrome.ConnectEvent {
	if r.eConnectExternal == nil {
		r.eConnectExternal = NewConnectEvent(r.runtime.Get(x.JsKeyOnConnectExternal))
	}

	return r.eConnectExternal
}

func (r *API) OnConnectNative() chrome.ConnectEvent {
	if r.eConnectNative == nil {
		r.eConnectNative = NewConnectEvent(r.runtime.Get(x.JsKeyOnConnectNative))
	}

	return r.eConnectNative
}

func (r *API) OnInstalled() chrome.InstallationEvent {
	if r.eInstall == nil {
		r.eInstall = NewInstallationEvent(r.runtime.Get(x.JsKeyOnInstalled))
	}

	return r.eInstall
}

func (r *API) OnMessage() chrome.RuntimeMessageEvent {
	if r.eMessage == nil {
		r.eMessage = NewMessageEvent(r.runtime.Get(x.JsKeyOnMessage))
	}

	return r.eMessage
}

func (r *API) OnMessageExternal() chrome.RuntimeMessageEvent {
	if r.eMessageExternal == nil {
		r.eMessageExternal = NewMessageEvent(r.runtime.Get(x.JsKeyOnMessageExternal))
	}

	return r.eMessageExternal
}

func (r *API) OnRestartRequired() chrome.RestartRequiredEvent {
	if r.eRestartRequired == nil {
		r.eRestartRequired = NewRestartRequiredEvent(r.runtime.Get(x.JsKeyOnRestartRequired))
	}

	return r.eRestartRequired
}

func (r *API) OnStartup() chrome.NotificationEvent {
	if r.eStartup == nil {
		r.eStartup = NewNotificationEvent(r.runtime.Get(x.JsKeyOnStartup))
	}

	return r.eStartup
}

func (r *API) OnSuspend() chrome.NotificationEvent {
	if r.eSuspend == nil {
		r.eSuspend = NewNotificationEvent(r.runtime.Get(x.JsKeyOnSuspend))
	}

	return r.eSuspend
}

func (r *API) OnSuspendCancelled() chrome.NotificationEvent {
	if r.eSuspendCanceled == nil {
		r.eSuspendCanceled = NewNotificationEvent(r.runtime.Get(x.JsKeyOnSuspendCanceled))
	}

	return r.eSuspendCanceled
}

func (r *API) OnUpdateAvailable() chrome.UpdateAvailableEvent {
	if r.eUpdateAvail == nil {
		r.eUpdateAvail = NewUpdateAvailableEvent(r.runtime.Get(x.JsKeyOnUpdateAvailable))
	}

	return r.eUpdateAvail
}

func (r *API) Release() {
	if r.eBrowserUpdate != nil {
		r.eBrowserUpdate.Release()
	}

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

	if r.eMessage != nil {
		r.eMessage.Release()
	}

	if r.eMessageExternal != nil {
		r.eMessageExternal.Release()
	}

	if r.eRestartRequired != nil {
		r.eRestartRequired.Release()
	}

	if r.eStartup != nil {
		r.eStartup.Release()
	}

	if r.eSuspend != nil {
		r.eSuspend.Release()
	}

	if r.eSuspendCanceled != nil {
		r.eSuspendCanceled.Release()
	}

	if r.eUpdateAvail != nil {
		r.eUpdateAvail.Release()
	}
}
