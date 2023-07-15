package main

import (
	"log"
	"time"
)

type gameSceneControllerCountdown2 struct {
	dummyUserHandler
	gsc *GameSceneController
}

func (currentState *gameSceneControllerCountdown2) onStateEnter() {
	currentState.gsc.reusableTimer = time.AfterFunc(time.Second, currentState.gsc.countdown)
}

func (currentState *gameSceneControllerCountdown2) onStateLeave() {
	currentState.gsc.reusableTimer.Stop()
}

func (currentState *gameSceneControllerCountdown2) countdown() {
	currentState.gsc.gameServerSessions.Broadcast(countdownEvent1)
	currentState.gsc.next(currentState.gsc.countdown1State)
}

func (currentState *gameSceneControllerCountdown2) timeout() {
	log.Println("it's possible but do nothing")
}

func (currentState *gameSceneControllerCountdown2) disconnect(gss *GameServerSession) {
	currentState.gsc.readyState.disconnect(gss)
}
