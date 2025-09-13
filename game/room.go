package game

import (
	"time"

	"golang.org/x/net/websocket"
)

type RoomDirectory struct {
	Rooms map[string]*Room
}

func NewRoomDirectory() *RoomDirectory {
	return &RoomDirectory{
		Rooms: make(map[string]*Room),
	}
}

type Room struct {
	p1   *Player
	p2   *Player
	ball *Ball

	isOpen bool // Used for Random Matchmaking
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

	for {
		if room.IsFull() {
			ball.x += ball.vx
			ball.y += ball.vy

			if ball.y > HEIGHT-RADIUS || ball.y < RADIUS {
				ball.vy = -ball.vy
			}

			if ball.x > WIDTH-RADIUS || ball.x < RADIUS {
				ball.vx = -ball.vx
			}
			room.broadcastState()
			time.Sleep(time.Second / 30)
		} else {
			// TODO Send "Player left the game", Reset Game
			break
		}
	}
}

func (room *Room) IsFull() bool {
	return room.p1 != nil && room.p2 != nil
}

func (room *Room) RemovePlayer(player *Player) bool {
	if room.p1 == player {
		room.p1 = nil
		return true
	} else if room.p2 == player {
		room.p2 = nil
		return true
	}

	return false
}

func (room *Room) AddPlayer(ws *websocket.Conn) (*Player, bool) {
	player := NewPlayer(ws)
	if room.p1 == nil {
		room.p1 = player
		return player, true
	} else if room.p2 == nil {
		room.p2 = player
		return player, true
	}

	return nil, false
}

func (rd *RoomDirectory) NewRoom() string {
	code := randRoomCode()
	for {
		_, ok := (*rd).Rooms[code]
		if !ok {
			break
		}
		code = randRoomCode()
	}

	(*rd).Rooms[code] = &Room{
		p1:     nil,
		p2:     nil,
		ball:   NewBall(),
		isOpen: true,
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

	room, ok := (*rd).Rooms[req.RoomCode]

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

	if room.IsFull() {
		go room.StartGame()
	}

	player.readLoop()

	// Remove Player from room
	room.RemovePlayer(player)

}

func (rd *RoomDirectory) FindGame() string {
	for k, room := range (*rd).Rooms {
		if room.isOpen && !room.IsFull() {
			return k
		}
	}

	return rd.NewRoom()
}
