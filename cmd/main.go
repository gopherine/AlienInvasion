package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"

	"github.com/gopherine/alien/internal/alien"
	"github.com/gopherine/alien/internal/world"
)

var (
	cmd                                  string
	numOfCities, numOfAliens, numOfSteps int
)

func init() {
	// Choose the action to be executed
	flag.StringVar(&cmd, "cmd", "start",
		`* generate: Generates new world
* start: Invade the world`)

	// number of Cities

	flag.IntVar(&numOfCities, "n", numOfCities, "enter number of cities to be generated")

	// number of aliens
	flag.IntVar(&numOfAliens, "a", numOfAliens, "enter number of aliens to be dropped from spaceships")

	// number of steps
	flag.IntVar(&numOfSteps, "s", numOfSteps, "enter number of steps each alien will travel")
	flag.Parse()
}
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msg("Error loading .env file")
	}

	switch cmd {
	case "generate":
		cities := world.Generate(numOfCities)
		world.WriteToFile(cities)
	case "start":
		worldMap := world.LoadCities()
		if numOfCities < 0 {
			break
		}

		var wg sync.WaitGroup
		for i := 0; i < numOfAliens; i++ {
			wg.Add(1)
			a := alien.New(fmt.Sprintf("alien-%d", i))
			go a.Invade(worldMap, numOfSteps, &wg)
		}
		wg.Wait()

		for _, v := range worldMap {
			fmt.Println(v)
		}
	}
}
