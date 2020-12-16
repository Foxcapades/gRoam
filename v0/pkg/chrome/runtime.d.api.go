package chrome

// chrome.runtime
//
// Use the chrome.runtime API to retrieve the background page, return details
// about the manifest, and listen for and respond to events in the app or
// extension lifecycle. You can also use this API to convert the relative path
// of URLs to fully-qualified URLs.
type RuntimeAPI interface {
	// The ID of the extension/app.
	ID() string
	LastError() interface{}

	// Attempts to connect to connect listeners within an extension/app (such as
	// the background page), or other extensions/apps.
	//
	// This is useful for content scripts connecting to their extension processes,
	// inter-app/extension communication, and web messaging. Note that this does
	// not connect to any listeners in a content script. Extensions may connect to
	// content scripts embedded in tabs via tabs.connect.
	//
	// Arguments:
	//   extID = The ID of the extension or app to connect to. If omitted, a
	//           connection will be attempted with your own extension. Required if
	//           sending messages from a web page for web messaging.
	//   info  =
	//
	// Returns:
	//   [0] = Port through which messages can be sent and received. The port's
	//         onDisconnect event is fired if the extension/app does not exist.
	//   [1] = Any error that occurred while attempting to make this Chrome API
	//         call.
	Connect(extID string, info ConnectInfo) Port

	// Connects to a native application in the host machine.
	//
	// See Native Messaging for more information
	//
	// Arguments:
	//   application = The name of the registered application to connect to.
	//
	// Returns:
	//   [0] = Port through which messages can be sent and received with the
	//         application
	//   [1] = Any error that occurred while attempting to make this Chrome API
	//         call.
	ConnectNative(application string) Port

	// Retrieves the JavaScript 'window' object for the background page running
	// inside the current extension/app.
	//
	// If the background page is an event page, the system will ensure it is
	// loaded before calling the callback.
	//
	// If there is no background page, an error is set.
	//
	// Returns:
	//   [0] = The JavaScript 'window' object for the background page.
	//   [1] = Any error that occurred while attempting to make this Chrome API
	//         call.
	GetBackgroundPage() Window

	// Retrieves the JavaScript 'window' object for the background page running
	// inside the current extension/app.
	//
	// If the background page is an event page, the system will ensure it is
	// loaded before calling the callback.
	//
	// If there is no background page, an error is set.
	//
	// Arguments:
	//   cb = Callback function that accepts a Window.
	GetBackgroundPageAsync(cb WindowCallback)

	// Returns details about the app or extension from the manifest.
	//
	// The object returned is a serialization of the full manifest file.
	//
	// Returns:
	//   [0] = The parsed JSON manifest for this extension.
	//   [1] = Any error that occurred while attempting to make this Chrome API
	//         call.
	GetManifest() interface{}

	// Returns a DirectoryEntry for the package directory.
	//
	// Returns:
	//   [0] =
	//   [1] = Any error that occurred while attempting to make this Chrome API
	//         call.
	GetPackageDirectoryEntry() DirectoryEntry

	// Returns a DirectoryEntry for the package directory.
	//
	// Arguments:
	//   cb = Callback function that accepts a DirectoryEntry instance.
	GetPackageDirectoryEntryAsync(cb DirectoryEntryCallback)

	// Returns information about the current platform.
	//
	// Returns:
	//   [0] = Information about the current platform.
	//   [1] = Any error that occurred while attempting to make this Chrome API
	//         call.
	GetPlatformInfo() PlatformInfo

	// Returns information about the current platform.
	//
	// Arguments:
	//   cb = Callback function that accepts a PlatformInfo instance and an error.
	GetPlatformInfoAsync(cb PlatformInfoCallback)

	// Converts a relative path within an app/extension install directory to a
	// fully-qualified URL.
	//
	// Arguments:
	//   path = A path to a resource within an app/extension expressed relative to
	//          its install directory.
	//
	// Returns:
	//   [0] = The fully-qualified URL to the resource.
	GetURL(path string) string

	// Open your Extension's options page, if possible.
	//
	// The precise behavior may depend on your manifest's options_ui or
	// options_page key, or what Chrome happens to support at the time. For
	// example, the page may be opened in a new tab, within chrome://extensions,
	// within an App, or it may just focus an open options page.  It will never
	// cause the caller page to reload.
	//
	// If your Extension does not declare an options page, or Chrome failed to
	// create one for some other reason, the callback will set lastError.
	OpenOptionsPage()

	// Open your Extension's options page, if possible.
	//
	// The precise behavior may depend on your manifest's options_ui or
	// options_page key, or what Chrome happens to support at the time. For
	// example, the page may be opened in a new tab, within chrome://extensions,
	// within an App, or it may just focus an open options page.  It will never
	// cause the caller page to reload.
	//
	// If your Extension does not declare an options page, or Chrome failed to
	// create one for some other reason, the callback will set lastError.
	//
	// Arguments:
	//   cb = A callback function that takes no arguments.
	OpenOptionsPageAsync(cb EmptyCallback)

	// Reloads the app or extension.
	//
	// This method is not supported in kiosk mode.  For kiosk mode, use
	// chrome.runtime.restart() method.
	Reload()

	// Requests an immediate update check be done for this app/extension.
	//
	// Important: Most extensions/apps should not use this method, since chrome
	// already does automatic checks every few hours, and you can listen for the
	// onUpdateAvailable event without needing to call requestUpdateCheck.
	//
	// This method is only appropriate to call in very limited circumstances, such
	// as if your extension/app talks to a backend service, and the backend
	// service has determined that the client extension/app version is very far
	// out of date and you'd like to prompt a user to update. Most other uses of
	// requestUpdateCheck, such as calling it unconditionally based on a repeating
	// timer, probably only serve to waste client, network, and server resources.
	RequestUpdateCheck() (RequestUpdateCheckStatus, UpdateAvailableDetails)

	// Asynchronously requests an immediate update check be done for this
	// app/extension.
	//
	// Important: Most extensions/apps should not use this method, since chrome
	// already does automatic checks every few hours, and you can listen for the
	// onUpdateAvailable event without needing to call requestUpdateCheck.
	//
	// This method is only appropriate to call in very limited circumstances, such
	// as if your extension/app talks to a backend service, and the backend
	// service has determined that the client extension/app version is very far
	// out of date and you'd like to prompt a user to update. Most other uses of
	// requestUpdateCheck, such as calling it unconditionally based on a repeating
	// timer, probably only serve to waste client, network, and server resources.
	//
	// Arguments:
	//   cb = A callback function takes 2 arguments: a RequestUpdateCheckStatus
	//        and an UpdateResponse.
	RequestUpdateCheckAsync(cb RequestUpdateCheckCallback)

	// Restart the ChromeOS device when the app runs in kiosk mode.
	//
	// Otherwise, it's no-op.
	Restart()

	// Restart the ChromeOS device when the app runs in kiosk mode after the given
	// seconds.
	//
	// If called again before the time ends, the reboot will be delayed.
	//
	// If called with a value of -1, the reboot will be cancelled.
	//
	// It's a no-op in non-kiosk mode.
	//
	// It's only allowed to be called repeatedly by the first extension to invoke
	// this API.
	//
	// Arguments:
	//   seconds = Time to wait in seconds before rebooting the device, or -1 to
	//             cancel a scheduled reboot.
	//   cb      = A callback to be invoked when a restart request was
	//             successfully rescheduled.
	RestartAfterDelayAsync(seconds int, cb EmptyCallback)

	// Sends a single message to event listeners within your extension/app or a
	// different extension/app.
	//
	// Similar to connect but only sends a single message, with an optional
	// response.
	//
	// If sending to your extension, the onMessage event will be fired in every
	// frame of your extension (except for the sender's frame), or
	// onMessageExternal, if a different extension.
	//
	// Note that extensions cannot send messages to content scripts using this
	// method.  To send messages to content scripts, use tabs.sendMessage.
	//
	// Arguments:
	//   extID = The ID of the extension/app to send the message to. If omitted,
	//           the message will be sent to your own extension/app. Required if
	//           sending messages from a web page for web messaging.
	//   msg   = The message to send. This message should be a JSON-ifiable
	//           object.
	//   opts  =
	//
	// Returns:
	//   [0] = the response message.
	SendMessage(extID *string, msg interface{}, opts MessageOptions) interface{}

	// Sends a single message to event listeners within your extension/app or a
	// different extension/app.
	//
	// Similar to connect but only sends a single message, with an optional
	// response.
	//
	// If sending to your extension, the onMessage event will be fired in every
	// frame of your extension (except for the sender's frame), or
	// onMessageExternal, if a different extension.
	//
	// Note that extensions cannot send messages to content scripts using this
	// method.  To send messages to content scripts, use tabs.sendMessage.
	//
	// Arguments:
	//   extID = The ID of the extension/app to send the message to. If omitted,
	//           the message will be sent to your own extension/app. Required if
	//           sending messages from a web page for web messaging.
	//   msg   = The message to send. This message should be a JSON-ifiable
	//           object.
	//   opts  =
	//   cb    = A callback function that should take one any/interface typed
	//           argument.  This argument will contain the response message.
	SendMessageAsync(extID *string, msg interface{}, opts MessageOptions, cb AnyCallback)

	// Send a single message to a native application.
	//
	// Arguments:
	//   app = The name of the native messaging host.
	//   msg = The message that will be passed to the native messaging host.
	//
	// Returns:
	//   [0] = This argument will contain the response message.
	SendNativeMessage(app string, msg interface{}) interface{}

	// Send a single message to a native application.
	//
	// Arguments:
	//   app = The name of the native messaging host.
	//   msg = The message that will be passed to the native messaging host.
	//   cb    = A callback function that should take one any/interface typed
	//           argument.  This argument will contain the response message.
	SendNativeMessageAsync(app string, msg interface{}, cb AnyCallback)

	// Sets the URL to be visited upon uninstallation.
	//
	// This may be used to clean up server-side data, do analytics, and implement
	// surveys. Maximum 255 characters.
	//
	// Arguments:
	//   url = URL to be opened after the extension is uninstalled.  This URL must
	//         have an http: or https: scheme. Set an empty string to not open a
	//         new tab upon uninstallation.
	SetUninstallURL(url string)

	// Sets the URL to be visited upon uninstallation.
	//
	// This may be used to clean up server-side data, do analytics, and implement
	// surveys. Maximum 255 characters.
	//
	// Arguments:
	//   url = URL to be opened after the extension is uninstalled.  This URL must
	//         have an http: or https: scheme. Set an empty string to not open a
	//         new tab upon uninstallation.
	//   cb  = A callback that takes no arguments.  Called when the uninstall URL
	//         is set. If the given URL is invalid, lastError will be set.
	SetUninstallURLAsync(url string, cb EmptyCallback)

	// Fired when a Chrome update is available, but isn't installed immediately
	// because a browser restart is required.
	OnBrowserUpdateAvailable() NotificationEvent

	// Fired when a connection is made from either an extension process or a
	// content script (by connect).
	OnConnect() ConnectEvent

	// Fired when a connection is made from another extension (by connect).
	OnConnectExternal() ConnectEvent

	// Fired when a connection is made from a native application.
	//
	// Currently only supported on Chrome OS.
	OnConnectNative() ConnectEvent

	// Fired when the extension is first installed, when the extension is updated
	// to a new version, and when Chrome is updated to a new version.
	OnInstalled() InstallationEvent

	// Fired when a message is sent from either an extension process (by
	// sendMessage) or a content script (by tabs.sendMessage).
	OnMessage() RuntimeMessageEvent

	// Fired when a message is sent from another extension/app (by sendMessage).
	//
	// Cannot be used in a content script.
	OnMessageExternal() RuntimeMessageEvent

	// Fired when an app or the device that it runs on needs to be restarted.
	//
	// The app should close all its windows at its earliest convenient time to let
	// the restart to happen.  If the app does nothing, a restart will be enforced
	// after a 24-hour grace period has passed.
	//
	// Currently, this event is only fired for Chrome OS kiosk apps.
	OnRestartRequired() RestartRequiredEvent

	// Fired when a profile that has this extension installed first starts up.
	//
	// This event is not fired when an incognito profile is started, even if this
	// extension is operating in 'split' incognito mode.
	OnStartup() NotificationEvent

	// Sent to the event page just before it is unloaded.
	//
	// This gives the extension opportunity to do some clean up.  Note that since
	// the page is unloading, any asynchronous operations started while handling
	// this event are not guaranteed to complete.
	//
	// If more activity for the event page occurs before it gets unloaded the
	// onSuspendCanceled event will be sent and the page won't be unloaded.
	OnSuspend() NotificationEvent

	// Sent after onSuspend to indicate that the app won't be unloaded after all.
	OnSuspendCancelled() NotificationEvent

	// Fired when an update is available, but isn't installed immediately because
	// the app is currently running.
	//
	// If you do nothing, the update will be installed the next time the
	// background page gets unloaded, if you want it to be installed sooner you
	// can explicitly call chrome.runtime.reload().  If your extension is using a
	// persistent background page, the background page of course never gets
	// unloaded, so unless you call chrome.runtime.reload() manually in response
	// to this event the update will not get installed until the next time chrome
	// itself restarts. If no handlers are listening for this event, and your
	// extension has a persistent background page, it behaves as if
	// chrome.runtime.reload() is called in response to this event.
	OnUpdateAvailable() UpdateAvailableEvent

	Release()
}

type RequestUpdateCheckCallback = func(RequestUpdateCheckStatus, UpdateAvailableDetails)

type MessageOptions interface {
	// Whether the TLS channel ID will be passed into onMessageExternal for
	// processes that are listening for the connection event.
	IncludeTLSChannelID() OptionalBool
}
