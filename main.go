package main

import (
	"github.com/caarlos0/env"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	var c config
	if err := env.Parse(&c); err != nil {
		panic(err)
	}
	grm := NewRoomsManager()
	// extend with some configuration management
	// provide some persistance layer for players/rooms/games
	gs := NewGameServer(new(websocket.Upgrader), grm, c)
	r.Handle("/game", gs)
	s := http.Server{
		Addr:    ":80",
		Handler: r,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
