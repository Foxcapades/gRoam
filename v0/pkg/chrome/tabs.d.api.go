package chrome

// chrome.tabs
//
// Use the chrome.tabs API to interact with the browser's tab system. You can
// use this API to create, modify, and rearrange tabs in the browser.
//
// Manifest
//
// You can use most chrome.tabs methods and events without declaring any
// permissions in the extension's manifest file. However, if you require access
// to the url, pendingUrl, title, or favIconUrl properties of tabs.Tab, you must
// declare the "tabs" permission in the manifest, as shown below:
//
//   {
//     "name": "My extension",
//     ...
//     "permissions": [
//       "tabs"
//     ],
//     ...
//   }
//
// Examples
//
// Two tabs in a window: https://developer.chrome.com/docs/extensions/reference/tabs/tabs.png
//
// You can find simple examples of manipulating tabs with the chrome.tabs API
// in the examples/api/tabs directory. For other examples and for help in
// viewing the source code, see Samples.
type TabAPI interface {
	// TAB_ID_NONE
	//
	// An ID that represents the absence of a browser tab.
	TabIDNone() TabID

	// Captures the visible area of the currently active tab in the specified
	// window.
	//
	// In order to call this method, the extension must have either the <all_urls>
	// permission or the activeTab permission.
	//
	// In addition to sites that extensions can normally access, this method
	// allows extensions to capture sensitive sites that are otherwise restricted,
	// including chrome:-scheme pages, other extensions' pages, and data: URLs.
	// These sensitive sites can only be captured with the activeTab permission.
	//
	// File URLs may be captured only if the extension has been granted file
	// access.
	//
	// Arguments:
	//   id      = The target window. Defaults to the current window.
	//   options =
	//
	// Returns:
	//   [0] = A data URL that encodes an image of the visible area of the
	//         captured tab. May be assigned to the 'src' property of an HTML
	//         img element for display.
	CaptureVisibleTab(id *WindowID, options ImageDetails) (dataURL string)

	// Captures the visible area of the currently active tab in the specified
	// window.
	//
	// In order to call this method, the extension must have either the <all_urls>
	// permission or the activeTab permission.
	//
	// In addition to sites that extensions can normally access, this method
	// allows extensions to capture sensitive sites that are otherwise restricted,
	// including chrome:-scheme pages, other extensions' pages, and data: URLs.
	// These sensitive sites can only be captured with the activeTab permission.
	//
	// File URLs may be captured only if the extension has been granted file
	// access.
	//
	// Arguments:
	//   id      = The target window. Defaults to the current window.
	//   options =
	//   cb      = A callback that takes a single string param for the dataURL.
	CaptureVisibleTabAsync(id *WindowID, options ImageDetails, cb StringCallback)


	// Connects to the content script(s) in the specified tab.
	//
	// The runtime.onConnect event is fired in each content script running in the
	// specified tab for the current extension.
	//
	// For more details, see Content Script Messaging.
	//
	// Arguments:
	//   id   =
	//   info =
	//
	// Returns:
	//   [0] = A port that can be used to communicate with the content scripts
	//         running in the specified tab. The port's runtime.Port event is
	//         fired if the tab closes or does not exist.
	Connect(id TabID, info ConnectInfo) Port

	// Creates a new tab.
	Create(TabCreateProperties) Tab

	// Creates a new tab.
	CreateAsync(TabCreateProperties, TabCallback)

	// Detects the primary language of the content in a tab.
	//
	// Arguments:
	//   id = Defaults to the active tab of the current window.
	//
	// Returns:
	//   [0] = An ISO language code such as en or fr. For a complete list of
	//         languages supported by this method, see kLanguageInfoTable. The
	//         second to fourth columns are checked and the first non-NULL value
	//         is returned, except for Simplified Chinese for which zh-CN is
	//         returned. For an unknown/undefined language, und is returned.
	DetectLanguage(id *TabID) string
	DetectLanguageAsync(id *TabID, cb StringCallback)

	// Discards a tab from memory.
	//
	// Discarded tabs are still visible on the tab strip and are reloaded when
	// activated.
	//
	// Arguments:
	//   id = The ID of the tab to be discarded. If specified, the tab is
	//        discarded unless it is active or already discarded. If omitted, the
	//        browser discards the least important tab. This can fail if no
	//        discardable tabs exist.
	//
	// Returns:
	//   [0] = The discarded tab, if it was successfully discarded; undefined
	//         otherwise.
	Discard(id TabID) Tab
	DiscardAsync(id TabID, cb TabCallback)

	// Duplicates a tab.
	//
	// Arguments:
	//   id = The ID of the tab to duplicate.
	//
	// Returns:
	//   [0] = Details about the duplicated tab. The Tab object does not contain
	//         url, pendingUrl, title, and favIconUrl if the "tabs" permission has
	//         not been requested.
	Duplicate(id TabID) Tab
	DuplicateAsync(id TabID, cb TabCallback)

	// Injects JavaScript code into a page.
	//
	// For details, see the programmatic injection section of the content scripts
	// doc.
	//
	// Arguments:
	//   id    = The ID of the tab in which to run the script; defaults to the active
	//           tab of the current window.
	//   props = Details of the script to run. Either the code or the file
	//           property must be set, but both may not be set at the same time.
	//
	// Returns:
	//   [0] = The result of the script in every injected frame.
	ExecuteScript(id TabID, props InjectDetails) []interface{}
	ExecuteScriptAsync(id TabID, props InjectDetails, cb AnySliceCallback)

	// Retrieves details about the specified tab.
	//
	// Arguments:
	//   id =
	//
	// Returns:
	//   [0] =
	Get(id TabID) Tab
	GetAsync(id TabID, cb TabCallback)

	// Gets details about all tabs in the specified window.
	//
	// Arguments:
	//   id = Optional window id. Defaults to the current window.
	//
	// Returns:
	//   [0] =
	GetAllInWindow(id *WindowID) []Tab
	GetAllInWindowAsync(id *WindowID, cb TabSliceCallback)

	// Gets the tab that this script call is being made from.
	//
	// May be undefined if called from a non-tab context (for example, a
	// background page or popup view).
	//
	// Returns:
	//   [0] =
	GetCurrent() Tab
	GetCurrentAsync(cb TabCallback)

	// Gets the tab that is selected in the specified window.
	GetSelected(id *WindowID) Tab
	GetSelectedAsync(id *WindowID, cb TabCallback)

	// Gets the current zoom factor of a specified tab.
	GetZoom(id *TabID) float32
	GetZoomAsync(id *TabID, cb F32Callback)

	// Gets the current zoom settings of a specified tab.
	GetZoomSettings(id *TabID) ZoomSettings
	GetZoomSettingsAsync(id *TabID, cb ZoomSettingsCallback)

	// Go back to the previous page, if one is available.
	GoBack(id *TabID)
	GoBackAsync(id *TabID, cb EmptyCallback)

	// Go foward to the next page, if one is available.
	GoForward(id *TabID)
	GoForwardAsync(id *TabID, cb EmptyCallback)

	// Adds one or more tabs to a specified group, or if no group is specified,
	// adds the given tabs to a newly created group.
	Group(opts GroupOptions) GroupID
	GroupAsync(opts GroupOptions, cb GroupIDCallback)

	// Highlights the given tabs and focuses on the first of group.
	//
	// Will appear to do nothing if the specified tab is currently active.
	Highlight(HighlightInfo) Window
	HighlightAsync(HighlightInfo, WindowCallback)

	// Injects CSS into a page. For details, see the programmatic injection
	// section of the content scripts doc.
	InsertCSS(id *TabID, details InjectDetails)
	InsertCSSAsync(id *TabID, details InjectDetails, cb EmptyCallback)

	// Moves one or more tabs to a new position within its window, or to a new
	// window.
	//
	// Note that tabs can only be moved to and from normal
	// (window.type === "normal") windows.
	MoveSingle(id TabID, props TabMoveProperties) Tab
	MoveSingleAsync(id TabID, props TabMoveProperties, cb TabCallback)
	MoveMultiple(ids []TabID, props TabMoveProperties) []Tab
	MoveMultipleAsync(ids []TabID, props TabMoveProperties, cb TabSliceCallback)

	//
	Query(TabQueryInfo) []Tab
	QueryAsync(TabQueryInfo, TabSliceCallback)

	// Reload a tab.
	Reload(TabID, TabReloadProperties)
	ReloadAsync(TabID, TabReloadProperties, EmptyCallback)

	// Closes one or more tabs.
	RemoveSingle(TabID)
	RemoveSingleAsync(TabID, EmptyCallback)
	RemoveMultiple([]TabID)
	RemoveMultipleAsync([]TabID, EmptyCallback)

	// Removes from a page CSS that was previously injected by a call to
	// insertCSS.
	RemoveCSS(id *TabID, details DeleteInjectionDetails)
	RemoveCSSAsync(id *TabID, details DeleteInjectionDetails, cb EmptyCallback)

	// Sends a single message to the content script(s) in the specified tab, with
	// an optional callback to run when a response is sent back.
	//
	// The runtime.onMessage event is fired in each content script running in the
	// specified tab for the current extension.
	SendMessage(id TabID, msg interface{}, opts SendMessageOptions) interface{}
	SendMessageAsync(id TabID, msg interface{}, opts SendMessageOptions, responseCB AnyCallback)

	// Sends a single request to the content script(s) in the specified tab, with
	// an optional callback to run when a response is sent back.
	//
	// The extension.onRequest event is fired in each content script running in
	// the specified tab for the current extension.
	SendRequest(id TabID, request interface{}) interface{}
	SendRequestAsync(id TabID, request interface{}, responseCB AnyCallback)

	// Zooms a specified tab.
	SetZoom(id *TabID, zoom float32)
	SetZoomAsync(id *TabID, zoom float32, cb EmptyCallback)

	// Sets the zoom settings for a specified tab, which define how zoom changes
	// are handled. These settings are reset to defaults upon navigating the tab.
	SetZoomSettings(id *TabID, zoom ZoomSettings)
	SetZoomSettingsAsync(id *TabID, zoom ZoomSettings, cb EmptyCallback)

	// Removes one or more tabs from their respective groups.
	//
	// If any groups become empty, they are deleted.
	UngroupSingle(TabID)
	UngroupSingleAsync(TabID, EmptyCallback)
	UngroupMultiple([]TabID)
	UngroupMultipleAsync([]TabID, EmptyCallback)

	// Modifies the properties of a tab. Properties that are not specified in
	// updateProperties are not modified.
	Update(id *TabID, update TabUpdateProperties) Tab
	UpdateAsync(id *TabID, update TabUpdateProperties, cb TabCallback)

	// Fires when the active tab in a window changes. Note that the tab's URL may
	// not be set at the time this event fired, but you can listen to onUpdated
	// events so as to be notified when a URL is set.
	OnActivated() TabActivatedEvent

	// Fires when the selected tab in a window changes. Note that the tab's URL
	// may not be set at the time this event fired, but you can listen to
	// onUpdated events so as to be notified when a URL is set.
	OnActiveChanged() TabSelectionEvent

	// Fired when a tab is attached to a window; for example, because it was moved
	// between windows.
	OnAttached() TabAttachEvent

	// Fired when a tab is created. Note that the tab's URL may not be set at the
	// time this event is fired, but you can listen to onUpdated events so as to
	// be notified when a URL is set.
	OnCreated() TabCreateEvent

	// Fired when a tab is detached from a window; for example, because it was
	// moved between windows.
	OnDetached() TabDetachEvent

	// Fired when a tab is detached from a window; for example, because it was
	// moved between windows.
	OnHighlightChanged() TabHighlightEvent

	// Fired when the highlighted or selected tabs in a window changes.
	OnHighlighted() TabHighlightEvent

	// Fired when a tab is moved within a window. Only one move event is fired,
	// representing the tab the user directly moved. Move events are not fired
	// for the other tabs that must move in response to the manually-moved tab.
	// This event is not fired when a tab is moved between windows; for details,
	// see onDetached.
	OnMoved() TabMoveEvent

	// Fired when a tab is closed.
	OnRemoved() TabRemovalEvent

	// Fired when a tab is replaced with another tab due to prerendering or
	// instant.
	OnReplaced() TabReplaceEvent

	// Fires when the selected tab in a window changes.
	OnSelectionChanged() TabSelectionEvent

	//
	OnUpdated() TabUpdateEvent

	//
	OnZoomChange() ZoomChangeEvent
}
