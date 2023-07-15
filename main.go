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
