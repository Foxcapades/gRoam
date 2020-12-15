package chrome

import (
	"github.com/foxcapades/groam/v0/pkg/chrome"
)

type OptionalMessageSender struct {
	v **chrome.MessageSender
}

func (o *OptionalMessageSender) IsPresent() bool {
	return *o.v != nil
}

func (o *OptionalMessageSender) IsAbsent() bool {
	return *o.v == nil
}

func (o *OptionalMessageSender) Clear() {
	*o.v = nil
}

func (o *OptionalMessageSender) Get() chrome.MessageSender {
	return **o.v
}

func (o *OptionalMessageSender) Set(sender chrome.MessageSender) {
	*o.v = &sender
}

func (o *OptionalMessageSender) OrElse(sender chrome.MessageSender) chrome.MessageSender {
	if *o.v == nil {
		return sender
	}

	return **o.v
}

func (o *OptionalMessageSender) With(f func(chrome.MessageSender)) {
	if *o.v != nil {
		f(**o.v)
	}
}
