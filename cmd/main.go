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

	// number of Cities default is 10

	flag.IntVar(&numOfCities, "n", 10, "enter number of cities to be generated")

	// number of aliens default is 2
	flag.IntVar(&numOfAliens, "a", 2, "enter number of aliens to be dropped from spaceships")

	// number of steps an alien can move default is 10000
	flag.IntVar(&numOfSteps, "s", 10000, "enter number of steps each alien will travel")
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
		if numOfCities < 0 || len(worldMap) == 0 {
			log.Error().Msg("Please generate or provide a map")
			break
		}
		fmt.Println(worldMap)
		fmt.Println(numOfSteps)
		var wg sync.WaitGroup
		for i := 0; i < numOfAliens; i++ {
			wg.Add(1)
			a := alien.New(fmt.Sprintf("alien-%d", i))
			go a.Invade(worldMap, numOfSteps, &wg)
		}
		wg.Wait()

		fmt.Printf("Total number of city destroyed %d \n", world.TotalCityDestroyed(worldMap))
	}
}
