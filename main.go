package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	grm := NewRoomsManager()
	// extend with some configuration management
	// provide some persistance layer for players/rooms/games
	gs := &GameServer{
		Upgrader:         new(websocket.Upgrader),
		GameRoomsManager: grm,
	}
	r.Handle("/game", http.HandlerFunc(gs.Handle))
	s := http.Server{
		Addr:    ":80",
		Handler: r,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
