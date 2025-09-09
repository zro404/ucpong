package game

type ErrorMsg struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type JoinRequest struct {
	RoomCode string `json:"roomCode"`
}

type PlayerInput struct {
	Action int `json:"action"`
}
