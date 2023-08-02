package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type GameServer struct {
	Upgrader         *websocket.Upgrader
	GameRoomsManager *GameRoomsManager
	config           config
}

func NewGameServer(u *websocket.Upgrader, g *GameRoomsManager, c config) *GameServer {
	return &GameServer{u, g, c}
}

func (g *GameServer) Handle(writer http.ResponseWriter, request *http.Request) {
	conn, err := g.Upgrader.Upgrade(writer, request, nil)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	user := request.URL.Query().Get(g.config.UsernameKey) // always uniq per player

	gss := NewGameServerSession(conn, user, g.config.SocTimeout)

	g.GameRoomsManager.Connect(gss)
	defer g.GameRoomsManager.Disconnect(gss)

	gss.ListenConn()
}
