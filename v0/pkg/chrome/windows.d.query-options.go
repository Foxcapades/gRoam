package chrome

type QueryOptions interface {
	// If true, the Window object has a tabs property that contains a list of the
	// Tab objects.  The Tab objects only contain the url, pendingUrl, title, and
	// favIconUrl properties if the extension's manifest file includes the "tabs"
	// permission.
	Populate() OptionalBool

	// The type of browser window this is. In some circumstances a window may not
	// be assigned a type property; for example, when querying closed windows from
	// the sessions API.
	WindowTypes() OptionalWindowTypeSlice
}