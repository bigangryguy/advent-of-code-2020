package main

import (
	"fmt"
	"testing"
)

func areShipsSame(s1 Ship, s2 Ship) bool {
	return s1.Facing == s2.Facing &&
		s1.Location.X == s2.Location.X &&
		s1.Location.Y == s2.Location.Y &&
		s1.Waypoint.X == s2.Waypoint.X &&
		s1.Waypoint.Y == s2.Waypoint.Y
}

func TestShip_doNavigation(t *testing.T) {
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
	instructions := []string { "N9", "S8", "E7", "W6", "L90", "R270", "F30" }
	expected := []Ship {
		Ship {
			Facing: East,
			Location: Coordinates{
				X: 0,
				Y: 9,
			},
			Waypoint: Coordinates{
				X: 10,
				Y: 1,
			},
		},
		Ship {
			Facing: East,
			Location: Coordinates{
				X: 0,
				Y: -8,
			},
			Waypoint: Coordinates{
				X: 10,
				Y: 1,
			},
		},
		Ship {
			Facing: East,
			Location: Coordinates{
				X: 7,
				Y: 0,
			},
			Waypoint: Coordinates{
				X: 10,
				Y: 1,
			},
		},
		Ship {
			Facing: East,
			Location: Coordinates{
				X: -6,
				Y: 0,
			},
			Waypoint: Coordinates{
				X: 10,
				Y: 1,
			},
		},
		Ship {
			Facing: North,
			Location: Coordinates{
				X: 0,
				Y: 0,
			},
			Waypoint: Coordinates{
				X: 10,
				Y: 1,
			},
		},
		Ship {
			Facing: North,
			Location: Coordinates{
				X: 0,
				Y: 0,
			},
			Waypoint: Coordinates{
				X: 10,
				Y: 1,
			},
		},
		Ship {
			Facing: East,
			Location: Coordinates{
				X: 30,
				Y: 0,
			},
			Waypoint: Coordinates{
				X: 10,
				Y: 1,
			},
		},
	}
	for i, instruction := range instructions {
		actual, err := ship.doNavigation(instruction)
		if err != nil {
			t.Error("Received unexpected error")
		}
		if !areShipsSame(actual, expected[i]) {
			t.Errorf("doNavigation = %v, expected %v for instruction %s\n",
				actual, expected[i], instruction)
		}
	}
}

func TestShip_doNavigationList(t *testing.T) {
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
	instructions := []string { "N9", "S8", "E7", "W6", "L90", "R270", "F30" }
	expected := Ship {
		Facing: West,
		Location: Coordinates{
			X: -29,
			Y: 1,
		},
		Waypoint: Coordinates{
			X: 10,
			Y: 1,
		},
	}
	actual, err := ship.doNavigationList(instructions)
	if err != nil {
		t.Error("Received unexpected error")
	}
	if !areShipsSame(actual, expected) {
		t.Errorf("doNavigationList = %v, expected %v\n", actual, expected)
	}
}

func TestShip_doNavigationByWaypoint(t *testing.T) {
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
	instructions := []string { "N9", "S8", "E7", "W6", "L90", "R270", "F30" }
	expected := []Ship {
		Ship {
			Facing: East,
			Location: Coordinates{
				X: 0,
				Y: 0,
			},
			Waypoint: Coordinates{
				X: 10,
				Y: 10,
			},
		},
		Ship {
			Facing: East,
			Location: Coordinates{
				X: 0,
				Y: 0,
			},
			Waypoint: Coordinates{
				X: 10,
				Y: -7,
			},
		},
		Ship {
			Facing: East,
			Location: Coordinates{
				X: 0,
				Y: 0,
			},
			Waypoint: Coordinates{
				X: 17,
				Y: 1,
			},
		},
		Ship {
			Facing: East,
			Location: Coordinates{
				X: 0,
				Y: 0,
			},
			Waypoint: Coordinates{
				X: 4,
				Y: 1,
			},
		},
		Ship {
			Facing: East,
			Location: Coordinates{
				X: 0,
				Y: 0,
			},
			Waypoint: Coordinates{
				X: -1,
				Y: 10,
			},
		},
		Ship {
			Facing: East,
			Location: Coordinates{
				X: 0,
				Y: 0,
			},
			Waypoint: Coordinates{
				X: -1,
				Y: 10,
			},
		},
		Ship {
			Facing: East,
			Location: Coordinates{
				X: 300,
				Y: 30,
			},
			Waypoint: Coordinates{
				X: 10,
				Y: 1,
			},
		},
	}
	for i, instruction := range instructions {
		actual, err := ship.doNavigationByWaypoint(instruction)
		if err != nil {
			t.Error("Received unexpected error")
		}
		if !areShipsSame(actual, expected[i]) {
			t.Errorf("doNavigationByWaypoint = %v, expected %v for instruction %s\n",
				actual, expected[i], instruction)
		}
	}
}

func TestShip_doNavigationByWaypointList(t *testing.T) {
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
	instructions := []string { "N9", "S8", "E7", "W6", "L90", "R270", "F30" }
	expected := Ship {
		Facing: East,
		Location: Coordinates{
			X: -330,
			Y: -60,
		},
		Waypoint: Coordinates{
			X: -11,
			Y: -2,
		},
	}
	actual, err := ship.doNavigationByWaypointList(instructions)
	if err != nil {
		t.Error("Received unexpected error")
	}
	if !areShipsSame(actual, expected) {
		t.Errorf("doNavigationByWaypointList = %v, expected %v\n", actual, expected)
	}
}

func Test_part1(t *testing.T) {
	lines, err := getInput("day12_test_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	expected := 25
	actual, err := part1(lines)
	if err != nil {
		t.Error("Received unexpected error")
	}
	if actual != expected {
		t.Errorf("part1 = %d, expected %d\n", actual, expected)
	}
}

func Test_part2(t *testing.T) {
	lines, err := getInput("day12_test_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	expected := 286
	actual, err := part2(lines)
	if err != nil {
		t.Error("Received unexpected error")
	}
	if actual != expected {
		t.Errorf("part1 = %d, expected %d\n", actual, expected)
	}
}
