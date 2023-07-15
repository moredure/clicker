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

type OpponentDisconnectedEvent struct {
}

func (OpponentDisconnectedEvent) Type() string {
	return "opponent_disconnected_event"
}

type OpponentReconnectedEvent struct {
}

func (OpponentReconnectedEvent) Type() string {
	return "opponent_reconnected_event"
}

type MessageEvent struct {
	Source string `json:"source"`
	Body   string `json:"body"`
}

func (*MessageEvent) Type() string {
	return "message_event"
}

var (
	startEvent      = new(StartEvent)
	countdownEvent3 = &CountdownEvent{Tick: 3}
	countdownEvent2 = &CountdownEvent{Tick: 2}
	countdownEvent1 = &CountdownEvent{Tick: 1}
)
