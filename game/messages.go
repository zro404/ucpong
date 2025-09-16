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

type BroadcastType string

const (
	InProgress BroadcastType = "inProgress"
	GameOver   BroadcastType = "gameOver"
	Disconnect BroadcastType = "disconnect"
)

type GameState struct {
	Type         BroadcastType `json:"type"`
	BallX        int           `json:"ballX"`
	BallY        int           `json:"ballY"`
	PlayerPos1   int           `json:"playerPos1"`
	PlayerPos2   int           `json:"playerPos2"`
	PlayerScore1 int           `json:"playerScore1"`
	PlayerScore2 int           `json:"playerScore2"`
}
