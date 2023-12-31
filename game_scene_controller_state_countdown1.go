package main

import (
	"time"
)

type gameSceneControllerCountdown1 struct {
	dummyUserHandler
	dummyTimeout
	gsc *GameSceneController
}

func (currentState *gameSceneControllerCountdown1) onStateEnter() {
	currentState.gsc.reusableTimer = time.AfterFunc(time.Second, currentState.gsc.countdown)
}

func (currentState *gameSceneControllerCountdown1) onStateLeave() {
	currentState.gsc.reusableTimer.Stop()
}

func (currentState *gameSceneControllerCountdown1) countdown() {
	currentState.gsc.gameServerSessions.Broadcast(startEvent)
	currentState.gsc.next(currentState.gsc.playingState)
}

func (currentState *gameSceneControllerCountdown1) disconnect(gss *GameServerSession) {
	currentState.gsc.readyState.disconnect(gss)
}
