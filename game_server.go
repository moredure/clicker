package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type Upgrader interface {
	Upgrade(http.ResponseWriter, *http.Request, http.Header) (*websocket.Conn, error)
}

type RoomsManager interface {
	Connect(*GameServerSession)
	Disconnect(*GameServerSession)
}

type GameServer struct {
	Upgrader         Upgrader
	GameRoomsManager RoomsManager
	config           config
}

func NewGameServer(u Upgrader, g RoomsManager, c config) *GameServer {
	return &GameServer{u, g, c}
}

func (g *GameServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
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
