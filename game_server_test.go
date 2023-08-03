package main

import (
	"github.com/gorilla/websocket"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type RoomsManagerMock struct {
	ConnectCalled    bool
	DisconnectCalled bool
}

func (rmm *RoomsManagerMock) Connect(*GameServerSession) {
	rmm.ConnectCalled = true
}

func (rmm *RoomsManagerMock) Disconnect(*GameServerSession) {
	rmm.DisconnectCalled = true
}

func TestGameServer_ServeHTTP(t *testing.T) {
	u := new(websocket.Upgrader)
	rmm := new(RoomsManagerMock)
	defer func() {
		if !rmm.DisconnectCalled || !rmm.ConnectCalled {
			t.Fatal("mock was not used")
		}
	}()
	c := config{
		GameDuration: 30 * time.Second,
		SocTimeout:   15 * time.Second,
		UsernameKey:  "username",
	}
	gs := NewGameServer(u, rmm, c)
	s := httptest.NewServer(gs)
	defer s.Close()

	var d websocket.Dialer
	uri := "ws" + strings.TrimPrefix(s.URL, "http")
	conn, _, err := d.Dial(uri+"/?"+c.UsernameKey+"=solo", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
}
