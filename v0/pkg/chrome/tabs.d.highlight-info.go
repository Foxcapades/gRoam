package chrome

type HighlightInfo interface {
	// One or more tab indices to highlight.
	Tabs() []uint16

	// The window that contains the tabs.
	WindowID() OptionalWindowID
}
