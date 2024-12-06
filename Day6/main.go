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
	direction string
	position  Coords[int, int]
	visited   map[Coords[int, int]]bool
	leftMap   bool
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
	guard := Guard{"", Coords[int, int]{}, make(map[Coords[int, int]]bool), false}
	for y, line := range strings.Split(string(data), "\r\n") {
		cells := strings.Split(line, "")
		if guard.direction == "" {
			for x, cell := range cells {
				if slices.Contains(directions, cell) {
					guard.direction = cell
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

	// Debugging commented out
	// index := 0
	for {
		// if index == 10 {
		// 	break
		// }
		move := guard.position
		// fmt.Printf("Round %d - Before move: %v\n", index, move)
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
		// fmt.Printf("Round %d - After move: %v\n", index, move)
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
		// index++
	}
	fmt.Println("Star1 count: ", len(guard.visited))
}
