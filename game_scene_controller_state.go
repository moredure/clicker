package main

type GameSceneControllerState interface {
	tap(*GameServerSession)
	timeout()
	countdown()
	disconnect(*GameServerSession)
}

type (
	stateEnterer interface {
		onStateEnter()
	}
	stateLeaver interface {
		onStateLeave()
	}
)

type dummyDisconnect struct {
}

func (dummyDisconnect) disconnect(*GameServerSession) {
}

type dummyCountdown struct {
}

func (dummyCountdown) countdown() {
}

type dummyTimeout struct {
}

func (dummyTimeout) timeout() {
}

type dummyUserHandler struct {
}

func (dummyUserHandler) tap(*GameServerSession) {
}
