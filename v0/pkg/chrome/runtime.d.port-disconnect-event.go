package chrome

type (
	PortDisconnectEvent interface {
		Event

		// AddListener registers an event listener callback to an event.
		AddListener(key interface{}, cb PortEventListener)
	}
)
