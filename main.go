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
	// объявить структур/объект
	//разделить на два

	//вставить файл конфигурации. порт подключения. задать имя пользвателя.
	//user := request.URL.Query().Get("username") // always uniq per player
	//port
	//rooms: make(map[string]*GameSceneController, 64),
	//deadline := time.Now().Add(15 * time.Second)
	//timer := NewPausableAfterFunc(30*time.Second, gsc.connectionTimeout) socet timuout game connectionTimeout
	gs := NewGameServer(new(websocket.Upgrader), grm, c)

	r.Handle("/game", http.HandlerFunc(gs.Handle))
	s := http.Server{
		Addr:    c.HTTPAddr,
		Handler: r,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
