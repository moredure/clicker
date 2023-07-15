package main

import "time"

type ReentrantTimer struct {
	left  time.Duration
	start time.Time
	t     *time.Timer
	f     func()
}

func (t *ReentrantTimer) Start() {
	t.start = time.Now()
	if t.t == nil {
		t.t = time.AfterFunc(t.left, t.f)
		return
	}
	t.t.Reset(t.left)
}

func (t *ReentrantTimer) Left() time.Duration {
	left := t.left - time.Since(t.start)
	if left < 0 {
		return 0
	}
	return left
}

func (t *ReentrantTimer) Pause() {
	t.left = t.left - time.Since(t.start)
	t.t.Stop()
}

func NewPausableAfterFunc(d time.Duration, f func()) *ReentrantTimer {
	return &ReentrantTimer{
		left: d,
		f:    f,
	}
}
