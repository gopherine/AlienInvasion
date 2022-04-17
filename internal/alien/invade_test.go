package alien_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/gopherine/alien/internal/alien"
	"github.com/gopherine/alien/internal/world"
)

// Tests invade function currently testing for hardcoded values could be improved with different scenarios
func TestInvade(t *testing.T) {
	numOfAliens := 20
	newMap := world.Generate(100)
	var wg sync.WaitGroup
	for i := 0; i < numOfAliens; i++ {
		wg.Add(1)
		a := alien.New(fmt.Sprintf("alien-%d", i))
		go a.Invade(newMap, 100000, &wg)
	}
	wg.Wait()

	numOfDestroyedCity := world.TotalCityDestroyed(newMap)
	if numOfDestroyedCity < 1 || numOfDestroyedCity > 100 {
		t.Errorf("Invade() Failed: number of city destroyed is %d which is out of range 1 to 100", numOfDestroyedCity)
	} else {
		t.Logf("Invade() PASSED")
	}

}
