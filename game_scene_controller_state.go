package main

type GameSceneControllerState interface {
	server
	timerTrigger
	userHandler
}

type (
	userHandler interface {
		userTap(gss *GameServerSession)
	}
	timerTrigger interface {
		timeout()
		countdown()
	}
	server interface {
		disconnect(session *GameServerSession)
	}
	stateEnterer interface {
		onStateEnter()
	}
	stateLeaver interface {
		onStateLeave()
	}
)

type dummyUserHandler struct {
}

func (dummyUserHandler) userTap(*GameServerSession) {
}
