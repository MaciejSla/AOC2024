package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

func compactPointArr(arr []utils.Point) []utils.Point {
	compacted := []utils.Point{}
	for _, trail := range arr {
		skip := false
		for _, val := range compacted {
			if trail.X == val.X && trail.Y == val.Y {
				skip = true
				break
			}
		}
		if !skip {
			compacted = append(compacted, trail)
		}
	}
	return compacted
}

func countPaths(prevVal int, coords utils.Point, tMap [][]int, bounds utils.Point) []utils.Point {
	if coords.InBounds(bounds) {
		val := tMap[coords.Y][coords.X]
		if val == prevVal+1 {
			if val == 9 {
				return []utils.Point{{X: coords.X, Y: coords.Y}}
			}
			returnArr := []utils.Point{}
			up := utils.Point{X: coords.X, Y: coords.Y - 1}
			returnArr = append(returnArr, countPaths(val, up, tMap, bounds)...)
			down := utils.Point{X: coords.X, Y: coords.Y + 1}
			returnArr = append(returnArr, countPaths(val, down, tMap, bounds)...)
			right := utils.Point{X: coords.X + 1, Y: coords.Y}
			returnArr = append(returnArr, countPaths(val, right, tMap, bounds)...)
			left := utils.Point{X: coords.X - 1, Y: coords.Y}
			returnArr = append(returnArr, countPaths(val, left, tMap, bounds)...)
			return returnArr
		} else {
			return []utils.Point{}
		}
	} else {
		return []utils.Point{}
	}
}

func main() {
	data := utils.ReadFile("data.txt")
	topographicMap := [][]int{}
	for _, val := range strings.Split(string(data), "\r\n") {
		row := []int{}
		for _, char := range val {
			numVal, err := strconv.Atoi(string(char))
			utils.ThrowErr(err)
			row = append(row, numVal)
		}
		topographicMap = append(topographicMap, row)
	}
	bounds := utils.Point{X: len(topographicMap) - 1, Y: len(topographicMap[0]) - 1}

	count1 := 0
	count2 := 0
	for y, line := range topographicMap {
		for x, val := range line {
			if val == 0 {
				trails := countPaths(-1, utils.Point{X: x, Y: y}, topographicMap, bounds)
				compacted := compactPointArr(trails)

				count1 += len(compacted)
				count2 += len(trails)
			}
		}
	}

	fmt.Println("Star1 count: ", count1)
	fmt.Println("Star2 count: ", count2)
}
