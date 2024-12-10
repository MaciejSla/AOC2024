package utils

import (
	"fmt"
	"os"
)

type Pair[T, U any] struct {
	First  T
	Second U
}

type Point struct {
	X int
	Y int
}

func (c *Point) Subtract(c2 Point) Point {
	return Point{c.X - c2.X, c.Y - c2.Y}
}

func (c *Point) Add(c2 Point) Point {
	return Point{c.X + c2.X, c.Y + c2.Y}
}

func (c *Point) InBounds(bounds Point) bool {
	return c.X <= bounds.X && c.X >= 0 && c.Y <= bounds.Y && c.Y >= 0
}

func ThrowErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ReadFile(name string) []byte {
	data, err := os.ReadFile(name)
	ThrowErr(err)
	return data
}
