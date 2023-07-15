package main

import "log"

type gameSceneControllerStateFinished struct {
	dummyUserHandler
	gsc *GameSceneController
}

func (currentState *gameSceneControllerStateFinished) onStateEnter() {
	currentState.gsc.onFinishedCallback()
}

func (currentState *gameSceneControllerStateFinished) countdown() {
	log.Println("do nothing")
}

func (currentState *gameSceneControllerStateFinished) disconnect(session *GameServerSession) {
	log.Println("it's not possible")
	return
}

func (currentState *gameSceneControllerStateFinished) timeout() {
	log.Println("it's possible but do nothing")
}
