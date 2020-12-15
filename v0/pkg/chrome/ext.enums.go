package chrome

type CSSOrigin string

const (
	CSSOriginAuthor CSSOrigin = "author"
	CSSOriginUser   CSSOrigin = "user"
)

type OptionalCSSOrigin interface {
	Optional

	Get() CSSOrigin
	Set(CSSOrigin)
	OrElse(CSSOrigin) CSSOrigin
	With(func(CSSOrigin))
}

type ImageFormat string

const (
	ImageFormatJPG ImageFormat = "jpeg"
	ImageFormatPNG ImageFormat = "png"
)

type OptionalImageFormat interface {
	Optional

	Get() ImageFormat
	Set(ImageFormat)
	OrElse(ImageFormat) ImageFormat
	With(func(ImageFormat))
}

type RunAt string

const (
	RunAtDocumentStart RunAt = "document_start"
	RunAtDocumentEnd   RunAt = "document_end"
	RunAtDocumentIdle  RunAt = "document_idle"
)

type OptionalRunAt interface {
	Optional

	Get() RunAt
	Set(RunAt)
	OrElse(RunAt) RunAt
	With(func(RunAt))
}