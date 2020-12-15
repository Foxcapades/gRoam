package chrome

type (
	TabActivatedListener = func(TabActivatedInfo)

	TabActivatedEvent interface {
		Event

		AddListener(TabActivatedListener)
	}

	TabActivatedInfo interface {
		// The ID of the tab that has become active.
		TabID() TabID

		// The ID of the window the active tab changed inside of.
		WindowID() WindowID
	}
)

type (
	TabSelectionListener = func(TabID, TabSelectInfo)

	TabSelectionEvent interface {
		Event

		AddListener(TabSelectionListener)
	}

	TabSelectInfo interface {
		// The ID of the window the selected tab changed inside of.
		WindowID() WindowID
	}
)

type (
	TabAttachListener = func(TabID, TabAttachInfo)

	TabAttachEvent interface {
		Event

		AddListener(TabAttachListener)
	}

	TabAttachInfo interface {
		NewPosition() uint16
		NewWindowID() WindowID
	}
)

type (
	TabCreateListener = func(Tab)

	TabCreateEvent interface {
		Event
		AddListener(TabCreateListener)
	}
)

type (
	TabDetachListener = func(TabID, TabDetachInfo)

	TabDetachEvent interface {
		Event
		AddListener(TabDetachListener)
	}

	TabDetachInfo interface {
		OldPosition() uint16
		OldWindowID() WindowID
	}
)

type (
	TabHighlightListener = func(TabHighlightChangeInfo)

	TabHighlightEvent interface {
		Event
		AddListener(TabHighlightListener)
	}

	TabHighlightChangeInfo interface {
		TabIDs() TabID
		WindowID() WindowID
	}
)

type (
	TabMoveListener = func(TabID, TabMoveInfo)

	TabMoveEvent interface {
		Event
		AddListener(TabMoveListener)
	}

	TabMoveInfo interface {
		FromIndex() uint16
		ToIndex() uint16
		WindowID() WindowID
	}
)

type (
	TabRemoveListener = func(TabID, TabRemovalInfo)

	TabRemovalEvent interface {
		Event
		AddListener(TabRemoveListener)
	}

	TabRemovalInfo interface {
		// True when the tab was closed because its parent window was closed.
		IsWindowClosing() bool

		// The window whose tab is closed.
		WindowID() WindowID
	}
)

type (
	TabReplaceListener = func(added, removed TabID)

	TabReplaceEvent interface {
		Event
		AddListener(TabReplaceListener)
	}
)

type (
	TabUpdateListener = func(TabID, TabChangeInfo)

	TabUpdateEvent interface {
		Event
		AddListener(TabUpdateListener)
	}

	// Lists the changes to the state of the tab that was updated.
	TabChangeInfo interface {
		// The tab's new audible state.
		Audible() OptionalBool

		// The tab's new auto-discardable state.
		AutoDiscardable() OptionalBool

		// The tab's new discarded state.
		Discarded() OptionalBool

		// The tab's new discarded state.
		FavIconURL() OptionalBool

		// The tab's new group.
		GroupID() OptionalGroupID

		// The tab's new muted state and the reason for the change.
		MutedInfo() OptionalMutedInfo

		// The tab's new pinned state.
		Pinned() OptionalBool

		// The tab's loading status.
		Status() OptionalTabStatus

		// The tab's new title.
		Title() OptionalString

		// The tab's URL if it has changed.
		URL() OptionalString
	}
)

type (
	ZoomChangeListener = func(ZoomChangeInfo)

	ZoomChangeEvent interface {
		Event
		AddListener(ZoomChangeListener)
	}

	ZoomChangeInfo interface {
		NewZoomFactor() float32
		OldZoomFactor() float32
		TabID() TabID
		ZoomSettings() ZoomSettings
	}
)
