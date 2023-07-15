package main

import (
	"bufio"
	"context"
	"github.com/gorilla/websocket"
	"os"
)

func main() {
	s := bufio.NewReader(os.Stdin)
	line, _, err := s.ReadLine()
	var d websocket.Dialer
	conn, _, err := d.DialContext(context.Background(), "http://localhost:80/game?username="+string(line), nil)
	if err != nil {
		panic(err)
	}
	// states as well
	for {
		conn.ReadMessage()
		line, _, err := s.ReadLine()
		err := conn.WriteMessage(websocket.TextMessage, []byte("tap"))
		if err != nil {
			panic(err)
		}
	}
}
