package main

type gameSceneControllerStatePlaying struct {
	dummyCountdown
	gsc *GameSceneController
}

func (currentState *gameSceneControllerStatePlaying) onStateEnter() {
	currentState.gsc.gameScene.deadlineTimer.Start()
}

func (currentState *gameSceneControllerStatePlaying) onStateLeave() {
	currentState.gsc.gameScene.deadlineTimer.Pause()
}

func (currentState *gameSceneControllerStatePlaying) disconnect(gss *GameServerSession) {
	currentState.gsc.readyState.disconnect(gss)
}

func (currentState *gameSceneControllerStatePlaying) tap(session *GameServerSession) {
	currentState.gsc.gameScene.Tap(session.Username)
	currentState.gsc.gameServerSessions.Publish(currentState.gsc.gameScene)
}

func (currentState *gameSceneControllerStatePlaying) timeout() {
	currentState.gsc.gameScene.ChooseWinner()
	currentState.gsc.gameServerSessions.Publish(currentState.gsc.gameScene)
	currentState.gsc.gameServerSessions.Close()
	currentState.gsc.next(currentState.gsc.finishedState)
}
