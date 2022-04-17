Alien invasion 
============== 

Mad aliens are about to invade the earth and you are tasked with simulating the invasion. 
You are given a map containing the names of cities in the non-existent world of X. The map is in a file, with one city per line. The city name is first, followed by 1-4 directions (north, south, east, or west). Each one represents a road to another city that lies in that direction. 
For example: 
```
Foo north=Bar west=Baz south=Qu-ux 
Bar south=Foo west=Bee 
```
The city and each of the pairs are separated by a single space, and the directions are separated from their respective cities with an equals (=) sign. 
You should create N aliens, where N is specified as a command-line argument. 
These aliens start out at random places on the map, and wander around randomly, following links. Each iteration, the aliens can travel in any of the directions leading out of a city. In our example above, an alien that starts at Foo can go north to Bar, west to Baz, or south to Qu-ux. 
When two aliens end up in the same place, they fight, and in the process kill each other and destroy the city. When a city is destroyed, it is removed from the map, and so are any roads that lead into or out of it. 
In our example above, if Bar were destroyed the map would now be something like: 
Foo west=Baz south=Qu-ux 
Once a city is destroyed, aliens can no longer travel to or through it. This may lead to aliens getting "trapped". 
You should create a program that reads in the world map, creates N aliens, and unleashes them. The program should run until all the aliens have been destroyed, or each alien has moved at least 10,000 times. When two aliens 

User Story
============== 

* Generate new map with randomly connected city
    * The generate map should be a practical map
* Save the generated map to a textfile
* Load any map from the textfile
    * Assuming the format of the map is in the format of above given example and is a practical map implementation
* Randomly generate aliens
* Aliens should move in and out of the city randomly make it concurrent
* Loop and move for the number of steps provided
* if two or more aleans move to same city return from goroutine and do not move any steps
* aliens can either be trapped or destroyed

Installation
============== 

* Clone this repository and go to root of the repository via terminal
* run `make build`
* run binary `bin/alien-invasion -h`
```
Usage of bin/alien-invasion:
  -a int
        enter number of aliens to be dropped from spaceships (default 2)
  -cmd string
        * generate: Generates new world
        * start: Invade the world (default "start")
  -n int
        enter number of cities to be generated (default 10)
  -s int
        enter number of steps each alien will travel (default 10000)
```
* If for some reason make does not work use below command to generate binary
```
go build -o bin/alien-invasion ./cmd/main.go
```

Usage
============== 
```
// Generate a new map
bin/alien-invasion -cmd generate

// Run Simulation
bin/alien-invasion
```

When explicit flag values are not provided default values are used.
default values are for the flag:
```
-cmd start
-a 2
-n 10
-s 10000
```

Program assumes that user can provide any practical number of steps, aliens or city within hardware limitation.
This package uses `github.com/goombaio/namegenerator` package to generate unique names this could have been manually randomized as well.
During testing this always generated n number of unique names but this is not guaranteed. An alternate solution to generate unique name
can easily be implemented.

This application also assumes that map can be provided in the above example format from any external source and need not be always
generated from the application. In case the provided map is of different name **PLEASE CHANGE FILENAME IN .env** file to load the map from
the provided file.

Test
============== 
```
# Run unit test with coverage
go test -cover ./...
```

Assumtion
============== 
* All the passed value such as number of aliens, number of city and steps respectively are always positive.
* City names cannot contain spaces