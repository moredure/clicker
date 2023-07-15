package main

import (
	"log"
	"sync"
)

type GameRoomsManager struct {
	sync.Mutex
	lobby map[*GameServerSession]struct{}
	rooms map[string]*GameSceneController
}

func NewRoomsManager() *GameRoomsManager {
	return &GameRoomsManager{
		lobby: make(map[*GameServerSession]struct{}, 2),
		rooms: make(map[string]*GameSceneController, 64),
	}
}

func (grm *GameRoomsManager) cleanup(sessionIDs []string) {
	grm.Lock()
	defer grm.Unlock()
	for _, id := range sessionIDs {
		delete(grm.rooms, id)
	}
	log.Println("finished game")
}

func (grm *GameRoomsManager) Connect(session *GameServerSession) {
	grm.Lock()
	defer grm.Unlock()

	grm.lobby[session] = struct{}{}
	if len(grm.lobby) != 2 {
		return
	}
	lobby := grm.lobby
	grm.lobby = make(map[*GameServerSession]struct{}, 2)
	gsc := NewGameSceneController(lobby, grm.cleanup)
	for ps := range lobby {
		ps.Subscribe(gsc.communicate)
	}
	go gsc.countdown()
	grm.rooms[session.Username] = gsc
	return
}

func (grm *GameRoomsManager) Disconnect(session *GameServerSession) {
	grm.Lock()
	defer grm.Unlock()

	if gsc := grm.rooms[session.Username]; gsc != nil {
		gsc.disconnect(session)
		return
	}
	delete(grm.lobby, session)
}
