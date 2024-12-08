package main

import (
	"fmt"
	"strings"
	"utils"
)

type Coords struct {
	X int
	Y int
}

func (c *Coords) sub(c2 Coords) Coords {
	return Coords{c.X - c2.X, c.Y - c2.Y}
}

func (c *Coords) add(c2 Coords) Coords {
	return Coords{c.X + c2.X, c.Y + c2.Y}
}

func (c *Coords) inBounds(bounds Coords) bool {
	return c.X <= bounds.X && c.X >= 0 && c.Y <= bounds.Y && c.Y >= 0
}

func main() {
	data := utils.ReadFile("data.txt")
	lines := strings.Split(string(data), "\r\n")
	frequencies := map[rune][]Coords{}
	for y, line := range lines {
		for x, char := range line {
			if char != '.' {
				frequencies[char] = append(frequencies[char], Coords{x, y})
			}
		}
	}
	bounds := Coords{len(lines) - 1, len(lines[0]) - 1}

	// Star1
	antinodes := map[Coords]bool{}
	for _, antennaArr := range frequencies {
		for i, antenna := range antennaArr {
			for j, antenna2 := range antennaArr {
				if i == j {
					continue
				}
				potential1 := antenna.add(antenna.sub(antenna2))
				if potential1.inBounds(bounds) {
					antinodes[potential1] = true
				}
				potential2 := antenna2.add(antenna2.sub(antenna))
				if potential2.inBounds(bounds) {
					antinodes[potential2] = true
				}
			}
		}

	}
	// Star2
	antinodes2 := map[Coords]bool{}
	for _, antennaArr := range frequencies {
		for i, antenna := range antennaArr {
			antinodes2[antenna] = true
			for j, antenna2 := range antennaArr {
				if i == j {
					continue
				}
				offset1 := antenna.sub(antenna2)
				offset2 := antenna2.sub(antenna)
				potential1 := antenna.add(offset1)
				potential2 := antenna2.add(offset2)
				for {
					if potential1.inBounds(bounds) {
						antinodes2[potential1] = true
						potential1 = potential1.add(offset1)
					} else {
						break
					}
				}
				for {
					if potential2.inBounds(bounds) {
						antinodes2[potential2] = true
						potential2 = potential2.add(offset2)
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
