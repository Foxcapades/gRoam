package chrome

type (
	StorageAreaChangeEventListener = func(StorageChange)

	StorageAreaChangeEvent interface {
		Event

		AddListener(StorageAreaChangeEventListener)
	}
)

type (
	StorageChangeEventListener = func(change StorageChange, areaName string)

	StorageChangeEvent interface {
		Event

		AddListener(StorageChangeEventListener)
	}
)
