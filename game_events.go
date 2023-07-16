package main

type event interface {
	Type() string
}

type CountdownEvent struct {
	Tick int `json:"tick"`
}

func (CountdownEvent) Type() string {
	return "countdown"
}

type StartEvent struct {
}

func (StartEvent) Type() string {
	return "start"
}

type PlayerStateEvent struct {
	State interface{} `json:"state"`
}

func (PlayerStateEvent) Type() string {
	return "player_state"
}

var (
	startEvent      = new(StartEvent)
	countdownEvent3 = &CountdownEvent{Tick: 3}
	countdownEvent2 = &CountdownEvent{Tick: 2}
	countdownEvent1 = &CountdownEvent{Tick: 1}
)
