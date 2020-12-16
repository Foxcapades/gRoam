package x

import "sync"

func NewWaitSingle() (out sync.WaitGroup) {
	out.Add(1)
	return
}
