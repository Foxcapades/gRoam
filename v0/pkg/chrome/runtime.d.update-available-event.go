package chrome

type (
	UpdateAvailableEventListener = func(UpdateAvailableDetails)

	UpdateAvailableEvent interface {
		Event

		// AddListener registers an event listener callback to an event.
		AddListener(key interface{}, cb UpdateAvailableEventListener)
	}

	UpdateAvailableDetails interface {
		Version() string
	}
)
