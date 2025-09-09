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

func (room *Room) AddPlayer(ws *websocket.Conn) bool {
	player := NewPlayer(ws)
	if room.p1 == nil {
		room.p1 = player
		player.readLoop()
		return true
	} else if room.p2 == nil {
		room.p2 = player
		player.readLoop()
		return true
	}

	return false
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

	ok = room.AddPlayer(ws)
	if !ok {
		websocket.JSON.Send(ws, ErrorMsg{false, "Room Full!"})
		ws.Close()
		return
	}

}
