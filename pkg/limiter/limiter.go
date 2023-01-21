package limiter

import (
	"errors"
	"sync/atomic"
)

type Limiter struct {
	maxEvents int64
	curEvents atomic.Int64
}

var errMaxLimitReached = errors.New("max number of requests reached")

func NewLimiter(maxEvents int64) *Limiter {
	return &Limiter{maxEvents: maxEvents}
}

func (l *Limiter) AddTask() error {
	if l.curEvents.Load() >= l.maxEvents {
		return errMaxLimitReached
	}
	l.curEvents.Add(1)
	return nil
}

func (l *Limiter) DoneTask() {
	if l.curEvents.Load() > 0 {
		l.curEvents.Add(-1)
	}
}
