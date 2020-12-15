package chrome

type (
	WindowRemoveEventListener = func(WindowID)

	WindowRemoveEvent interface {
		Event

		AddListener(WindowRemoveEventListener)
	}
)

type (
	WindowFocusEventListener = func(WindowID)

	WindowFocusEvent interface {
		Event

		AddListener(WindowFocusEventListener)
	}
)

type (
	WindowCreatedEventListener = func(Window)

	WindowCreatedEvent interface {
		Event

		AddListener(WindowCreatedEventListener)
	}
)

type (
	WindowBoundsChangeEventListener = func(Window)

	WindowBoundsChangeEvent interface {
		Event

		AddListener(WindowBoundsChangeEventListener)
	}
)
