package world_test

import (
	"testing"

	"github.com/gopherine/alien/internal/world"
)

func TestEnterCity(t *testing.T) {
	var city1, city2 world.City

	t.Run("City should be destroyed", func(t *testing.T) {
		city1 = world.City{
			Name:     "test",
			Invaders: []string{"test"},
			Roads: map[string]*world.City{
				"north": &city2, // using same city just for simplification
				"south": &city2,
				"east":  &city2,
				"west":  &city2,
			},
			Destroyed: false,
		}

		city2 = world.City{
			Name:     "test",
			Invaders: []string{"test"},
			Roads: map[string]*world.City{
				"north": &city1, // using same city just for simplification
				"south": &city1,
				"east":  &city1,
				"west":  &city1,
			},
			Destroyed: false,
		}

		city1.EnterCity("test2")
		if city1.Destroyed != true {
			t.Errorf("EnterCity(test2) FAILED: expected city to be destroyed")
		}

		if len(city1.Roads) != 0 && len(city2.Roads) == 0 {
			t.Errorf("EnterCity(test2) FAILED: city still has connected roads")
		} else {
			t.Logf("EnterCity(test2) PASSED")
		}
	})
}

func TestLeaveCity(t *testing.T) {
	city := world.City{
		Name:      "test",
		Invaders:  []string{"test"},
		Destroyed: false,
	}

	city.LeaveCity()
	if len(city.Invaders) != 0 {
		t.Errorf("LeaveCity() FAILED: invader unable to leave city")
	} else {
		t.Logf("LeaveCity() PASSED")
	}
}
