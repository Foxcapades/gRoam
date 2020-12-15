package chrome

type StorageChange interface {
	// The new value of the item, if there is a new value.
	NewValue() OptionalAny

	// The old value of the item, if there was an old value.
	OldValue() OptionalAny
}