package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type GameServer struct {
	Upgrader         *websocket.Upgrader
	GameRoomsManager *GameRoomsManager
}

func (g *GameServer) Handle(writer http.ResponseWriter, request *http.Request) {
	conn, err := g.Upgrader.Upgrade(writer, request, nil)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	user := request.Header.Get("User") // always uniq per player

	ps := &GameServerSession{
		Username: user,
		conn:     conn,
	}

	g.GameRoomsManager.Connect(ps)
	defer g.GameRoomsManager.Disconnect(ps)
	ps.ListenConn()
}
