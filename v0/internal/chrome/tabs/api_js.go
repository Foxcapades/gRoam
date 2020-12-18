// +build js,wasm
// +build chrome

package tabs

import (
	"syscall/js"

	"github.com/foxcapades/groam/v0/internal/chrome/x"
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type API struct {
	root js.Value

	onActivated        chrome.TabActivatedEvent
	onActiveChanged    chrome.TabSelectionEvent
	onAttached         chrome.TabAttachEvent
	onCreated          chrome.TabCreateEvent
	onDetached         chrome.TabDetachEvent
	onHighlightChanged chrome.TabHighlightEvent
	onHighlighted      chrome.TabHighlightEvent
	onMoved            chrome.TabMoveEvent
	onRemoved          chrome.TabRemovalEvent
	onReplaced         chrome.TabReplaceEvent
	onSelectionChanged chrome.TabSelectionEvent
	onUpdated          chrome.TabUpdateEvent
	onZoomChange       chrome.ZoomChangeEvent
}

func (a *API) TabIDNone() chrome.TabID {
	return chrome.TabID(a.root.Get(x.JsKeyTabIDNone).Int())
}

func (a *API) CaptureVisibleTab(id *chrome.WindowID, opts chrome.ImageDetails) (dataURL string) {
	panic("implement me")
}

func (a *API) CaptureVisibleTabAsync(id *chrome.WindowID, opts chrome.ImageDetails, cb chrome.StringCallback) {
	panic("implement me")
}

func (a *API) Connect(id chrome.TabID, info chrome.ConnectInfo) chrome.Port {
	return NewPort(a.root.Call(x.JsFnCaptureVisibleTab))
}

func (a *API) Create(props chrome.TabCreateProperties) chrome.Tab {
	panic("implement me")
}

func (a *API) CreateAsync(props chrome.TabCreateProperties, cb chrome.TabCallback) {
	panic("implement me")
}

func (a *API) DetectLanguage(id *chrome.TabID) string {
	panic("implement me")
}

func (a *API) DetectLanguageAsync(id *chrome.TabID, cb chrome.StringCallback) {
	panic("implement me")
}

func (a *API) Discard(id chrome.TabID) chrome.Tab {
	panic("implement me")
}

func (a *API) DiscardAsync(id chrome.TabID, cb chrome.TabCallback) {
	panic("implement me")
}
/home/ellie/Code/go/src/github.com/gojp
func (a *API) Duplicate(id chrome.TabID) chrome.Tab {
	panic("implement me")
}

func (a *API) DuplicateAsync(id chrome.TabID, cb chrome.TabCallback) {
	panic("implement me")
}

func (a *API) ExecuteScript(id chrome.TabID, props chrome.InjectDetails) []interface{} {
	panic("implement me")
}

func (a *API) ExecuteScriptAsync(id chrome.TabID, props chrome.InjectDetails, cb chrome.AnySliceCallback) {
	panic("implement me")
}

func (a *API) Get(id chrome.TabID) chrome.Tab {
	panic("implement me")
}

func (a *API) GetAsync(id chrome.TabID, cb chrome.TabCallback) {
	panic("implement me")
}

func (a *API) GetAllInWindow(id *chrome.WindowID) []chrome.Tab {
	panic("implement me")
}

func (a *API) GetAllInWindowAsync(id *chrome.WindowID, cb chrome.TabSliceCallback) {
	panic("implement me")
}

func (a *API) GetCurrent() chrome.Tab {
	panic("implement me")
}

func (a *API) GetCurrentAsync(cb chrome.TabCallback) {
	panic("implement me")
}

func (a *API) GetSelected(id *chrome.WindowID) chrome.Tab {
	panic("implement me")
}

func (a *API) GetSelectedAsync(id *chrome.WindowID, cb chrome.TabCallback) {
	panic("implement me")
}

func (a *API) GetZoom(id *chrome.TabID) float32 {
	panic("implement me")
}

func (a *API) GetZoomAsync(id *chrome.TabID, cb chrome.F32Callback) {
	panic("implement me")
}

func (a *API) GetZoomSettings(id *chrome.TabID) chrome.ZoomSettings {
	panic("implement me")
}

func (a *API) GetZoomSettingsAsync(id *chrome.TabID, cb chrome.ZoomSettingsCallback) {
	panic("implement me")
}

func (a *API) GoBack(id *chrome.TabID) {
	panic("implement me")
}

func (a *API) GoBackAsync(id *chrome.TabID, cb chrome.EmptyCallback) {
	panic("implement me")
}

func (a *API) GoForward(id *chrome.TabID) {
	panic("implement me")
}

func (a *API) GoForwardAsync(id *chrome.TabID, cb chrome.EmptyCallback) {
	panic("implement me")
}

func (a *API) Group(opts chrome.GroupOptions) chrome.GroupID {
	panic("implement me")
}

func (a *API) GroupAsync(opts chrome.GroupOptions, cb chrome.GroupIDCallback) {
	panic("implement me")
}

func (a *API) Highlight(h chrome.TabHighlightInfo) chrome.Window {
	panic("implement me")
}

func (a *API) HighlightAsync(h chrome.TabHighlightInfo, callback chrome.WindowCallback) {
	panic("implement me")
}

func (a *API) InsertCSS(id *chrome.TabID, details chrome.InjectDetails) {
	panic("implement me")
}

func (a *API) InsertCSSAsync(id *chrome.TabID, details chrome.InjectDetails, cb chrome.EmptyCallback) {
	panic("implement me")
}

func (a *API) MoveSingle(id chrome.TabID, props chrome.TabMoveProperties) chrome.Tab {
	panic("implement me")
}

func (a *API) MoveSingleAsync(id chrome.TabID, props chrome.TabMoveProperties, cb chrome.TabCallback) {
	panic("implement me")
}

func (a *API) MoveMultiple(ids []chrome.TabID, props chrome.TabMoveProperties) []chrome.Tab {
	panic("implement me")
}

func (a *API) MoveMultipleAsync(ids []chrome.TabID, props chrome.TabMoveProperties, cb chrome.TabSliceCallback) {
	panic("implement me")
}

func (a *API) Query(info chrome.TabQueryInfo) []chrome.Tab {
	panic("implement me")
}

func (a *API) QueryAsync(info chrome.TabQueryInfo, cb chrome.TabSliceCallback) {
	panic("implement me")
}

func (a *API) Reload(id chrome.TabID, props chrome.TabReloadProperties) {
	panic("implement me")
}

func (a *API) ReloadAsync(id chrome.TabID, props chrome.TabReloadProperties, cb chrome.EmptyCallback) {
	panic("implement me")
}

func (a *API) RemoveSingle(id chrome.TabID) {
	panic("implement me")
}

func (a *API) RemoveSingleAsync(id chrome.TabID, cb chrome.EmptyCallback) {
	panic("implement me")
}

func (a *API) RemoveMultiple(ids []chrome.TabID) {
	panic("implement me")
}

func (a *API) RemoveMultipleAsync(ids []chrome.TabID, callback chrome.EmptyCallback) {
	panic("implement me")
}

func (a *API) RemoveCSS(id *chrome.TabID, details chrome.DeleteInjectionDetails) {
	panic("implement me")
}

func (a *API) RemoveCSSAsync(id *chrome.TabID, details chrome.DeleteInjectionDetails, cb chrome.EmptyCallback) {
	panic("implement me")
}

func (a *API) SendMessage(id chrome.TabID, msg interface{}, opts chrome.SendMessageOptions) (out interface{}) {
	wg := x.NewWaitSingle()
	fn := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		out = args[0]
		wg.Done()

		return nil
	})
	defer fn.Release()

	a.root.Call(x.JsFnSendMessage, id, msg, SendMessageOptionsToJS(opts), fn)
	wg.Wait()

	return
}

func (a *API) SendMessageAsync(id chrome.TabID, msg interface{}, opts chrome.SendMessageOptions, cb chrome.AnyCallback) {
	fn := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		if cb != nil {
			cb(args[0])
		}

		return nil
	})
	defer fn.Release()

	a.root.Call(x.JsFnSendMessage, id, msg, SendMessageOptionsToJS(opts), fn)
}

func (a *API) SendRequest(id chrome.TabID, request interface{}) (out interface{}) {
	wg := x.NewWaitSingle()
	fn := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		out = args[0]
		wg.Done()

		return nil
	})
	defer fn.Release()

	a.root.Call(x.JsFnSendRequest, id, request, fn)
	wg.Wait()

	return
}

func (a *API) SendRequestAsync(id chrome.TabID, request interface{}, cb chrome.AnyCallback) {
	fn := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		if cb != nil {
			cb(args[0])
		}

		return nil
	})
	defer fn.Release()

	a.root.Call(x.JsFnSendRequest, id, request, fn)
}

func (a *API) SetZoom(id *chrome.TabID, zoom float32) {
	wg := x.NewWaitSingle()
	fn := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		wg.Done()
		return nil
	})
	defer fn.Release()

	if id == nil {
		a.root.Call(x.JsFnSetZoom, zoom, fn)
	} else {
		a.root.Call(x.JsFnSetZoom, int(*id), zoom, fn)
	}

	wg.Wait()
}

func (a *API) SetZoomAsync(id *chrome.TabID, zoom float32, cb chrome.EmptyCallback) {
	fn := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		cb()
		return nil
	})
	defer fn.Release()

	if id == nil {
		a.root.Call(x.JsFnSetZoom, zoom, fn)
	} else {
		a.root.Call(x.JsFnSetZoom, int(*id), zoom, fn)
	}
}

func (a *API) SetZoomSettings(id *chrome.TabID, zoom chrome.ZoomSettings) {
	wg := x.NewWaitSingle()
	fn := js.FuncOf(func(js.Value, []js.Value) interface{} {
		wg.Done()
		return nil
	})
	defer fn.Release()

	if id == nil {
		a.root.Call(x.JsFnSetZoomSettings, ZoomSettingsToJS(zoom), fn)
	} else {
		a.root.Call(x.JsFnSetZoomSettings, int(*id), ZoomSettingsToJS(zoom), fn)
	}

	wg.Wait()
}

func (a *API) SetZoomSettingsAsync(
	id *chrome.TabID,
	zoom chrome.ZoomSettings,
	cb chrome.EmptyCallback,
) {
	fn := js.FuncOf(func(js.Value, []js.Value) interface{} {
		cb()
		return nil
	})
	defer fn.Release()

	if id == nil {
		a.root.Call(x.JsFnSetZoomSettings, ZoomSettingsToJS(zoom), fn)
	} else {
		a.root.Call(x.JsFnSetZoomSettings, int(*id), ZoomSettingsToJS(zoom), fn)
	}
}

func (a *API) UngroupSingle(id chrome.TabID) {
	wg := x.NewWaitSingle()
	fn := js.FuncOf(func(js.Value, []js.Value) interface{} {
		wg.Done()
		return nil
	})
	defer fn.Release()

	a.root.Call(x.JsFnUngroup, int(id), fn)
	wg.Wait()
}

func (a *API) UngroupSingleAsync(id chrome.TabID, cb chrome.EmptyCallback) {
	fn := js.FuncOf(func(js.Value, []js.Value) interface{} {
		cb()
		return nil
	})
	defer fn.Release()

	a.root.Call(x.JsFnUngroup, int(id), fn)
}

func (a *API) UngroupMultiple(ids []chrome.TabID) {
	input := x.JsArray.New(len(ids))
	for i := range ids {
		input.SetIndex(i, int(ids[i]))
	}

	wg := x.NewWaitSingle()
	fn := js.FuncOf(func(js.Value, []js.Value) interface{} {
		wg.Done()
		return nil
	})
	defer fn.Release()

	a.root.Call(x.JsFnUngroup, input, fn)
	wg.Wait()
}

func (a *API) UngroupMultipleAsync(ids []chrome.TabID, cb chrome.EmptyCallback) {
	input := x.JsArray.New(len(ids))
	for i := range ids {
		input.SetIndex(i, int(ids[i]))
	}

	fn := js.FuncOf(func(js.Value, []js.Value) interface{} {
		cb()
		return nil
	})
	defer fn.Release()

	a.root.Call(x.JsFnUngroup, input, fn)

}

func (a *API) Update(id *chrome.TabID, update chrome.TabUpdateProperties) (tab chrome.Tab) {
	wg := x.NewWaitSingle()
	fn := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		tab = NewTab(args[0])
		wg.Done()
		return nil
	})
	defer fn.Release()

	if id == nil {
		a.root.Call(x.JsFnUpdate, UpdatePropertiesToJS(update), fn)
	} else {
		a.root.Call(x.JsFnUpdate, int(*id), UpdatePropertiesToJS(update), fn)
	}

	wg.Wait()

	return
}

func (a *API) UpdateAsync(id *chrome.TabID, update chrome.TabUpdateProperties, cb chrome.TabCallback) {
	fn := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		cb(NewTab(args[0]))
		return nil
	})
	defer fn.Release()

	if id == nil {
		a.root.Call(x.JsFnUpdate, UpdatePropertiesToJS(update), fn)
	} else {
		a.root.Call(x.JsFnUpdate, int(*id), UpdatePropertiesToJS(update), fn)
	}
}

func (a *API) OnActivated() chrome.TabActivatedEvent {
	if a.onActivated == nil {
		a.onActivated = NewActivatedEvent(a.root.Get(x.JsKeyOnActivated))
	}

	return a.onActivated
}

func (a *API) OnActiveChanged() chrome.TabSelectionEvent {
	if a.onActiveChanged == nil {
		a.onActiveChanged = NewSelectionEvent(a.root.Get(x.JsKeyOnActiveChanged))
	}

	return a.onActiveChanged
}

func (a *API) OnAttached() chrome.TabAttachEvent {
	if a.onAttached == nil {
		a.onAttached = NewAttachEvent(a.root.Get(x.JsKeyOnAttached))
	}

	return a.onAttached
}

func (a *API) OnCreated() chrome.TabCreateEvent {
	if a.onCreated == nil {
		a.onCreated = NewCreateEvent(a.root.Get(x.JsKeyOnCreated))
	}

	return a.onCreated
}

func (a *API) OnDetached() chrome.TabDetachEvent {
	if a.onDetached == nil {
		a.onDetached = NewDetachEvent(a.root.Get(x.JsKeyOnDetached))
	}

	return a.onDetached
}

func (a *API) OnHighlightChanged() chrome.TabHighlightEvent {
	if a.onHighlightChanged == nil {
		a.onHighlightChanged = NewHighlightEvent(a.root.Get(x.JsKeyOnHighlightChanged))
	}

	return a.onHighlightChanged
}

func (a *API) OnHighlighted() chrome.TabHighlightEvent {
	if a.onHighlighted == nil {
		a.onHighlighted = NewHighlightEvent(a.root.Get(x.JsKeyOnHighlighted))
	}

	return a.onHighlighted
}

func (a *API) OnMoved() chrome.TabMoveEvent {
	if a.onMoved == nil {
		a.onMoved = NewMoveEvent(a.root.Get(x.JsKeyOnMoved))
	}

	return a.onMoved
}

func (a *API) OnRemoved() chrome.TabRemovalEvent {
	if a.onRemoved == nil {
		a.onRemoved = NewRemoveEvent(a.root.Get(x.JsKeyOnRemoved))
	}

	return a.onRemoved
}

func (a *API) OnReplaced() chrome.TabReplaceEvent {
	if a.onReplaced == nil {
		a.onReplaced = NewReplaceEvent(a.root.Get(x.JsKeyOnReplaced))
	}

	return a.onReplaced
}

func (a *API) OnSelectionChanged() chrome.TabSelectionEvent {
	if a.onSelectionChanged == nil {
		a.onSelectionChanged = NewSelectionEvent(a.root.Get(x.JsKeyOnSelectionChanged))
	}

	return a.onSelectionChanged
}

func (a *API) OnUpdated() chrome.TabUpdateEvent {
	if a.onUpdated == nil {
		a.onUpdated = NewUpdateEvent(a.root.Get(x.JsKeyOnUpdated))
	}

	return a.onUpdated
}

func (a *API) OnZoomChange() chrome.ZoomChangeEvent {
	if a.onZoomChange == nil {
		a.onZoomChange = NewZoomChangeEvent(a.root.Get(x.JsKeyOnZoomChange))
	}

	return a.onZoomChange
}

func (a *API) Release() {
	if a.onActivated != nil {
		a.onActivated.Release()
	}
	if a.onActiveChanged != nil {
		a.onActiveChanged.Release()
	}
	if a.onAttached != nil {
		a.onAttached.Release()
	}
	if a.onCreated != nil {
		a.onCreated.Release()
	}
	if a.onDetached != nil {
		a.onDetached.Release()
	}
	if a.onHighlightChanged != nil {
		a.onHighlightChanged.Release()
	}
	if a.onHighlighted != nil {
		a.onHighlighted.Release()
	}
	if a.onMoved != nil {
		a.onMoved.Release()
	}
	if a.onRemoved != nil {
		a.onRemoved.Release()
	}
	if a.onReplaced != nil {
		a.onReplaced.Release()
	}
	if a.onSelectionChanged != nil {
		a.onSelectionChanged.Release()
	}
	if a.onUpdated != nil {
		a.onUpdated.Release()
	}
	if a.onZoomChange != nil {
		a.onZoomChange.Release()
	}

}
