package alien

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"golang.org/x/exp/maps"

	w "github.com/gopherine/alien/internal/world"
	"github.com/gopherine/alien/util"
)

// for synchronizing individual city operations
type Alien struct {
	Name string
}

var mutex sync.Mutex

func New(alienName string) *Alien {
	return &Alien{alienName}
}

// Invade function invades cities with aliens
func (a *Alien) Invade(worldMap []*w.City, steps int, wg *sync.WaitGroup) {
	defer wg.Done()
	if steps <= 0 {
		log.Error().Msg("Number of steps should be a positive integer greater then 0")
		return
	}

	// random city where this aliens spaceship will land
	city := worldMap[util.RandomInt(len(worldMap))]
	for i := 0; i < steps; i++ {
		mutex.Lock()
		city.EnterCity(a.Name)
		mutex.Unlock()

		// notify if the alien is dead or trapped
		if len(city.Roads) == 0 {
			if city.Destroyed {
				fmt.Printf("%s is Dead \n", a.Name)
				return
			}

			fmt.Printf("%s is trapped \n", a.Name)
			return
		}

		mutex.Lock()
		city.LeaveCity()
		mutex.Unlock()

		mutex.Lock()
		if len(city.Roads) != 0 {
			fmt.Printf("%s moved to city %s \n", a.Name, city.Name)
			city = city.Roads[MapRandomKeyGet(city.Roads)]
		}
		mutex.Unlock()
	}

	return
}

// Helper func to get random map keys
func MapRandomKeyGet(mapI map[string]*w.City) string {
	if len(mapI) == 0 {
		return ""
	}
	rand.Seed(time.Now().UnixNano())
	keys := maps.Keys(mapI)
	return keys[rand.Intn(len(keys))]
}
