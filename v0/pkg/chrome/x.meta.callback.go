package chrome

type (
	AnyCallback      = func(interface{})
	AnySliceCallback = func([]interface{})
	BoolCallback     = func(bool)
	EmptyCallback    = func()
	F32Callback      = func(float32)
	StringCallback   = func(string)
	U32Callback      = func(uint32)
)
