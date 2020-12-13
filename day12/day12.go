package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Direction int
const (
	East Direction = iota
	South
	West
	North
)
func (d Direction) String() string {
	return [...]string{"East", "South", "West", "North"}[d]
}

type Coordinates struct {
	X, Y int
}

type Ship struct {
	Facing   Direction
	Location Coordinates
	Waypoint Coordinates
}

func (s Ship) doNavigation(instruction string) (ship Ship, err error){
	action := instruction[0]
	var value int
	value, err = strconv.Atoi(instruction[1:])
	if err != nil {
		return
	}

	ship.Facing = s.Facing
	ship.Location = s.Location
	ship.Waypoint = s.Waypoint
	switch action {
	case 'N':
		ship.Location.Y += value
	case 'S':
		ship.Location.Y -= value
	case 'E':
		ship.Location.X += value
	case 'W':
		ship.Location.X -= value
	case 'L':
		turns := value / 90
		temp := int(s.Facing) - turns
		if temp < 0 {
			ship.Facing = Direction(4 + temp)
		} else {
			ship.Facing = Direction(temp)
		}
	case 'R':
		turns := value / 90
		temp := int(s.Facing) + turns
		if temp > 3 {
			ship.Facing = Direction(temp - 4)
		} else {
			ship.Facing = Direction(temp)
		}
	case 'F':
		switch s.Facing {
		case East:
			ship.Location.X += value
		case South:
			ship.Location.Y -= value
		case West:
			ship.Location.X -= value
		case North:
			ship.Location.Y += value
		}
	}
	return
}

func (s Ship) doNavigationByWaypoint(instruction string) (ship Ship, err error) {
	action := instruction[0]
	var value int
	value, err = strconv.Atoi(instruction[1:])
	if err != nil {
		return
	}

	ship.Facing = s.Facing
	ship.Location = s.Location
	ship.Waypoint = s.Waypoint
	switch action {
	case 'N':
		ship.Waypoint.Y += value
	case 'S':
		ship.Waypoint.Y -= value
	case 'E':
		ship.Waypoint.X += value
	case 'W':
		ship.Waypoint.X -= value
	case 'L':
		turns := value / 90
		oldx := ship.Waypoint.X
		oldy := ship.Waypoint.Y
		switch turns {
		case 1:
			ship.Waypoint.X = -oldy
			ship.Waypoint.Y = oldx
		case 2:
			ship.Waypoint.X = -oldx
			ship.Waypoint.Y = -oldy
		case 3:
			ship.Waypoint.X = oldy
			ship.Waypoint.Y = -oldx
		}
	case 'R':
		turns := value / 90
		oldx := ship.Waypoint.X
		oldy := ship.Waypoint.Y
		switch turns {
		case 1:
			ship.Waypoint.X = oldy
			ship.Waypoint.Y = -oldx
		case 2:
			ship.Waypoint.X = -oldx
			ship.Waypoint.Y = -oldy
		case 3:
			ship.Waypoint.X = -oldy
			ship.Waypoint.Y = oldx
		}
	case 'F':
		ship.Location.X += ship.Waypoint.X * value
		ship.Location.Y += ship.Waypoint.Y * value
	}
	return
}

func (s Ship) doNavigationList(instructions []string) (ship Ship, err error) {
	ship.Facing = s.Facing
	ship.Location = s.Location
	ship.Waypoint = s.Waypoint
	for _, instruction := range instructions {
		ship, err = ship.doNavigation(instruction)
		if err != nil {
			return
		}
	}
	return
}

func (s Ship) doNavigationByWaypointList(instructions []string) (ship Ship, err error) {
	ship.Facing = s.Facing
	ship.Location = s.Location
	ship.Waypoint = s.Waypoint
	for _, instruction := range instructions {
		ship, err = ship.doNavigationByWaypoint(instruction)
		if err != nil {
			return
		}
	}
	return
}

func getInput(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error reading file: %q", err))
	}

	return strings.Split(string(data), "\n"), nil
}

func manhattanDistance(loc1 Coordinates, loc2 Coordinates) int {
	return int(math.Abs(float64(loc1.X) - float64(loc2.X)) +
		math.Abs(float64(loc1.Y) - float64(loc2.Y)))
}

func part1(lines []string) (result int, err error) {
	ship := Ship {
		Facing: East,
		Location: Coordinates{
			X: 0,
			Y: 0,
		},
		Waypoint: Coordinates{
			X: 10,
			Y: 1,
		},
	}
	ship, err = ship.doNavigationList(lines)
	if err != nil {
		return
	}
	result = manhattanDistance(Coordinates{ X: 0, Y: 0 }, ship.Location)
	return
}

func part2(lines []string) (result int, err error) {
	ship := Ship {
		Facing: East,
		Location: Coordinates{
			X: 0,
			Y: 0,
		},
		Waypoint: Coordinates{
			X: 10,
			Y: 1,
		},
	}
	ship, err = ship.doNavigationByWaypointList(lines)
	if err != nil {
		return
	}
	result = manhattanDistance(Coordinates{ X: 0, Y: 0 }, ship.Location)
	return
}

func main() {
	lines, err := getInput("day12_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	part1Result, err := part1(lines)
	if err != nil {
		fmt.Printf("Received unexpected error: %v\n", err)
	}

	part2Result, err := part2(lines)
	if err != nil {
		fmt.Printf("Received unexpected error: %v\n", err)
	}

	fmt.Printf("Part 1 answer: %v\n", part1Result)
	fmt.Printf("Part 2 answer: %v\n", part2Result)
}
