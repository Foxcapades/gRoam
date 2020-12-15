package chrome

type (
	PortMessageEventListener = func(any interface{}, port Port)

	PortMessageEvent interface {
		Event

		// AddListener registers an event listener callback to an event.
		AddListener(key interface{}, cb PortMessageEventListener)
	}
)
