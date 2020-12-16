// +build js,wasm
// +build chrome

package tabs

import (
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type API struct{}

func (A *API) TabIDNone() chrome.TabID {
	panic("implement me")
}

func (A *API) CaptureVisibleTab(id *chrome.WindowID, opts chrome.ImageDetails) (dataURL string) {
	panic("implement me")
}

func (A *API) CaptureVisibleTabAsync(id *chrome.WindowID, opts chrome.ImageDetails, cb chrome.StringCallback) {
	panic("implement me")
}

func (A *API) Connect(id chrome.TabID, info chrome.ConnectInfo) chrome.Port {
	panic("implement me")
}

func (A *API) Create(props chrome.TabCreateProperties) chrome.Tab {
	panic("implement me")
}

func (A *API) CreateAsync(props chrome.TabCreateProperties, cb chrome.TabCallback) {
	panic("implement me")
}

func (A *API) DetectLanguage(id *chrome.TabID) string {
	panic("implement me")
}

func (A *API) DetectLanguageAsync(id *chrome.TabID, cb chrome.StringCallback) {
	panic("implement me")
}

func (A *API) Discard(id chrome.TabID) chrome.Tab {
	panic("implement me")
}

func (A *API) DiscardAsync(id chrome.TabID, cb chrome.TabCallback) {
	panic("implement me")
}

func (A *API) Duplicate(id chrome.TabID) chrome.Tab {
	panic("implement me")
}

func (A *API) DuplicateAsync(id chrome.TabID, cb chrome.TabCallback) {
	panic("implement me")
}

func (A *API) ExecuteScript(id chrome.TabID, props chrome.InjectDetails) []interface{} {
	panic("implement me")
}

func (A *API) ExecuteScriptAsync(id chrome.TabID, props chrome.InjectDetails, cb chrome.AnySliceCallback) {
	panic("implement me")
}

func (A *API) Get(id chrome.TabID) chrome.Tab {
	panic("implement me")
}

func (A *API) GetAsync(id chrome.TabID, cb chrome.TabCallback) {
	panic("implement me")
}

func (A *API) GetAllInWindow(id *chrome.WindowID) []chrome.Tab {
	panic("implement me")
}

func (A *API) GetAllInWindowAsync(id *chrome.WindowID, cb chrome.TabSliceCallback) {
	panic("implement me")
}

func (A *API) GetCurrent() chrome.Tab {
	panic("implement me")
}

func (A *API) GetCurrentAsync(cb chrome.TabCallback) {
	panic("implement me")
}

func (A *API) GetSelected(id *chrome.WindowID) chrome.Tab {
	panic("implement me")
}

func (A *API) GetSelectedAsync(id *chrome.WindowID, cb chrome.TabCallback) {
	panic("implement me")
}

func (A *API) GetZoom(id *chrome.TabID) float32 {
	panic("implement me")
}

func (A *API) GetZoomAsync(id *chrome.TabID, cb chrome.F32Callback) {
	panic("implement me")
}

func (A *API) GetZoomSettings(id *chrome.TabID) chrome.ZoomSettings {
	panic("implement me")
}

func (A *API) GetZoomSettingsAsync(id *chrome.TabID, cb chrome.ZoomSettingsCallback) {
	panic("implement me")
}

func (A *API) GoBack(id *chrome.TabID) {
	panic("implement me")
}

func (A *API) GoBackAsync(id *chrome.TabID, cb chrome.EmptyCallback) {
	panic("implement me")
}

func (A *API) GoForward(id *chrome.TabID) {
	panic("implement me")
}

func (A *API) GoForwardAsync(id *chrome.TabID, cb chrome.EmptyCallback) {
	panic("implement me")
}

func (A *API) Group(opts chrome.GroupOptions) chrome.GroupID {
	panic("implement me")
}

func (A *API) GroupAsync(opts chrome.GroupOptions, cb chrome.GroupIDCallback) {
	panic("implement me")
}

func (A *API) Highlight(h chrome.HighlightInfo) chrome.Window {
	panic("implement me")
}

func (A *API) HighlightAsync(h chrome.HighlightInfo, callback chrome.WindowCallback) {
	panic("implement me")
}

func (A *API) InsertCSS(id *chrome.TabID, details chrome.InjectDetails) {
	panic("implement me")
}

func (A *API) InsertCSSAsync(id *chrome.TabID, details chrome.InjectDetails, cb chrome.EmptyCallback) {
	panic("implement me")
}

func (A *API) MoveSingle(id chrome.TabID, props chrome.TabMoveProperties) chrome.Tab {
	panic("implement me")
}

func (A *API) MoveSingleAsync(id chrome.TabID, props chrome.TabMoveProperties, cb chrome.TabCallback) {
	panic("implement me")
}

func (A *API) MoveMultiple(ids []chrome.TabID, props chrome.TabMoveProperties) []chrome.Tab {
	panic("implement me")
}

func (A *API) MoveMultipleAsync(ids []chrome.TabID, props chrome.TabMoveProperties, cb chrome.TabSliceCallback) {
	panic("implement me")
}

func (A *API) Query(info chrome.TabQueryInfo) []chrome.Tab {
	panic("implement me")
}

func (A *API) QueryAsync(info chrome.TabQueryInfo, callback chrome.TabSliceCallback) {
	panic("implement me")
}

func (A *API) Reload(id chrome.TabID, properties chrome.TabReloadProperties) {
	panic("implement me")
}

func (A *API) ReloadAsync(id chrome.TabID, properties chrome.TabReloadProperties, callback chrome.EmptyCallback) {
	panic("implement me")
}

func (A *API) RemoveSingle(id chrome.TabID) {
	panic("implement me")
}

func (A *API) RemoveSingleAsync(id chrome.TabID, callback chrome.EmptyCallback) {
	panic("implement me")
}

func (A *API) RemoveMultiple(ids []chrome.TabID) {
	panic("implement me")
}

func (A *API) RemoveMultipleAsync(ids []chrome.TabID, callback chrome.EmptyCallback) {
	panic("implement me")
}

func (A *API) RemoveCSS(id *chrome.TabID, details chrome.DeleteInjectionDetails) {
	panic("implement me")
}

func (A *API) RemoveCSSAsync(id *chrome.TabID, details chrome.DeleteInjectionDetails, cb chrome.EmptyCallback) {
	panic("implement me")
}

func (A *API) SendMessage(id chrome.TabID, msg interface{}, opts chrome.SendMessageOptions) interface{} {
	panic("implement me")
}

func (A *API) SendMessageAsync(id chrome.TabID, msg interface{}, opts chrome.SendMessageOptions, responseCB chrome.AnyCallback) {
	panic("implement me")
}

func (A *API) SendRequest(id chrome.TabID, request interface{}) interface{} {
	panic("implement me")
}

func (A *API) SendRequestAsync(id chrome.TabID, request interface{}, responseCB chrome.AnyCallback) {
	panic("implement me")
}

func (A *API) SetZoom(id *chrome.TabID, zoom float32) {
	panic("implement me")
}

func (A *API) SetZoomAsync(id *chrome.TabID, zoom float32, cb chrome.EmptyCallback) {
	panic("implement me")
}

func (A *API) SetZoomSettings(id *chrome.TabID, zoom chrome.ZoomSettings) {
	panic("implement me")
}

func (A *API) SetZoomSettingsAsync(id *chrome.TabID, zoom chrome.ZoomSettings, cb chrome.EmptyCallback) {
	panic("implement me")
}

func (A *API) UngroupSingle(id chrome.TabID) {
	panic("implement me")
}

func (A *API) UngroupSingleAsync(id chrome.TabID, callback chrome.EmptyCallback) {
	panic("implement me")
}

func (A *API) UngroupMultiple(ids []chrome.TabID) {
	panic("implement me")
}

func (A *API) UngroupMultipleAsync(ids []chrome.TabID, callback chrome.EmptyCallback) {
	panic("implement me")
}

func (A *API) Update(id *chrome.TabID, update chrome.TabUpdateProperties) chrome.Tab {
	panic("implement me")
}

func (A *API) UpdateAsync(id *chrome.TabID, update chrome.TabUpdateProperties, cb chrome.TabCallback) {
	panic("implement me")
}

func (A *API) OnActivated() chrome.TabActivatedEvent {
	panic("implement me")
}

func (A *API) OnActiveChanged() chrome.TabSelectionEvent {
	panic("implement me")
}

func (A *API) OnAttached() chrome.TabAttachEvent {
	panic("implement me")
}

func (A *API) OnCreated() chrome.TabCreateEvent {
	panic("implement me")
}

func (A *API) OnDetached() chrome.TabDetachEvent {
	panic("implement me")
}

func (A *API) OnHighlightChanged() chrome.TabHighlightEvent {
	panic("implement me")
}

func (A *API) OnHighlighted() chrome.TabHighlightEvent {
	panic("implement me")
}

func (A *API) OnMoved() chrome.TabMoveEvent {
	panic("implement me")
}

func (A *API) OnRemoved() chrome.TabRemovalEvent {
	panic("implement me")
}

func (A *API) OnReplaced() chrome.TabReplaceEvent {
	panic("implement me")
}

func (A *API) OnSelectionChanged() chrome.TabSelectionEvent {
	panic("implement me")
}

func (A *API) OnUpdated() chrome.TabUpdateEvent {
	panic("implement me")
}

func (A *API) OnZoomChange() chrome.ZoomChangeEvent {
	panic("implement me")
}
