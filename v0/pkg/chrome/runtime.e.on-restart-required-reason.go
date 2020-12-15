package chrome

// OnRestartRequiredReason
//
// The reason that the event is being dispatched.
//
// 'app_update' is used when the restart is needed because the application is
// updated to a newer version.
//
// 'os_update' is used when the restart is needed because the browser/OS is
// updated to a newer version.
//
// 'periodic' is used when the system runs for more than the permitted uptime
// set in the enterprise policy.
type OnRestartReason string

const (
	OnRestartReasonAppUpdate OnRestartReason = "app_update"
	OnRestartReasonOSUpdate  OnRestartReason = "os_update"
	OnRestartReasonPeriodic  OnRestartReason = "periodic"
)
