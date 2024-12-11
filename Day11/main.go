package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"utils"
)

func countDigits(stone int64) int {
	digitCount := 0
	for stone > 0 {
		stone /= 10
		digitCount++
	}

	return digitCount
}

func recurseBlink(timesLeft int, stone int64) int {
	if timesLeft == 0 {
		return 1
	}

	switch stone {
	case 0:
		return recurseBlink(timesLeft-1, 1)
	default:
		digits := countDigits(stone)
		if digits%2 == 0 {
			// Split the number
			halfDigits := digits / 2
			divisor := int64(math.Pow10(halfDigits))
			part1 := stone / divisor
			part2 := stone % divisor

			return recurseBlink(timesLeft-1, part1) + recurseBlink(timesLeft-1, part2)
		} else {
			return recurseBlink(timesLeft-1, stone*2024)
		}
	}
}

func threadedWrap(timesLeft int, stone int64, c chan int) {
	c <- recurseBlink(timesLeft, stone)
}

func main() {
	data := utils.ReadFile("data.txt")
	stones := strings.Split(string(data), " ")

	start := time.Now()
	result := 0
	c := make(chan int)
	for _, stone := range stones {
		numStone, err := strconv.Atoi(stone)
		utils.ThrowErr(err)
		go threadedWrap(25, int64(numStone), c)
	}
	for range len(stones) {
		result += <-c
	}
	fmt.Println("Star1 count: ", result)
	fmt.Println("Time 1: ", time.Since(start))

	start = time.Now()
	result = 0
	for _, stone := range stones {
		numStone, err := strconv.Atoi(stone)
		utils.ThrowErr(err)
		go threadedWrap(75, int64(numStone), c)
	}
	for range len(stones) {
		result += <-c
	}
	fmt.Println("Star2 count: ", result)
	fmt.Println("Time 2: ", time.Since(start))
}
