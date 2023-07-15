package main

import (
	"time"
)

type gameSceneControllerCountdown3 struct {
	dummyUserHandler
	dummyTimeout
	gsc *GameSceneController
}

func (currentState *gameSceneControllerCountdown3) onStateEnter() {
	currentState.gsc.reusableTimer = time.AfterFunc(time.Second, currentState.gsc.countdown)
}

func (currentState *gameSceneControllerCountdown3) onStateLeave() {
	currentState.gsc.reusableTimer.Stop()
}

func (currentState *gameSceneControllerCountdown3) countdown() {
	currentState.gsc.gameServerSessions.Broadcast(countdownEvent2)
	currentState.gsc.next(currentState.gsc.countdown2State)
}

func (currentState *gameSceneControllerCountdown3) disconnect(gss *GameServerSession) {
	currentState.gsc.readyState.disconnect(gss)
}
