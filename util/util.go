package util

import (
	"math/rand"
	"time"
)

// Helper func to get random integers
func RandomInt(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}
