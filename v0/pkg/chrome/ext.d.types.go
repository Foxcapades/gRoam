package chrome

// DeleteInjectionDetails
//
// Details of the CSS to remove. Either the code or the file property must be
// set, but both may not be set at the same time.
type DeleteInjectionDetails interface {
	// If allFrames is true, implies that the CSS should be removed from all
	// frames of current page. By default, it's false and is only removed from the
	// top frame. If true and frameId is set, then the code is removed from the
	// selected frame and all of its child frames.
	AllFrames() OptionalBool

	// CSS code to remove.
	Code() OptionalString

	// The origin of the CSS to remove. Defaults to "author".
	CSSOrigin() OptionalCSSOrigin

	// CSS file to remove.
	File() OptionalString

	// The frame from where the CSS should be removed.
	//
	// Defaults to 0 (the top-level frame).
	FrameID() OptionalFrameID

	// If matchAboutBlank is true, then the code is also removed from about:blank
	// and about:srcdoc frames if your extension has access to its parent
	// document. By default it is false.
	MatchAboutBlank() OptionalBool
}

type ImageDetails interface {
	// The format of the resulting image. Default is "jpeg".
	Format() OptionalImageFormat
	// When format is "jpeg", controls the quality of the resulting image.
	//
	// This value is ignored for PNG images.
	//
	// As quality is decreased, the resulting image will have more visual
	// artifacts, and the number of bytes needed to store it will decrease.
	//
	// Valid range is 0-100
	Quality() OptionalU8
}

type InjectDetails interface {
	// If allFrames is true, implies that the JavaScript or CSS should be injected
	// into all frames of current page.
	//
	// By default, it's false and is only injected into the top frame. If true and
	// frameId is set, then the code is inserted in the selected frame and all of
	// its child frames.
	AllFrames() OptionalBool
	SetAllFrames(bool) InjectDetails

	// JavaScript or CSS code to inject.
	//
	// Warning:
	//
	// Be careful using the code parameter. Incorrect use of it may open your
	// extension to cross site scripting attacks.
	Code() OptionalString
	SetCode(string) InjectDetails

	// The origin of the CSS to inject. This may only be specified for CSS, not
	// JavaScript. Defaults to "author".
	CSSOrigin() OptionalCSSOrigin
	SetCSSOrigin(CSSOrigin) InjectDetails

	// JavaScript or CSS file to inject.
	File() OptionalString
	SetFile(string) InjectDetails

	// The frame where the script or CSS should be injected.
	//
	// Defaults to 0 (the top-level frame).
	FrameID() OptionalFrameID
	SetFrameID(FrameID) InjectDetails

	// If matchAboutBlank is true, then the code is also injected in about:blank
	// and about:srcdoc frames if your extension has access to its parent
	// document. Code cannot be inserted in top-level about:-frames. By default it
	// is false.
	MatchAboutBlank() OptionalBool
	SetMatchAboutBlank(bool) InjectDetails

	// The soonest that the JavaScript or CSS will be injected into the tab.
	//
	// Defaults to "document_idle".
	RunAt() OptionalRunAt
	SetRunAt(RunAt) InjectDetails
}