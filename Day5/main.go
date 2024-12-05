package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Rule struct {
	before, after []string
}

func orderUpdate(updateArr []string, rulebook map[string]Rule) []string {
	if len(updateArr) <= 1 {
		return updateArr
	}
	pivot := updateArr[0]
	rest := updateArr[1:]

	pageRule := rulebook[pivot]
	beforeArr := []string{}
	afterArr := []string{}

	for _, page := range rest {
		if slices.Contains(pageRule.before, page) {
			beforeArr = append(beforeArr, page)
		} else if slices.Contains(pageRule.after, page) {
			afterArr = append(afterArr, page)
		}
	}
	return append(append(orderUpdate(beforeArr, rulebook), pivot), orderUpdate(afterArr, rulebook)...)
}

func main() {
	input, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	data := strings.Split(string(input), "\n\r")
	rules, updates := strings.Split(strings.TrimSpace(data[0]), "\n"), strings.Split(strings.TrimSpace(data[1]), "\n")

	// Create a rulebook that stores the rules for each page
	rulebook := make(map[string]Rule)

	for _, rule := range rules {
		splitRule := strings.Split(strings.TrimSpace(rule), "|")
		before, after := splitRule[0], splitRule[1]

		beforeRule := rulebook[before]
		if beforeRule.after == nil {
			beforeRule.after = []string{after}
		} else {
			beforeRule.after = append(beforeRule.after, after)
		}
		rulebook[before] = beforeRule

		afterRule := rulebook[after]
		if afterRule.before == nil {
			afterRule.before = []string{before}
		} else {
			afterRule.before = append(afterRule.before, before)
		}
		rulebook[after] = afterRule
	}
	
	count1 := 0
	count2 := 0
	incorrectUpdates := []string{}

	// Star 1 - only correctly-ordered updates
	for _, update := range updates {
		valid := true
		pages := strings.Split(strings.TrimSpace(update), ",")

		outer:
		for i, page := range pages {
			pageRule := rulebook[page]
			for x := i; x >= 0; x-- {
				if slices.Contains(pageRule.after, pages[x]) {
					incorrectUpdates = append(incorrectUpdates, update)
					valid = false
					break outer
				}
			}
		}
		if valid {
			middleIndex := len(pages) / 2
			middlePage, err := strconv.Atoi(pages[middleIndex])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			count1 += middlePage
		}
	}

	// Star 2 - incorrectly-ordered updates
	for _, update := range incorrectUpdates {
		pages := strings.Split(strings.TrimSpace(update), ",")
		orderedArr := orderUpdate(pages, rulebook)
		middleIndex := len(pages) / 2
		middlePage, err := strconv.Atoi(orderedArr[middleIndex])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		count2 += middlePage
	}

	fmt.Println("Star1 count: ", count1)
	fmt.Println("Star2 count: ", count2)
}