package chrome


type Browser interface {
	// AccessibilityFeatures() AccessibilityFeaturesAPI
	// Alarms() AlarmAPI
	// Bookmarks() BookmarkAPI
	// BrowserAction() BrowserActionAPI
	// BrowsingData() BrowsingDataAPI
	// CertificateProvider() CertificateProviderAPI
	// Commands() CommandAPI
	// ContentSettings() ContentSettingsAPI
	// ContextMenus() ContextMenuAPI
	// Cookies() CookieAPI
	// Debugger() DebuggerAPI
	// DeclarativeContent() DeclarativeContentAPI
	// DeclarativeNetRequest() DeclarativeNetRequestAPI
	// DesktopCapture() DesktopCaptureAPI
	// DevTools() DevToolsAPI
	// DocumentScan() DocumentScanAPI
	// Downloads() DownloadsAPI
	// Enterprise() EnterpriseAPI
	// Events() EventAPI
	// Extension() ExtensionAPI
	// ExtensionTypes() ExtensionTypesAPI
	// FileBrowserHandler() FileBrowserHandlerAPI
	// FileSystemProvider() FileSystemProviderAPI
	// FontSettings() FontSettingsAPI
	// GoogleCloudMessaging() GoogleCloudMessagingAPI
	// History() HistoryAPI
	// I18N() I18NAPI
	// Identity() IdentityAPI
	// Input() InputAPI
	// InstanceID() InstanceIDAPI
	// LoginState() LoginStateAPI
	// Management() ManagementAPI
	// Notifications() NotificationsAPI
	// OmniBox() OmniBoxAPI
	// PageAction() PageActionAPI
	// PageCapture() PageCaptureAPI
	// Permissions() PermissionsAPI
	// PlatformKeys() PlatformKeysAPI
	// Power() PowerAPI
	// PrinterProvider() PrinterProviderAPI
	// Printing() PrintingAPI
	// PrintingMetrics() PrintingMetricsAPI
	// Privacy() PrivacyAPI
	// Proxy() ProxyAPI
	Runtime() RuntimeAPI
	// Search() SearchAPI
	// Sessions() SessionsAPI
	Storage() StorageAPI
	// System() SystemAPI
	// TabCapture() TabCaptureAPI
	Tabs() TabAPI
	// TopSites() TopSitesAPI
	// TextToSpeech() TextToSpeechAPI
	// TTSEngine() TTSEngineAPI
	// Types() TypeAPI
	// VPNProvider() VPNProviderAPI
	// Wallpaper() WallpaperAPI
	// WebNavigation() WebNavigationAPI
	// WebRequest() WebRequestAPI

	// Use the chrome.windows API to interact with browser windows.
	//
	// You can use this API to create, modify, and rearrange windows in the
	// browser.
	Windows() WindowAPI
}
