package main

import (
	"log"
)

type gameSceneControllerStatePlaying struct {
	gsc *GameSceneController
}

func (currentState *gameSceneControllerStatePlaying) onStateEnter() {
	currentState.gsc.gameScene.deadlineTimer.Start()
}

func (currentState *gameSceneControllerStatePlaying) onStateLeave() {
	currentState.gsc.gameScene.deadlineTimer.Pause()
}

func (currentState *gameSceneControllerStatePlaying) countdown() {
	log.Println("it's possible but do nothing")
}

func (currentState *gameSceneControllerStatePlaying) disconnect(gss *GameServerSession) {
	currentState.gsc.readyState.disconnect(gss)
}

func (currentState *gameSceneControllerStatePlaying) userTap(session *GameServerSession) {
	currentState.gsc.gameScene.Tap(session.Username)
	currentState.gsc.gameServerSessions.Publish(currentState.gsc.gameScene)
}

func (currentState *gameSceneControllerStatePlaying) timeout() {
	currentState.gsc.gameScene.Finish()
	currentState.gsc.gameServerSessions.Publish(currentState.gsc.gameScene)
	currentState.gsc.gameServerSessions.Close()
	currentState.gsc.next(currentState.gsc.finishedState)
}
