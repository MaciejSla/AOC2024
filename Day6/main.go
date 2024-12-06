package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Coords[T, U any] struct {
	X T
	Y U
}

type Guard struct {
	direction         string
	startingDirection string
	startingPosition  Coords[int, int]
	position          Coords[int, int]
	visited           map[Coords[int, int]]bool
	paradoxPath       []Coords[Coords[int, int], string]
	leftMap           bool
}

func coordInBounds(coords Coords[int, int], bounds Coords[int, int]) bool {
	return coords.X <= bounds.X && coords.X >= 0 && coords.Y <= bounds.Y && coords.Y >= 0
}

func turnRight(direction string) string {
	switch direction {
	case "^":
		return ">"
	case ">":
		return "v"
	case "v":
		return "<"
	case "<":
		return "^"
	default:
		return ""
	}
}

func main() {
	directions := []string{"^", ">", "<", "v"}
	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pathMap := [][]string{}
	guard := Guard{"", "", Coords[int, int]{}, Coords[int, int]{}, make(map[Coords[int, int]]bool), []Coords[Coords[int, int], string]{}, false}
	for y, line := range strings.Split(string(data), "\r\n") {
		cells := strings.Split(line, "")
		if guard.direction == "" {
			for x, cell := range cells {
				if slices.Contains(directions, cell) {
					guard.direction = cell
					guard.startingDirection = cell
					guard.startingPosition = Coords[int, int]{x, y}
					guard.position = Coords[int, int]{x, y}
					guard.visited[guard.position] = true
					cells[x] = "."
					break
				}
			}
		}
		pathMap = append(pathMap, cells)
	}
	mapBounds := Coords[int, int]{len(pathMap[0]) - 1, len(pathMap) - 1}

	for {
		move := guard.position
		switch guard.direction {
		case "^":
			move.Y -= 1
		case ">":
			move.X += 1
		case "<":
			move.X -= 1
		case "v":
			move.Y += 1
		}
		if coordInBounds(move, mapBounds) {
			switch pathMap[move.Y][move.X] {
			case "#":
				guard.direction = turnRight(guard.direction)
			case ".":
				guard.visited[move] = true
				guard.position = move
			}
		} else {
			guard.leftMap = true
			break
		}
	}
	fmt.Println("Star1 count: ", len(guard.visited))

	paradoxCount := 0
	guard.visited[guard.startingPosition] = false

	// This is probably slow but oh well
	for position, valid := range guard.visited {
		if valid {
			guard.paradoxPath = []Coords[Coords[int, int], string]{}
			guard.position = guard.startingPosition
			guard.direction = guard.startingDirection

			// damn you shallow copies!!!!!!
			tempMap := make([][]string, len(pathMap))
			for i := range pathMap {
				tempMap[i] = make([]string, len(pathMap[i]))
				copy(tempMap[i], pathMap[i])
			}

			// About 1h spent on tempMap[position.Y][position.Y]... kill me
			tempMap[position.Y][position.X] = "#"

			for {
				move := guard.position
				switch guard.direction {
				case "^":
					move.Y -= 1
				case ">":
					move.X += 1
				case "<":
					move.X -= 1
				case "v":
					move.Y += 1
				}
				if coordInBounds(move, mapBounds) {
					coords := Coords[Coords[int, int], string]{move, guard.direction}
					if slices.Contains(guard.paradoxPath, coords) {
						paradoxCount++
						break
					}
					switch tempMap[move.Y][move.X] {
					case "#":
						guard.direction = turnRight(guard.direction)
					case ".":
						guard.paradoxPath = append(guard.paradoxPath, Coords[Coords[int, int], string]{move, guard.direction})
						guard.position = move
					}
				} else {
					break
				}
			}
		}
	}
	fmt.Println("Star2 count: ", paradoxCount)
}
