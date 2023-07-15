package main

type GameSceneControllerState interface {
	serverController
	timerController
	userController
}

type (
	userController interface {
		tap(gss *GameServerSession)
	}
	timerController interface {
		timeout()
		countdown()
	}
	serverController interface {
		disconnect(*GameServerSession)
	}
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
