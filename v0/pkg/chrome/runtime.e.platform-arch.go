package chrome

// PlatformArch
//
// The machine's processor architecture.
type PlatformArch string

const (
	PlatformArchARM    PlatformArch = "arm"
	PlatformArchARM64  PlatformArch = "arm64"
	PlatformArchX8632  PlatformArch = "x86-32"
	PlatformArchX8664  PlatformArch = "x86-64"
	PlatformArchMIPS   PlatformArch = "mips"
	PlatformArchMIPS64 PlatformArch = "mips64"
)
