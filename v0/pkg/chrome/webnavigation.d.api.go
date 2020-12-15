package chrome

type FrameID int32

type OptionalFrameID interface {
	Optional

	Get() FrameID
	Set(FrameID)
	OrElse(FrameID) FrameID
	With(func(FrameID))
}
