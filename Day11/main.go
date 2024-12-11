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

// Previous (bruteforce) solution

// func recurseBlink(timesLeft int, stone int64) int {
// 	if timesLeft == 0 {
// 		return 1
// 	}

// 	switch stone {
// 	case 0:
// 		return recurseBlink(timesLeft-1, 1)
// 	default:
// 		digits := countDigits(stone)
// 		if digits%2 == 0 {
// 			// Split the number
// 			halfDigits := digits / 2
// 			divisor := int64(math.Pow10(halfDigits))
// 			part1 := stone / divisor
// 			part2 := stone % divisor

// 			return recurseBlink(timesLeft-1, part1) + recurseBlink(timesLeft-1, part2)
// 		} else {
// 			return recurseBlink(timesLeft-1, stone*2024)
// 		}
// 	}
// }

// func threadedWrap(timesLeft int, stone int64, c chan int) {
// 	c <- recurseBlink(timesLeft, stone)
// }

func blink(stones []string, blinks int) int {
	result := 0
	stoneMap := make(map[int64]int)
	for _, stone := range stones {
		numStone, err := strconv.Atoi(stone)
		utils.ThrowErr(err)
		stoneMap[int64(numStone)] += 1
	}
	for range blinks {
		newMap := make(map[int64]int)
		for stone := range stoneMap {
			digits := countDigits(int64(stone))
			if stone == 0 {
				newMap[1] += stoneMap[0]
			} else if digits%2 == 0 {
				halfDigits := digits / 2
				divisor := int64(math.Pow10(halfDigits))
				part1 := stone / divisor
				part2 := stone % divisor

				newMap[part1] += stoneMap[stone]
				newMap[part2] += stoneMap[stone]
			} else {
				newMap[int64(stone*2024)] = stoneMap[stone]
			}
		}
		stoneMap = newMap
	}
	for _, value := range stoneMap {
		result += value
	}
	return result
}

func main() {
	data := utils.ReadFile("data.txt")
	stones := strings.Split(string(data), " ")

	start := time.Now()
	result := blink(stones, 25)
	fmt.Println("Star1 count: ", result)
	fmt.Println("Time 1: ", time.Since(start))

	start = time.Now()
	result = blink(stones, 75)
	fmt.Println("Star2 count: ", result)
	fmt.Println("Time 2: ", time.Since(start))
}
