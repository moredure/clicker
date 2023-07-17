package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

type GameServerSession struct {
	sync.Mutex
	f func(*GameServerSession, []byte)

	conn     *websocket.Conn
	Username string
}

func NewGameServerSession(conn *websocket.Conn, username string) *GameServerSession {
	return &GameServerSession{conn: conn, Username: username}
}

func (gss *GameServerSession) Subscribe(f func(*GameServerSession, []byte)) {
	gss.Lock()
	defer gss.Unlock()
	gss.f = f
}

func (gss *GameServerSession) Notify(data []byte) {
	gss.Lock()
	f := gss.f
	gss.Unlock()
	if f == nil {
		return
	}
	f(gss, data)
}

func (gss *GameServerSession) Close() {
	gss.conn.Close()
}

func (gss *GameServerSession) ListenConn() {
	for {
		deadline := time.Now().Add(15 * time.Second)
		if err := gss.conn.SetReadDeadline(deadline); err != nil {
			return
		}
		_, data, err := gss.conn.ReadMessage()
		if err != nil {
			return
		}
		gss.Notify(data)
	}
}

func (gss *GameServerSession) emit(e event) {
	log.Println("emitted", gss.Username, e.Type())
	m := map[string]any{
		"type":  e.Type(),
		"value": e,
	}
	data, _ := json.Marshal(m)
	if err := gss.conn.SetWriteDeadline(time.Now().Add(200 * time.Millisecond)); err != nil {
		log.Println("failed to set write deadline", err)
		return
	}
	if err := gss.conn.WriteMessage(websocket.TextMessage, data); err != nil {
		log.Println("conn failed to WriteMessage", err)
		return
	}
}
