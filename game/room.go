package game

import (
	"golang.org/x/net/websocket"
)

type RoomDirectory map[string]*Room

func NewRoomDirectory() *RoomDirectory {
	return &RoomDirectory{}
}

type Room struct {
	p1   *Player
	p2   *Player
	ball *Ball
}

func (room *Room) broadcastState() {
	state := GameState{
		BallX:   room.ball.x,
		BallY:   room.ball.y,
		Player1: room.p1.pos,
		Player2: room.p2.pos,
	}
	if room.p1 != nil {
		websocket.JSON.Send(room.p1.conn, state)
	}
	if room.p2 != nil {
		websocket.JSON.Send(room.p2.conn, state)
	}
}

func (room *Room) StartGame() {
	ball := room.ball
	ball.vx = 10
	ball.vy = 10

	for {
		ball.x += ball.vx
		ball.y += ball.vy

		if ball.y > 100 || ball.y < -100 {
			ball.vy = -ball.vy
		}
		room.broadcastState()
	}
}

func (room *Room) IsFull() bool {
	return room.p1 != nil && room.p2 != nil
}

func (room *Room) AddPlayer(ws *websocket.Conn) (*Player, bool) {
	player := NewPlayer(ws)
	if room.p1 == nil {
		room.p1 = player
		return player, true
	} else if room.p2 == nil {
		room.p2 = player
		// Start the game
		go room.StartGame()

		return player, true
	}

	return nil, false
}

func (rd *RoomDirectory) NewRoom() string {
	code := randRoomCode()
	for {
		_, ok := (*rd)[code]
		if !ok {
			break
		}
		code = randRoomCode()
	}

	(*rd)[code] = &Room{
		p1:   nil,
		p2:   nil,
		ball: NewBall(),
	}

	return code

}

func (rd *RoomDirectory) HandleNewPlayer(ws *websocket.Conn) {
	req := JoinRequest{}
	err := websocket.JSON.Receive(ws, &req)

	if err != nil {
		websocket.JSON.Send(ws, ErrorMsg{false, "Room Code Required!"})
		ws.Close()
		return
	}

	room, ok := (*rd)[req.RoomCode]

	if !ok {
		websocket.JSON.Send(ws, ErrorMsg{false, "Invalid Room Code!"})
		ws.Close()
		return
	}

	player, ok := room.AddPlayer(ws)
	if !ok {
		websocket.JSON.Send(ws, ErrorMsg{false, "Room Full!"})
		ws.Close()
		return
	}

	player.readLoop()
}
