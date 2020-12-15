package chrome

type (
	PortCallback = func(Port, error)
)

// Port
//
// An object which allows two way communication with other pages.
//
// See Long-lived connections for more information.
type Port interface {
	// Disconnect
	//
	// Immediately disconnect the port.
	//
	// Calling disconnect() on an already-disconnected port has no effect.
	//
	// When a port is disconnected, no new events will be dispatched to this port.
	Disconnect()

	// Name
	//
	// The name of the port, as specified in the call to connect.
	Name() string

	// OnDisconnect
	//
	// Fired when the port is disconnected from the other end(s).
	//
	// lastError may be set if the port was disconnected by an error.
	//
	// If the port is closed via disconnect, then this event is only fired on the
	// other end.
	//
	// This event is fired at most once (see also Port lifetime).
	//
	// The first and only parameter to the event handler is this disconnected port.
	OnDisconnect() ConnectEvent

	// OnMessage
	//
	// This event is fired when postMessage is called by the other end of the
	// port.
	//
	// The first parameter is the message,
	// the second parameter is the port that received the message.
	OnMessage() PortMessageEvent

	// PostMessage
	//
	// Send a message to the other end of the port.
	//
	// If the port is disconnected, an error is thrown.
	PostMessage(any interface{})

	// Sender
	//
	// This property will only be present on ports passed to onConnect /
	// onConnectExternal / onConnectNative listeners.
	Sender() OptionalMessageSender

	Release()
}
