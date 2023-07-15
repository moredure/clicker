package main

import "time"

type Timer interface {
	Pause()
	Left() time.Duration
	Start()
}
