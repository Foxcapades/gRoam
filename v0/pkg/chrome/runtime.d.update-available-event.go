package chrome

type (
	UpdateAvailableEventListener = func(any interface{}, port Port)

	UpdateAvailableEvent interface {
		Event

		// AddListener registers an event listener callback to an event.
		AddListener(key interface{}, cb UpdateAvailableEventListener)
	}

	UpdateAvailableDetails interface {
		Version() string
	}
)

