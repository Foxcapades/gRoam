package chrome

type ConnectInfo interface {
	// Whether the TLS channel ID will be passed into onConnectExternal for
	// processes that are listening for the connection event.
	IncludeTLSChannelID() OptionalBool

	// Will be passed into onConnect for processes that are listening for the
	// connection event.
	Name() OptionalString
}
