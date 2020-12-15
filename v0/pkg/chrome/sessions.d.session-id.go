package chrome

type SessionID string

type OptionalSessionID interface {
	Optional

	Get() SessionID
	Set(SessionID)
	OrElse(SessionID)
	With(func(SessionID))
}
