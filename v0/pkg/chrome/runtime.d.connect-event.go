package chrome

type (
	PortEventListener = func(port Port)

	ConnectEvent interface {
		Event

		// AddListener registers an event listener callback to an event.
		AddListener(key interface{}, cb PortEventListener)
	}
)
