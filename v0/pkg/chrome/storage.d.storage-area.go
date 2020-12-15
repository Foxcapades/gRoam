package chrome

type StorageArea interface {
	// Fired when one or more items change.
	OnChanged() StorageAreaChangeEvent

	// Removes all items from storage.
	Clear()
	ClearAsync()

	// Gets one or more items from storage.
	GetSingle(key string) StorageData
	GetSingleAsync(key string, cb StorageDataCallback)

	// Gets one or more items from storage.
	GetMultiple(keys []string) StorageData
	GetMultipleAsync(keys []string, cb StorageDataCallback)

	// Gets one or more items from storage.
	GetWithDefaults(obj StorageData) StorageData
	GetWithDefaultsAsync(obj StorageData, cb StorageDataCallback)

	// Gets the amount of space (in bytes) being used by one or more items.
	GetBytesInUseSingle(key string) uint32
	GetBytesInUseSingleAsync(key string, cb U32Callback)

	// Gets the amount of space (in bytes) being used by one or more items.
	GetBytesInUseMulti(keys []string) uint32
	GetBytesInUseMultiAsync(keys []string, cb U32Callback)

	// Removes one or more items from storage.
	RemoveSingle(key string)
	RemoveSingleAsync(key string, cb EmptyCallback)

	// Removes one or more items from storage.
	RemoveMultiple(keys []string)
	RemoveMultipleAsync(keys []string, cb EmptyCallback)

	// Sets multiple items.
	Set(StorageData)
	SetAsync(StorageData, EmptyCallback)
}

type (
	StorageData = map[string]interface{}
	StorageDataCallback = func(StorageData)
)