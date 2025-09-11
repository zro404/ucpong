package game

import (
	"fmt"

	"golang.org/x/net/websocket"
)

type Ball struct {
	x, y   int
	vx, vy int
}

func NewBall() *Ball {
	return &Ball{0, 0, 0, 0}
}

type Player struct {
	conn *websocket.Conn
	pos  int
}

func NewPlayer(ws *websocket.Conn) *Player {
	return &Player{ws, 0}
}

func (p *Player) readLoop() {
	for {
		var msg PlayerInput
		err := websocket.JSON.Receive(p.conn, &msg)
		if err != nil {
			p.conn.Close()
			return
		}
		fmt.Println(msg)
		// TODO Handle input
	}
}
