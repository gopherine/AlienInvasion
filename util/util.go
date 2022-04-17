package util

import (
	"math/rand"
	"time"

	"golang.org/x/exp/maps"

	"github.com/gopherine/alien/internal/world"
)

// Helper func to get random integers
func RandomInt(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

// Helper func to get random map keys
func MapRandomKeyGet(mapI map[string]*world.City) string {
	rand.Seed(time.Now().UnixNano())
	keys := maps.Keys(mapI)
	return keys[rand.Intn(len(keys))]
}
