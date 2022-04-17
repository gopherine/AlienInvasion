package world

import (
	"fmt"
)

func (c *City) EnterCity(alienName string) {
	c.Invaders = append(c.Invaders, alienName)
	if len(c.Invaders) >= 2 && !c.Destroyed {
		c.DestroyCity()
		fmt.Printf("%s and %s destroyed city %s\n", c.Invaders[0], c.Invaders[1], c.Name)
	}
}

func (c *City) LeaveCity() {
	c.Invaders = c.Invaders[1:]
}

func (c *City) DestroyCity() {
	// Disconnecting roads to stop any other alien entering this city
	if c.Roads["north"] != nil {
		delete(c.Roads["north"].Roads, "south")
	}
	if c.Roads["south"] != nil {
		delete(c.Roads["south"].Roads, "north")
	}
	if c.Roads["east"] != nil {
		delete(c.Roads["east"].Roads, "west")
	}
	if c.Roads["west"] != nil {
		delete(c.Roads["west"].Roads, "east")
	}

	// Destroying city
	delete(c.Roads, "north")
	delete(c.Roads, "south")
	delete(c.Roads, "east")
	delete(c.Roads, "west")

	c.Destroyed = true
}
