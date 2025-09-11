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

type GameState struct {
	BallX   int `json:"ballX"`
	BallY   int `json:"ballY"`
	Player1 int `json:"player1"`
	Player2 int `json:"player2"`
}
