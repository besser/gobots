package main

import (
	"fmt"
	"math"
)

type position struct {
	x, y int
}

func (p position) isNeighbour(x, y int) (is bool) {
	if !(p.x == x && p.y == y) && math.Abs((float64)(p.x-x)) <= 1 && math.Abs((float64)(p.y-y)) <= 1 {
		is = true
	}
	return is
}

func main() {
	fmt.Println("Hello, Bots!")
}
