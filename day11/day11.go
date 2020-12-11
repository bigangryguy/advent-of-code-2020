package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

const Floor string = "."
const Empty string = "L"
const Occupied string = "#"

type Layout struct {
	Width, Height int
	Seats []string
}

func (l Layout) seatAt(x int, y int) string {
	return l.Seats[x + (y * l.Height)]
}

func (l Layout) setSeatAt(x int, y int, value string) {
	l.Seats[x + (y * l.Height)] = value
}

func (l Layout) occupiedSeats() (count int) {
	for _, seat := range l.Seats {
		if seat == Occupied {
			count++
		}
	}
	return
}

func (l Layout) copy() Layout {
	seatCopy := make([]string, len(l.Seats))
	copy(seatCopy, l.Seats)
	return Layout { Width: l.Width, Height: l.Height, Seats: seatCopy }
}

func (l Layout) print() {
	for y := 0; y < l.Height; y++ {
		for x := 0; x < l.Width; x++ {
			fmt.Printf("%s", l.seatAt(x, y))
		}
		fmt.Println()
	}
	fmt.Println()
}

func areLayoutsSame(l1 Layout, l2 Layout) bool {
	same := l1.Width == l2.Width && l1.Height == l2.Height && len(l1.Seats) == len(l2.Seats)
	if same {
		for i := 0; i < len(l1.Seats); i++ {
			if l1.Seats[i] != l2.Seats[i] {
				same = false
				break
			}
		}
	}
	return same
}

func getInput(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error reading file: %q", err))
	}

	return strings.Split(string(data), "\n"), nil
}

func parseInput(lines []string) Layout {
	width := len(lines[0])
	height := len(lines)
	seats := make([]string, height * width)
	for i, line := range lines {
		for j, char := range line {
			seats[j + (i * height)] = string(char)
		}
	}
	return Layout{ Width: width, Height: height, Seats: seats }
}

func intMax(i1 int, i2 int) int {
	if i1 >= i2 {
		return i1
	}
	return i2
}

func intMin(i1 int, i2 int) int {
	if i1 <= i2 {
		return i1
	}
	return i2
}

func applyRules(layout Layout, tolerance int) Layout {
	applied := layout.copy()
	for y := 0; y < layout.Height; y++ {
		for x := 0; x < layout.Width; x++ {
			seat := layout.seatAt(x, y)
			if seat == Floor {
				continue
			} else {
				starty := intMax(0, y-1)
				endy := intMin(y+1, layout.Height-1)
				startx := intMax(0, x-1)
				endx := intMin(x+1, layout.Width-1)
				var seatsOccupied int
				for ay := starty; ay <= endy; ay++ {
					for ax := startx; ax <= endx; ax++ {
						if ax == x && ay == y {
							continue
						}
						if layout.seatAt(ax, ay) == Occupied {
							seatsOccupied++
						}
					}
				}
				if seatsOccupied == 0 && seat == Empty {
					applied.setSeatAt(x, y, Occupied)
				} else if seatsOccupied >= tolerance && seat == Occupied {
					applied.setSeatAt(x, y, Empty)
				}
			}
		}
	}
	return applied
}

func applyRulesWithRange(layout Layout, tolerance int) Layout {
	applied := layout.copy()
	for y := 0; y < layout.Height; y++ {
		for x:= 0; x < layout.Width; x++ {
			seat := layout.seatAt(x, y)
			if seat == Floor {
				continue
			} else {
				var seatsOccupied int
				// Up
				for ay := intMax(0, y-1); ay >= 0; ay-- {
					if ay == y {
						break
					}
					if layout.seatAt(x, ay) != Floor {
						if layout.seatAt(x, ay) == Occupied {
							seatsOccupied++
						}
						break
					}
				}
				// Down
				for ay := intMin(y+1, layout.Height-1); ay < layout.Height; ay++ {
					if ay == y {
						break
					}
					if layout.seatAt(x, ay)  != Floor {
						if layout.seatAt(x, ay) == Occupied {
							seatsOccupied++
						}
						break
					}
				}
				// Left
				for ax := intMax(0, x-1); ax >= 0; ax-- {
					if ax == x {
						break
					}
					if layout.seatAt(ax, y) != Floor {
						if layout.seatAt(ax, y) == Occupied {
							seatsOccupied++
						}
						break
					}
				}
				// Right
				for ax := intMin(x+1, layout.Width-1); ax < layout.Width; ax++ {
					if ax == x {
						break
					}
					if layout.seatAt(ax, y) != Floor {
						if layout.seatAt(ax, y) == Occupied {
							seatsOccupied++
						}
						break
					}
				}
				// Up/Left
				ax := x-1
				ay := y-1
				for {
					if ax < 0 || ay < 0 {
						break
					}
					if layout.seatAt(ax, ay) != Floor {
						if layout.seatAt(ax, ay) == Occupied {
							seatsOccupied++
						}
						break
					}
					ax--
					ay--
				}
				// Up/Right
				ax = x+1
				ay = y-1
				for {
					if ax >= layout.Width || ay < 0 {
						break
					}
					if layout.seatAt(ax, ay) != Floor {
						if layout.seatAt(ax, ay) == Occupied {
							seatsOccupied++
						}
						break
					}
					ax++
					ay--
				}
				// Down/Left
				ax = x-1
				ay = y+1
				for {
					if ax < 0 || ay >= layout.Height {
						break
					}
					if layout.seatAt(ax, ay) != Floor {
						if layout.seatAt(ax, ay) == Occupied {
							seatsOccupied++
						}
						break
					}
					ax--
					ay++
				}
				// Down/Right
				ax = x+1
				ay = y+1
				for {
					if ax >= layout.Width || ay >= layout.Height {
						break
					}
					if layout.seatAt(ax, ay) != Floor {
						if layout.seatAt(ax, ay) == Occupied {
							seatsOccupied++
						}
						break
					}
					ax++
					ay++
				}
				if seatsOccupied == 0 && seat == Empty {
					applied.setSeatAt(x, y, Occupied)
				} else if seatsOccupied >= tolerance && seat == Occupied {
					applied.setSeatAt(x, y, Empty)
				}
			}
		}
	}
	return applied
}

func applyRulesUntilStable(layout Layout, tolerance int) Layout {
	prevLayout := layout.copy()
	var applied Layout
	for {
		applied = applyRules(prevLayout, tolerance)
		if areLayoutsSame(prevLayout, applied) {
			break
		}
		prevLayout = applied.copy()
	}
	return applied
}

func applyRulesWithRangeUntilStable(layout Layout, tolerance int) Layout {
	prevLayout := layout.copy()
	var applied Layout
	for {
		applied = applyRulesWithRange(prevLayout, tolerance)
		if areLayoutsSame(prevLayout, applied) {
			break
		}
		prevLayout = applied.copy()
	}
	return applied
}

func part1(lines []string) int {
	layout := parseInput(lines)
	applied := applyRulesUntilStable(layout, 4)
	return applied.occupiedSeats()
}

func part2(lines []string) int {
	layout := parseInput(lines)
	applied := applyRulesWithRangeUntilStable(layout, 5)
	return applied.occupiedSeats()
}

func main() {
	lines, err := getInput("day11_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	part1Result := part1(lines)
	fmt.Printf("Part 1 answer: %d\n", part1Result)

	part2Result := part2(lines)
	fmt.Printf("Part 2 answer: %d\n", part2Result)
}
