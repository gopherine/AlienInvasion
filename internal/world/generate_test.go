package world_test

import (
	"math"
	"testing"

	"github.com/gopherine/alien/internal/world"
)

func TestGenerate(t *testing.T) {
	testArgs := []struct {
		name      string
		numOfCity int
	}{
		{
			name:      "Testing for negative integer",
			numOfCity: -5,
		},
		{
			name:      "Testing for 0",
			numOfCity: 0,
		},
		{
			name:      "Testing for positive integer 20",
			numOfCity: 20,
		},
		{
			name:      "Testing for positive integer 100",
			numOfCity: 100,
		},
	}

	for _, v := range testArgs {
		t.Run(v.name, func(t *testing.T) {
			cities := world.Generate(v.numOfCity)

			if math.Signbit(float64(v.numOfCity)) {
				if len(cities) != 0 {
					t.Errorf("Generate(%d) FAILED: expected %d got %d", v.numOfCity, 0, len(cities))
				}
			} else if len(cities) != v.numOfCity {
				t.Errorf("Generate(%d) FAILED: expected %d got %d", v.numOfCity, v.numOfCity, len(cities))
			} else {
				t.Logf("Generate(%d) PASSED", v.numOfCity)
			}
		})
	}

}
