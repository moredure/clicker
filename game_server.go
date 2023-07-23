package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type GameServer struct {
	Upgrader         *websocket.Upgrader
	GameRoomsManager *GameRoomsManager
}

func (g *GameServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	conn, err := g.Upgrader.Upgrade(writer, request, nil)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	user := request.URL.Query().Get("username") // always uniq per player

	gss := NewGameServerSession(conn, user)

	g.GameRoomsManager.Connect(gss)
	defer g.GameRoomsManager.Disconnect(gss)

	gss.ListenConn()
}
