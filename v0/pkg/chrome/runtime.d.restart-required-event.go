package chrome

type (
	RestartRequiredEventListener = func(OnRestartReason)

	RestartRequiredEvent interface {
		Event

		// AddListener registers an event listener callback to an event.
		AddListener(key interface{}, cb RestartRequiredEventListener)
	}
)

