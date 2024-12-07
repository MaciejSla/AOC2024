package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"utils"
)

func main() {
	data := utils.ReadFile("data.txt")
	lines := strings.Split(string(data), "\r\n")

	sum := 0
	for _, line := range lines {
		split := strings.Split(line, ": ")
		testVal, err := strconv.Atoi(split[0])
		utils.ThrowErr(err)

		var rest []int
		for _, val := range strings.Split(split[1], " ") {
			intVal, err := strconv.Atoi(val)
			utils.ThrowErr(err)
			rest = append(rest, intVal)
		}
		solutions := []int{rest[0]}
		for _, val := range rest[1:] {
			tempSolutions := []int{}
			for _, solution := range solutions {
				tempSolutions = append(tempSolutions, solution*val)
				tempSolutions = append(tempSolutions, solution+val)
			}
			solutions = tempSolutions
		}
		if slices.Contains(solutions, testVal) {
			sum += testVal
		}
	}
	fmt.Println("Star1 sum: ", sum)
}
