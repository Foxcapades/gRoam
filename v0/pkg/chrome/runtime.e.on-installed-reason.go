package chrome

// OnInstalledReason
//
// The reason that an event is being dispatched.
type OnInstalledReason string

const (
	OnInstalledReasonInstall            OnInstalledReason = "install"
	OnInstalledReasonUpdate             OnInstalledReason = "update"
	OnInstalledReasonChromeUpdate       OnInstalledReason = "chrome_update"
	OnInstalledReasonSharedModuleUpdate OnInstalledReason = "shared_module_update"
)
