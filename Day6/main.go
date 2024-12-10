package main

import (
	"fmt"
	"slices"
	"strings"
	"utils"
)

// type Coords[T, U any] struct {
// 	X T
// 	Y U
// }

type Guard struct {
	direction         string
	startingDirection string
	startingPosition  utils.Point
	position          utils.Point
	visited           map[utils.Point]bool
	paradoxPath       []utils.Pair[utils.Point, string]
	leftMap           bool
}

func pointInBounds(point utils.Point, bounds utils.Point) bool {
	return point.X <= bounds.X && point.X >= 0 && point.Y <= bounds.Y && point.Y >= 0
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
	data := utils.ReadFile("data.txt")
	pathMap := [][]string{}
	guard := Guard{"", "", utils.Point{}, utils.Point{}, make(map[utils.Point]bool), []utils.Pair[utils.Point, string]{}, false}
	for y, line := range strings.Split(string(data), "\r\n") {
		cells := strings.Split(line, "")
		if guard.direction == "" {
			for x, cell := range cells {
				if slices.Contains(directions, cell) {
					guard.direction = cell
					guard.startingDirection = cell
					guard.startingPosition = utils.Point{X: x, Y: y}
					guard.position = utils.Point{X: x, Y: y}
					guard.visited[guard.position] = true
					cells[x] = "."
					break
				}
			}
		}
		pathMap = append(pathMap, cells)
	}
	mapBounds := utils.Point{X: len(pathMap[0]) - 1, Y: len(pathMap) - 1}

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
		if pointInBounds(move, mapBounds) {
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
			guard.paradoxPath = []utils.Pair[utils.Point, string]{}
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
				if pointInBounds(move, mapBounds) {
					coords := utils.Pair[utils.Point, string]{First: move, Second: guard.direction}
					if slices.Contains(guard.paradoxPath, coords) {
						paradoxCount++
						break
					}
					switch tempMap[move.Y][move.X] {
					case "#":
						guard.direction = turnRight(guard.direction)
					case ".":
						guard.paradoxPath = append(guard.paradoxPath, utils.Pair[utils.Point, string]{First: move, Second: guard.direction})
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
