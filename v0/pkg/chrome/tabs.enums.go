package chrome

type MutedInfoReason string

const (
	MutedInfoReasonUser      MutedInfoReason = "user"
	MutedInfoReasonCapture   MutedInfoReason = "capture"
	MutedInfoReasonExtension MutedInfoReason = "extension"
)

type OptionalMutedInfoReason interface {
	Optional

	Get() MutedInfoReason
	Set(MutedInfoReason)
	OrElse(MutedInfoReason) MutedInfoReason
	With(func(MutedInfoReason))
}

type TabStatus string

const (
	TabStatusUnloaded TabStatus = "unloaded"
	TabStatusLoading  TabStatus = "loading"
	TabStatusComplete TabStatus = "complete"
)

type OptionalTabStatus interface {
	Optional

	Get() TabStatus
	Set(TabStatus)
	OrElse(TabStatus) TabStatus
	With(func(TabStatus))
}

type ZoomSettingsMode string

const (
	ZoomSettingsModeAutomatic ZoomSettingsMode = "automatic"
	ZoomSettingsModeManual    ZoomSettingsMode = "manual"
	ZoomSettingsModeDisabled  ZoomSettingsMode = "disabled"
)

type OptionalZoomSettingsMode interface {
	Optional

	Get() ZoomSettingsMode
	Set(ZoomSettingsMode)
	OrElse(ZoomSettingsMode) ZoomSettingsMode
	With(func(ZoomSettingsMode))
}

type ZoomSettingsScope string

const (
	ZoomSettingsScopePerOrigin ZoomSettingsScope = "per-origin"
	ZoomSettingsScopePerTab    ZoomSettingsScope = "per-tab"
)

type OptionalZoomSettingsScope interface {
	Optional

	Get() ZoomSettingsScope
	Set(ZoomSettingsScope)
	OrElse(ZoomSettingsScope) ZoomSettingsScope
	With(func(ZoomSettingsScope))
}
