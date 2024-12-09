package main

import (
	"fmt"
	"strconv"
	"utils"
)

func findLastFileIndex(arr []int) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] >= 0 {
			return i
		}
	}
	return -1
}

func printString(arr []int) {
	for _, val := range arr {
		if val == -1 || val == -2 {
			fmt.Print(".")
		} else {
			fmt.Print(val)
		}
	}
	fmt.Println()
}

func main() {
	diskMap := utils.ReadFile("data.txt")
	var mapped []int
	idNumber := 0

	for i, char := range string(diskMap) {
		num, err := strconv.Atoi(string(char))
		utils.ThrowErr(err)
		var val int
		switch i % 2 {
		case 0:
			val = idNumber
			idNumber++
		case 1:
			val = -1
		}
		for range num {
			mapped = append(mapped, val)
		}
	}
	// printString(mapped)

	newMapped := make([]int, len(mapped))
	copy(newMapped, mapped)

	// Star 1
outer:
	for i, val := range newMapped {
		switch val {
		case -1:
			lastId := findLastFileIndex(newMapped)
			newMapped[i], newMapped[lastId] = newMapped[lastId], -2
		case -2:
			break outer
		}
	}
	// printString(newMapped)

	checksum := 0
	for i, val := range newMapped {
		if val < 0 {
			break
		}
		checksum += i * val
	}

	fmt.Println("Star1 checksum : ", checksum)
}
