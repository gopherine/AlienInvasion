package world

import (
	"bufio"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
)

type City struct {
	//name of the city
	Name string `csv:"name"`
	//This map contains directions like north south east west as keys and the corresponding city as value
	Roads map[string]*City `csv:"roads"`
	//number of aliens in the city
	Invaders []string `csv:"invaders"`
	//Keeps track of city status if destroyed or not
	Destroyed bool `csv:"destroyed"`
}

func LoadCities() []*City {
	f, err := os.Open(os.Getenv("FILENAME"))
	if err != nil {
		log.Fatal().Msgf("Unable to open file %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	citiesMap := make(map[string]*City)
	var cities []*City
	for scanner.Scan() {
		// line := strings.Trim(scanner.Text(), " \t\n\r")
		line := strings.Fields(scanner.Text())

		city := new(City)
		city.Roads = make(map[string]*City)
		city.Name = line[0]

		citiesMap[line[0]] = city
		var nextCity *City

		for i := 1; i < len(line); i++ {
			tuple := strings.Split(line[i], "=")

			if _, ok := citiesMap[tuple[1]]; !ok {
				nextCity = new(City)
				nextCity.Roads = make(map[string]*City)
				nextCity.Name = tuple[1]
				citiesMap[tuple[1]] = nextCity
			} else {
				nextCity = citiesMap[tuple[1]]
			}

			switch tuple[0] {
			case "north":
				city.Roads["north"] = nextCity
				nextCity.Roads["south"] = city
			case "south":
				city.Roads["south"] = nextCity
				nextCity.Roads["north"] = city
			case "east":
				city.Roads["east"] = nextCity
				nextCity.Roads["west"] = city
			case "west":
				city.Roads["west"] = nextCity
				nextCity.Roads["east"] = city
			}
		}
		cities = append(cities, city)
	}

	return cities
}
