package main

import (
	"bytes"
	"log"
	"reflect"
	"sync"
	"time"
)

type GameSceneController struct {
	sync.Mutex
	readyState      GameSceneControllerState
	countdown2State GameSceneControllerState
	countdown3State GameSceneControllerState
	countdown1State GameSceneControllerState
	playingState    GameSceneControllerState
	finishedState   GameSceneControllerState

	gameScene *GameScene

	currentState       GameSceneControllerState
	reusableTimer      *time.Timer
	gameServerSessions GameServerSessions
	onFinishedCallback func()
}

func (gsc *GameSceneController) communicate(s *GameServerSession, command []byte) {
	gsc.Lock()
	defer gsc.Unlock()

	switch {
	case bytes.Contains(command, tapCommand):
		gsc.currentState.tap(s)
	default:
	}
}

func (gsc *GameSceneController) disconnect(session *GameServerSession) {
	gsc.Lock()
	defer gsc.Unlock()
	gsc.currentState.disconnect(session)
}

func (gsc *GameSceneController) countdown() {
	gsc.Lock()
	defer gsc.Unlock()
	gsc.currentState.countdown()
}

func (gsc *GameSceneController) timeout() {
	gsc.Lock()
	defer gsc.Unlock()
	gsc.currentState.timeout()
}

func (gsc *GameSceneController) next(gscs GameSceneControllerState) {
	log.Println("next", reflect.TypeOf(gscs))
	leaver, ok := gsc.currentState.(stateLeaver)
	if ok {
		leaver.onStateLeave()
	}
	enter, ok := gscs.(stateEnterer)
	if ok {
		enter.onStateEnter()
	}
	gsc.currentState = gscs
}

func NewGameSceneController(lobby map[*GameServerSession]struct{}, callback func([]string)) *GameSceneController {
	usernames := make([]string, 0, 2)
	playerNameToSession := make(map[string]*GameServerSession, 2)
	for gss := range lobby {
		playerNameToSession[gss.Username] = gss
		usernames = append(usernames, gss.Username)
	}
	gsc := &GameSceneController{
		gameServerSessions: playerNameToSession,
		onFinishedCallback: func() {
			callback(usernames)
		},
	}
	gs := new(GameScene)
	gs.stats = make(map[string]int, 2)
	gs.deadlineTimer = NewPausableAfterFunc(30*time.Second, gsc.timeout)
	gsc.gameScene = gs
	gsc.readyState = &gameSceneControllerStateReady{gsc: gsc}
	gsc.countdown3State = &gameSceneControllerCountdown3{gsc: gsc}
	gsc.countdown2State = &gameSceneControllerCountdown2{gsc: gsc}
	gsc.countdown1State = &gameSceneControllerCountdown1{gsc: gsc}
	gsc.playingState = &gameSceneControllerStatePlaying{gsc: gsc}
	gsc.finishedState = &gameSceneControllerStateFinished{gsc: gsc}
	gsc.next(gsc.readyState)
	return gsc
}
