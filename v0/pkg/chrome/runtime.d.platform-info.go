package chrome

type PlatformInfoCallback = func(PlatformInfo)

// PlatformInfo
//
// An object containing information about the current platform.
type PlatformInfo interface {
	// Arch
	//
	// The machine's processor architecture.
	Arch() PlatformArch

	// Nacl_arch
	//
	// The native client architecture. This may be different from arch on some platforms.
	NaclArch() PlatformNACLArch

	// Os
	//
	// The operating system chrome is running on.
	OS() PlatformOS
}
