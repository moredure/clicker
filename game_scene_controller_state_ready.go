package main

import (
	"log"
	"time"
)

type gameSceneControllerStateReady struct {
	dummyUserHandler
	gsc *GameSceneController
}

func (currentState *gameSceneControllerStateReady) onStateEnter() {
	currentState.gsc.reusableTimer = time.AfterFunc(time.Second, currentState.gsc.countdown)
}

func (currentState *gameSceneControllerStateReady) onStateLeave() {
	currentState.gsc.reusableTimer.Stop()
}

func (currentState *gameSceneControllerStateReady) countdown() {
	currentState.gsc.gameServerSessions.Broadcast(countdownEvent3)
	currentState.gsc.next(currentState.gsc.countdown3State)
}

func (currentState *gameSceneControllerStateReady) timeout() {
	log.Println("it's possible but do nothing")
}

func (currentState *gameSceneControllerStateReady) disconnect(session *GameServerSession) {
	delete(currentState.gsc.gameServerSessions, session.Username)
	currentState.gsc.gameScene.Loose(session.Username)
	currentState.gsc.gameServerSessions.Publish(currentState.gsc.gameScene)
	currentState.gsc.gameServerSessions.Close()
	currentState.gsc.next(currentState.gsc.finishedState)
}
