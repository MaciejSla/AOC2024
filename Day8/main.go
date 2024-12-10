package main

import (
	"fmt"
	"strings"
	"utils"
)

func main() {
	data := utils.ReadFile("data.txt")
	lines := strings.Split(string(data), "\r\n")
	frequencies := map[rune][]utils.Point{}
	for y, line := range lines {
		for x, char := range line {
			if char != '.' {
				frequencies[char] = append(frequencies[char], utils.Point{X: x, Y: y})
			}
		}
	}
	bounds := utils.Point{X: len(lines) - 1, Y: len(lines[0]) - 1}

	// Star1
	antinodes := map[utils.Point]bool{}
	for _, antennaArr := range frequencies {
		for i, antenna := range antennaArr {
			for j, antenna2 := range antennaArr {
				if i == j {
					continue
				}
				potential1 := antenna.Add(antenna.Subtract(antenna2))
				if potential1.InBounds(bounds) {
					antinodes[potential1] = true
				}
				potential2 := antenna2.Add(antenna2.Subtract(antenna))
				if potential2.InBounds(bounds) {
					antinodes[potential2] = true
				}
			}
		}

	}
	// Star2
	antinodes2 := map[utils.Point]bool{}
	for _, antennaArr := range frequencies {
		for i, antenna := range antennaArr {
			antinodes2[antenna] = true
			for j, antenna2 := range antennaArr {
				if i == j {
					continue
				}
				offset1 := antenna.Subtract(antenna2)
				offset2 := antenna2.Subtract(antenna)
				potential1 := antenna.Add(offset1)
				potential2 := antenna2.Add(offset2)
				for {
					if potential1.InBounds(bounds) {
						antinodes2[potential1] = true
						potential1 = potential1.Add(offset1)
					} else {
						break
					}
				}
				for {
					if potential2.InBounds(bounds) {
						antinodes2[potential2] = true
						potential2 = potential2.Add(offset2)
					} else {
						break
					}
				}
			}
		}
	}

	fmt.Println("Star1 count: ", len(antinodes))
	fmt.Println("Star2 count: ", len(antinodes2))
}
