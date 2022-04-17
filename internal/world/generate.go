package world

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/goombaio/namegenerator"
	"github.com/rs/zerolog/log"
)

// Generate - generates a new world
func Generate(numOfCities int) []*City {
	// number of cities cannot be less then or equal 0
	if numOfCities <= 0 {
		log.Fatal().Msg("Cannot generate a world please provide a positive integer greater then 0 to generate world")
	}

	// Initialize list of city
	cities := make([]*City, numOfCities)
	for i := range cities {
		cities[i] = newCity()
	}

	connectRoads(cities)
	writeToFile(cities)
	return cities
}

// Initialize an empty city to avoid nil pointers
func newCity() *City {
	city := new(City)
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	// generate random names with name generator library this is not 100 percent unique but a temporary solution
	city.Name = nameGenerator.Generate()
	city.Roads = make(map[string]*City)

	return city
}

// Connect roads randomly for the provided list of cities
func connectRoads(c []*City) []*City {
	// Initializing directions on list as  0==North  1==South  2==East 3==West
	// if an index value is 1 - that indicates a city is connected to it by road
	rand.Seed(time.Now().UnixNano())
	directions := make([]int, 4)

	// counter to point to next index
	next := 1
	for i := range c {
		// generate random directions
		for i := range directions {
			directions[i] = rand.Intn(2)
		}

		if next < len(c) && c[i].Roads["north"] == nil && directions[0] == 1 {
			c[i].Roads["north"] = c[next]
			c[next].Roads["south"] = c[i]
			next++
		}

		if next < len(c) && c[i].Roads["south"] == nil && directions[1] == 1 {
			c[i].Roads["south"] = c[next]
			c[next].Roads["north"] = c[i]
			next++
		}

		if next < len(c) && c[i].Roads["east"] == nil && directions[2] == 1 {
			c[i].Roads["east"] = c[next]
			c[next].Roads["west"] = c[i]
			next++
		}

		if next < len(c) && c[i].Roads["west"] == nil && directions[3] == 1 {
			c[i].Roads["west"] = c[next]
			c[next].Roads["east"] = c[i]
			next++
		}
	}

	return c
}

// WriteToFile writes list of cities to text file
func writeToFile(c []*City) {
	f, err := os.Create(os.Getenv("FILENAME"))
	if err != nil {
		log.Fatal().Msgf("failed creating file: %s", err)
	}
	defer f.Close()

	for _, city := range c {
		row := city.Name
		for k, v := range city.Roads {
			row = fmt.Sprintf("%s %s=%s", row, k, v.Name)
		}
		f.WriteString(fmt.Sprintf("%s \n", row))
	}

}
