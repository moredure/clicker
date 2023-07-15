package main

type gameSceneControllerStateFinished struct {
	dummyUserHandler
	dummyTimeout
	dummyCountdown
	dummyDisconnect
	gsc *GameSceneController
}

func (currentState *gameSceneControllerStateFinished) onStateEnter() {
	currentState.gsc.onFinishedCallback()
}
