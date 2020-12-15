package chrome

type TabID int32

type OptionalTabID interface {
	Optional

	Get() TabID
	Set(TabID)
	OrElse(TabID) TabID
	With(func(TabID))
}

type (
	GroupID int32

	OptionalGroupID interface {
		Optional

		Get() GroupID
		Set(GroupID)
		OrElse(GroupID) GroupID
		With(func(GroupID))
	}

	GroupIDCallback = func(GroupID)
)

// MutedInfo
//
// The tab's muted state and the reason for the last state change.
type MutedInfo interface {
	// The ID of the extension that changed the muted state.
	//
	// Not set if an extension was not the reason the muted state last changed.
	ExtensionID() OptionalString

	// Whether the tab is muted (prevented from playing sound).
	//
	// The tab may be muted even if it has not played or is not currently playing
	// sound.
	//
	// Equivalent to whether the 'muted' audio indicator is showing.
	Muted() bool

	// The reason the tab was muted or unmuted. Not set if the tab's mute state
	// has never been changed.
	Reason() OptionalMutedInfoReason
}

type OptionalMutedInfo interface {
	Optional

	Get() MutedInfo
	Set(MutedInfo)
	OrElse(MutedInfo)
	With(func(MutedInfo))
}

type (
	TabCallback      = func(Tab)
	TabSliceCallback = func([]Tab)
)
type Tab interface {
	// Whether the tab is active in its window.
	//
	// Does not necessarily mean the window is focused.
	Active() bool

	// Whether the tab has produced sound over the past couple of seconds (but it
	// might not be heard if also muted).
	//
	// Equivalent to whether the 'speaker audio' indicator is showing.
	Audible() OptionalBool

	// Whether the tab can be discarded automatically by the browser when
	// resources are low.
	AutoDiscardable() bool

	// Whether the tab is discarded. A discarded tab is one whose content has been
	// unloaded from memory, but is still visible in the tab strip. Its content is
	// reloaded the next time it is activated.
	Discarded() bool

	// The URL of the tab's favicon. This property is only present if the
	// extension's manifest includes the "tabs" permission. It may also be an
	// empty string if the tab is loading.
	FavIconURL() OptionalString

	// The ID of the group that the tab belongs to.
	GroupID() GroupID

	// The height of the tab in pixels.
	Height() OptionalHeight

	// Whether the tab is highlighted.
	Highlighted() bool

	// The ID of the tab. Tab IDs are unique within a browser session. Under some
	// circumstances a tab may not be assigned an ID; for example, when querying
	// foreign tabs using the sessions API, in which case a session ID may be
	// present. Tab ID can also be set to chrome.tabs.TAB_ID_NONE for apps and
	// devtools windows.
	TabID() OptionalTabID

	// Whether the tab is in an incognito window.
	Incognito() bool

	// The zero-based index of the tab within its window.
	Index() uint16

	// The tab's muted state and the reason for the last state change.
	MutedInfo() OptionalMutedInfo

	// The ID of the tab that opened this tab, if any.
	//
	// This property is only present if the opener tab still exists.
	OpenerTabID() OptionalTabID

	// The URL the tab is navigating to, before it has committed.
	//
	// This property is only present if the extension's manifest includes the
	// "tabs" permission and there is a pending navigation.
	PendingURL() OptionalString

	// Whether the tab is pinned.
	Pinned() bool

	// Whether the tab is selected.
	Selected() bool

	// The session ID used to uniquely identify a tab obtained from the sessions
	// API.
	SessionID() OptionalSessionID

	// The tab's loading status.
	Status() OptionalTabStatus

	// The title of the tab.
	//
	// This property is only present if the extension's manifest includes the
	// "tabs" permission.
	Title() OptionalString

	// The last committed URL of the main frame of the tab.
	//
	// This property is only present if the extension's manifest includes the
	// "tabs" permission and may be an empty string if the tab has not yet
	// committed.
	//
	// See also Tab.pendingUrl.
	URL() OptionalString

	// The width of the tab in pixels.
	Width() OptionalWidth

	// The ID of the window that contains the tab.
	WindowID() WindowID
}

type OptionalTab interface {
	Optional

	Get() Tab
	Set(Tab)
	OrElse(Tab)
	With(func(Tab))
}

type OptionalTabSlice interface {
	Optional

	Get() []Tab
	Set([]Tab)
	OrElse([]Tab)
	With(func([]Tab))
}

// ZoomSettings
//
// Defines how zoom changes in a tab are handled and at what scope.
type ZoomSettings interface {
	// Used to return the default zoom level for the current tab in calls to
	// tabs.getZoomSettings.
	DefaultZoomFactor() OptionalF32

	// Defines how zoom changes are handled, i.e., which entity is responsible for
	// the actual scaling of the page; defaults to automatic.
	Mode() OptionalZoomSettingsMode

	// Defines whether zoom changes persist for the page's origin, or only take
	// effect in this tab; defaults to per-origin when in automatic mode, and
	// per-tab otherwise.
	Scope() OptionalZoomSettingsScope
}

type ZoomSettingsCallback = func(ZoomSettings)

type TabCreateProperties interface {
	// Whether the tab should become the active tab in the window.
	//
	// Does not affect whether the window is focused (see windows.update).
	//
	// Defaults to true.
	Active() OptionalBool
	SetActive(bool) TabCreateProperties

	// The position the tab should take in the window.
	//
	// The provided value is clamped to between zero and the number of tabs in the
	// window.
	Index() OptionalU16
	SetIndex(uint16) TabCreateProperties

	// The ID of the tab that opened this tab.
	//
	// If specified, the opener tab must be in the same window as the newly
	// created tab.
	OpenerTabID() OptionalTabID
	SetOpenerTabID(TabID) TabCreateProperties

	// Whether the tab should be pinned.
	//
	// Defaults to false.
	Pinned() OptionalBool
	SetPinned(bool) TabCreateProperties

	// Whether the tab should become the selected tab in the window.
	//
	// Defaults to true.
	Selected() OptionalBool
	SetSelected(bool) TabCreateProperties

	// The URL to initially navigate the tab to.
	//
	// Fully-qualified URLs must include a scheme (i.e., 'http://www.google.com',
	// not 'www.google.com').
	//
	// Relative URLs are relative to the current page within the extension.
	//
	// Defaults to the New Tab Page.
	URL() OptionalString
	SetURL(string) TabCreateProperties

	// The window in which to create the new tab.
	//
	// Defaults to the current window.
	WindowID() OptionalWindowID
	SetWindowID(WindowID) TabCreateProperties
}

type GroupOptions interface {
	// Configurations for creating a group. Cannot be used if groupId is already
	// specified.
	CreateProperties() OptionalGroupCreateProperties
	SetCreateProperties(GroupCreateProperties)

	GroupID() OptionalGroupID
	SetGroupID(GroupID)

	TabIDs() []TabID
	SetTabIDs([]TabID)
	AddTabID(TabID)
}

type GroupCreateProperties interface {
	// The window of the new group. Defaults to the current window.
	WindowID() OptionalWindowID
	SetWindowID() GroupCreateProperties
}

type OptionalGroupCreateProperties interface {
	Optional

	Get() GroupCreateProperties
	Set(GroupCreateProperties)
	OrElse(GroupCreateProperties) GroupCreateProperties
	With(func(GroupCreateProperties))
}

type TabMoveProperties interface {
	Index() int16
	SetIndex(int16) TabMoveProperties

	WindowID() OptionalWindowID
	SetWindowID(WindowID) TabMoveProperties
}

// Gets all tabs that have the specified properties, or all tabs if no
// properties are specified.
type TabQueryInfo interface {
	// Whether the tabs are active in their windows.
	Active() OptionalBool
	SetActive(bool)

	// Whether the tabs are audible.
	Audible() OptionalBool
	SetAudible(bool) TabQueryInfo

	// Whether the tabs can be discarded automatically by the browser when
	// resources are low.
	AutoDiscardable() OptionalBool
	SetAutoDiscardable(bool) TabQueryInfo

	// Whether the tabs are in the current window.
	CurrentWindow() OptionalBool
	SetCurrentWindow(bool) TabQueryInfo

	// Whether the tabs are discarded. A discarded tab is one whose content has
	// been unloaded from memory, but is still visible in the tab strip.
	// Its content is reloaded the next time it is activated.
	Discarded() OptionalBool
	SetDiscarded(bool) TabQueryInfo

	// The ID of the group that the tabs are in, or tabGroups.TAB_GROUP_ID_NONE
	// for ungrouped tabs.
	GroupID() OptionalGroupID
	SetGroupID(GroupID) TabQueryInfo

	// Whether the tabs are highlighted.
	Highlighted() OptionalBool
	SetHighlighted(bool) TabQueryInfo

	// The position of the tabs within their windows.
	Index() OptionalU16
	SetIndex(uint16) TabQueryInfo

	// Whether the tabs are in the last focused window.
	LastFocusedWindow() OptionalBool
	SetLastFocusedWindow(bool) TabQueryInfo

	// Whether the tabs are muted.
	Muted() OptionalBool
	SetMuted(bool) TabQueryInfo

	// Whether the tabs are pinned.
	Pinned() OptionalBool
	SetPinned(bool) TabQueryInfo

	// The tab loading status.
	Status() OptionalTabStatus
	SetStatus(TabStatus) TabQueryInfo

	// Match page titles against a pattern. This property is ignored if the
	// extension does not have the "tabs" permission.
	Title() OptionalString
	SetTitle(string) TabQueryInfo

	// Match tabs against one or more URL patterns. Fragment identifiers are not
	// matched. This property is ignored if the extension does not have the "tabs"
	// permission.
	URL() OptionalStringSlice
	SetURL([]string) TabQueryInfo
	AddURL(string) TabQueryInfo

	// The ID of the parent window, or windows.WINDOW_ID_CURRENT for the current
	// window.
	WindowID() OptionalWindowID
	SetWindowID(WindowID) TabQueryInfo

	// The type of window the tabs are in.
	WindowType() OptionalWindowType
	SetWindowType(WindowType) TabQueryInfo
}

type TabReloadProperties interface {
	BypassCache() OptionalBool
	SetBypassCache(bool) TabReloadProperties
}

type SendMessageOptions interface {
	FrameID() OptionalFrameID
	SetFrameID(FrameID) SendMessageOptions
}

type TabUpdateProperties interface {
	// Whether the tab should be active. Does not affect whether the window is
	// focused (see windows.update).
	Active() OptionalBool
	SetActive(bool) TabUpdateProperties

	// Whether the tab should be discarded automatically by the browser when
	// resources are low.
	AutoDiscardable() OptionalBool
	SetAutoDiscardable(bool) TabUpdateProperties

	// Adds or removes the tab from the current selection.
	Highlighted() OptionalBool
	SetHighlighted(bool) TabUpdateProperties

	// Whether the tab should be muted.
	Muted() OptionalBool
	SetMuted(bool) TabUpdateProperties

	// The ID of the tab that opened this tab. If specified, the opener tab must
	// be in the same window as this tab.
	OpenerTabID() OptionalTabID
	SetOpener(TabID) TabUpdateProperties

	// Whether the tab should be pinned.
	Pinned() OptionalBool
	SetPinned(bool) TabUpdateProperties

	// Whether the tab should be selected.
	Selected() OptionalBool
	SetSelected(bool) TabUpdateProperties

	// A URL to navigate the tab to. JavaScript URLs are not supported; use
	// executeScript instead.
	URL() OptionalString
	SetURL(string) TabUpdateProperties
}