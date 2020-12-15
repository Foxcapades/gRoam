package chrome

type (
	RuntimeMessageEventListener = func(msg interface{}, sender MessageSender, cb RuntimeMessageResponseCallback)

	// TODO: docs say this returns nothing... i don't believe that
	RuntimeMessageResponseCallback = func()

	RuntimeMessageEvent interface {
		Event

		// AddListener registers an event listener callback to an event.
		AddListener(key interface{}, cb RuntimeMessageEventListener)
	}
)
