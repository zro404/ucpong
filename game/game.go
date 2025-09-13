package game

import (
	"fmt"

	"golang.org/x/net/websocket"
)

const WIDTH int = 1000
const HEIGHT int = 600
const RADIUS int = 15
const VX = 15
const VY = 10

type Ball struct {
	x, y   int
	vx, vy int
}

func NewBall() *Ball {
	return &Ball{WIDTH / 2, HEIGHT / 2, VX, VY}
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
