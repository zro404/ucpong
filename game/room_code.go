package game

import (
	"math/rand"
)

func randRoomCode() string {
	b := make([]byte, 6)
	for i := range b {
		b[i] = byte(rand.Intn(10) + 48)
	}

	return string(b)
}
