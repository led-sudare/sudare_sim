package util

import "time"

type InlineTicker struct {
	time     time.Time
	duration time.Duration
}

func NewInlineTicker(duration time.Duration) *InlineTicker {
	i := InlineTicker{
		time:     time.Now(),
		duration: duration,
	}
	return &i
}

func (i *InlineTicker) DoIfFire(callback func()) {

	d := time.Since(i.time)
	if d > i.duration {
		callback()
		i.time = time.Now()
	}
}
