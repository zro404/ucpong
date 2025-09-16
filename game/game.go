package game

import (
	"sync"

	"golang.org/x/net/websocket"
)

const WIDTH int = 1000
const HEIGHT int = 600
const RADIUS int = 15
const VX = 10
const VY = 10

const PADDLE_HEIGHT = 150
const PADDLE_WIDTH = 25
const PADDLE_VELOCITY = 25

type Ball struct {
	x, y   int
	vx, vy int
}

func NewBall() *Ball {
	return &Ball{WIDTH / 2, HEIGHT / 2, VX, VY}
}

func (b *Ball) reset() {
	b.x = WIDTH / 2
	b.y = HEIGHT / 2
	b.vx = -b.vx
}

type Player struct {
	conn  *websocket.Conn
	pos   int
	score int
	ready bool
	lock  sync.Mutex
}

func NewPlayer(ws *websocket.Conn) *Player {
	return &Player{ws, HEIGHT / 2, 0, false, sync.Mutex{}}
}

func (p *Player) reset() {
	p.pos = HEIGHT / 2
	p.ready = false
}

func (p *Player) readLoop() {
	for {
		var msg PlayerInput
		err := websocket.JSON.Receive(p.conn, &msg)
		if err != nil {
			p.lock.Lock()
			p.conn.Close()
			p.lock.Unlock()
			return
		}

		switch msg.Action {
		case 1:
			if p.pos != HEIGHT-PADDLE_HEIGHT/2 {
				p.pos += PADDLE_VELOCITY
			}
		case 2:
			if p.pos != PADDLE_HEIGHT/2 {
				p.pos -= PADDLE_VELOCITY
			}

		case 3:
			p.ready = true
		}
	}
}
