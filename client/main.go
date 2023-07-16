package main

import (
	"bufio"
	"context"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"sync"
)

func main() {
	s := bufio.NewReader(os.Stdin)
	line, _, err := s.ReadLine()
	var d websocket.Dialer
	conn, _, err := d.DialContext(context.Background(), "ws://localhost:80/game?username="+string(line), nil)
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			t, data, err := conn.ReadMessage()
			if err != nil {
				log.Println("close", string(data))
				return
			}
			log.Println(t, string(data))
		}
	}()
	go func() {
		defer wg.Done()
		for {
			s.ReadLine()
			conn.WriteMessage(websocket.TextMessage, []byte("tap"))
		}
	}()
	wg.Wait()
}
