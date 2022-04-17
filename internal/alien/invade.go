package world

import (
	"fmt"
	"sync"

	"github.com/rs/zerolog/log"

	"github.com/gopherine/alien/internal/world"
	"github.com/gopherine/alien/util"
)

// for synchronizing individual city operations
var mutex sync.Mutex

// Invade function invades cities with aliens
func Invade(worldMap []*world.City, alienName string, steps int, wg *sync.WaitGroup) {
	defer wg.Done()
	if steps <= 0 {
		log.Error().Msg("Number of steps should be a positive integer greater then 0")
		return
	}

	// random city where this aliens spaceship will land
	city := worldMap[util.RandomInt(len(worldMap))]
	for i := 0; i < steps; i++ {
		mutex.Lock()
		city.EnterCity(alienName)
		mutex.Unlock()

		// notify if the alien is dead or trapped
		if len(city.Roads) == 0 {
			if city.Destroyed {
				fmt.Printf("%s is Dead \n", alienName)
				return
			}

			fmt.Printf("%s is trapped \n", alienName)
			return
		}

		mutex.Lock()
		city.LeaveCity()
		mutex.Unlock()

		mutex.Lock()
		if len(city.Roads) != 0 {
			fmt.Printf("%s moved to city %s \n", alienName, city.Name)
			city = city.Roads[util.MapRandomKeyGet(city.Roads)]
		}
		mutex.Unlock()
	}

	return
}
