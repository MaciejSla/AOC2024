package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func isInBounds(x int, y int, data [][]string) bool {
	return x >= 0 && x < len(data) && y >= 0 && y < len(data[x])
}

func appendToWord(x int, y int, data [][]string) string {
	if isInBounds(x, y, data) {
		return data[x][y]
	} else {
		return ""
	}
}

type DirectionsXMAS struct {
	right,left,up,down,rightUp,rightDown,leftUp,leftDown string
}

// Slow ass solution bruh
func findXMAS(x int, y int, data [][]string) int {
	directions := DirectionsXMAS{
		right: "X",
		left: "X",
		up: "X",
		down: "X",
		rightUp: "X",
		rightDown: "X",
		leftUp: "X",
		leftDown: "X",
	}

	for i := 1; i < 4; i++ {
		// Right
		directions.right += appendToWord(x, y+i, data)
		// Left
		directions.left += appendToWord(x, y-i, data)
		// Up
		directions.up += appendToWord(x+i, y, data)
		// Down
		directions.down += appendToWord(x-i, y, data)
		// Right Up
		directions.rightUp += appendToWord(x+i, y+i, data)
		// Right Down
		directions.rightDown += appendToWord(x+i, y-i, data)
		// Left Up
		directions.leftUp += appendToWord(x-i, y+i, data)
		// Left Down
		directions.leftDown += appendToWord(x-i, y-i, data)
	}
	count := 0
	v := reflect.ValueOf(directions)
    for i := 0; i < v.NumField(); i++ {
        if v.Field(i).String() == "XMAS" {
			count++
		}
    }
	return count
}

func findX_MAS(x int, y int, data [][]string) int {
	leftUp := appendToWord(x-1, y+1, data)
	rightUp := appendToWord(x+1, y+1, data)
	leftDown := appendToWord(x-1, y-1, data)
	rightDown := appendToWord(x+1, y-1, data)

	diag1a := leftUp + "A" + rightDown
	diag1b := rightDown + "A" + leftUp
	diag2a := rightUp + "A" + leftDown
	diag2b := leftDown + "A" + rightUp

	if (diag1a == "MAS" || diag1b == "MAS") && (diag2a == "MAS" || diag2b == "MAS") {
		return 1
	} else {
		return 0
	}
}

func main() {
	input, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(string(input), "\n")
	data := make([][]string, len(lines))
	for i := range lines {
		data[i] = strings.Split(lines[i], "")
	}

	count1, count2 := 0, 0
	for x := range len(lines) {
		for y := range len(lines) {
			if data[x][y] == "A" {
				count2 += findX_MAS(x, y, data)
			}
			if data[x][y] == "X" {
				count1 += findXMAS(x, y, data)
			}
		}
	}
	fmt.Println("Star 1 Count: ", count1)
	fmt.Println("Star 2 Count: ", count2)
}