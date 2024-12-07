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
	sum2 := 0
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

		// Star 1
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

		// Star 2
		solutions2 := []int{rest[0]}
		for _, val := range rest[1:] {
			tempSolutions := []int{}
			for _, solution := range solutions2 {
				tempSolutions = append(tempSolutions, solution*val)
				tempSolutions = append(tempSolutions, solution+val)
				concat, err := strconv.Atoi(fmt.Sprintf("%d%d", solution, val))
				utils.ThrowErr(err)
				tempSolutions = append(tempSolutions, concat)
			}
			solutions2 = tempSolutions
		}
		if slices.Contains(solutions2, testVal) {
			sum2 += testVal
		}
	}
	fmt.Println("Star1 sum: ", sum)
	fmt.Println("Star2 sum: ", sum2)

}
