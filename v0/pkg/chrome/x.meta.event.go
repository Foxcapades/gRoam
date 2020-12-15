package chrome

type Event interface {

	// HasListener returns whether any listeners are registered to this event
	// with the given key.
	HasListener(key interface{}) bool

	// HasListener returns whether any listeners are registered to this event
	// with any of the given keys.
	HasListeners(keys []interface{}) bool

	// RemoveListener removes a registered listener from this event by key.
	//
	// If no such listener was registered with this event, this method does
	// nothing.
	RemoveListener(key interface{})

	Release()
}

type NotificationEvent interface {
	Event

	// AddListener adds a no-arg listener that will be called as a notification of
	// the wrapped event type.
	AddListener(key interface{}, cb EmptyCallback)
}
