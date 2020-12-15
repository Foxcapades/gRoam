package chrome

type RequestUpdateCheckStatus string

const (
	RequestUpdateCheckStatusThrottled       RequestUpdateCheckStatus = "throttled"
	RequestUpdateCheckStatusNoUpdate        RequestUpdateCheckStatus = "no_update"
	RequestUpdateCheckStatusUpdateAvailable RequestUpdateCheckStatus = "update_available"
)
