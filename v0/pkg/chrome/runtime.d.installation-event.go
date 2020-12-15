package chrome

type (
	InstallationEventListener = func(InstallationEventDetails)

	InstallationEvent interface {
		Event

		// AddListener registers an event listener callback to an event.
		AddListener(key interface{}, cb InstallationEventListener)
	}

	InstallationEventDetails interface {
		// Indicates the ID of the imported shared module extension which updated.
		//
		// This is present only if 'reason' is 'shared_module_update'.
		ID() OptionalString

		// Indicates the previous version of the extension, which has just been
		// updated. This is present only if 'reason' is 'update'.
		PreviousVersion() OptionalString

		// The reason that this event is being dispatched.
		Reason() OnInstalledReason
	}
)
