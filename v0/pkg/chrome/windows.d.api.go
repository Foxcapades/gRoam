package chrome

// chrome.windows
//
// Use the chrome.windows API to interact with browser windows.  You can use
// this API to create, modify, and rearrange windows in the browser.
//
// Manifest
//
// When requested, a windows.Window will contain an array of tabs.Tab objects.
// You must declare the "tabs" permission in your manifest if you require access
// to the url, pendingUrl, title, or favIconUrl properties of tabs.Tab. For
// example:
//
//   {
//     "name": "My extension",
//     ...
//     "permissions": ["tabs"],
//     ...
//   }
//
// The current window
//
// Many functions in the extension system take an optional windowId parameter,
// which defaults to the current window.
//
// The current window is the window that contains the code that is currently
// executing. It's important to realize that this can be different from the
// topmost or focused window.
//
// For example, say an extension creates a few tabs or windows from a single
// HTML file, and that the HTML file contains a call to tabs.query. The current
// window is the window that contains the page that made the call, no matter
// what the topmost window is.
//
// In the case of the event page, the value of the current window falls back to
// the last active window. Under some circumstances, there may be no current
// window for background pages.
//
// Examples
//
// Two windows, each with one tab: https://developer.chrome.com/docs/extensions/reference/windows/windows.png
//
// You can find simple examples of using the windows module in the
// `examples/api/windows` directory. Another example is in the tabs_api.html
// file of the inspector example. For other examples and for help in viewing the
// source code, see Samples.
type WindowAPI interface {
	// The windowId value that represents the current window.
	WindowIDCurrent() WindowID

	// The windowId value that represents the absence of a chrome browser window.
	WindowIDNone() WindowID

	// Creates (opens) a new browser window with any optional sizing, position, or
	// default URL provided.
	Create(data WindowCreateData) Window

	// Creates (opens) a new browser window with any optional sizing, position, or
	// default URL provided.
	CreateAsync(data WindowCreateData, cb WindowCallback)

	// Gets details about a window.
	Get(id WindowID, opts QueryOptions) Window

	// Gets details about a window.
	GetAsync(id WindowID, opts QueryOptions, cb WindowCallback)

	// Gets all windows.
	GetAll(opts QueryOptions) []Window

	// Gets all windows.
	GetAllAsync(opts QueryOptions, cb MultiWindowCallback)

	// Gets the current window.
	GetCurrent() Window

	// Gets the current window.
	GetCurrentAsync(cb WindowCallback)

	// Gets the window that was most recently focused — typically the window 'on
	// top'.
	GetLastFocused() Window

	// Gets the window that was most recently focused — typically the window 'on
	// top'.
	GetLastFocusedAsync(cb WindowCallback)

	// Removes (closes) a window and all the tabs inside it.
	Remove(id WindowID)

	// Removes (closes) a window and all the tabs inside it.
	RemoveAsync(id WindowID, cb EmptyCallback)

	// Updates the properties of a window.
	//
	// Specify only the properties that to be changed; unspecified properties are
	// unchanged.
	Update(id WindowID, upd WindowUpdateData) Window

	// Updates the properties of a window.
	//
	// Specify only the properties that to be changed; unspecified properties are
	// unchanged.
	UpdateAsync(id WindowID, upd WindowUpdateData, cb WindowCallback)

	// Fired when a window has been resized; this event is only dispatched when
	// the new bounds are committed, and not for in-progress changes.
	OnBoundsChanged() WindowBoundsChangeEvent

	// Fired when a window is created.
	OnCreated() WindowCreatedEvent

	// Fired when the currently focused window changes.
	//
	// Returns chrome.windows.WINDOW_ID_NONE if all Chrome windows have lost
	// focus.
	//
	// Note: On some Linux window managers, WINDOW_ID_NONE is always sent
	// immediately preceding a switch from one Chrome window to another.
	OnFocusChanged() WindowFocusEvent

	// Fired when a window is removed (closed).
	OnRemoved() WindowRemoveEvent
}
