package main

import (
	"math/rand"
	"math"
)

func main() {
	setup();
}

const (
	MINES uint = 40
	GRIDLENGTH uint = 16
)

var (
	firstClick bool
	revealed [GRIDLENGTH][GRIDLENGTH]bool
	flagged [GRIDLENGTH][GRIDLENGTH]bool
	grid [GRIDLENGTH][GRIDLENGTH]int
)

startGame() {
	for x := 0; x < GRIDLENGTH; x++ {
		for y := 0; y < GRIDLENGTH; y++ {
			revealed[x][y] = false
			flagged[x][y] = false
			grid[x][y] = 0
		}
	}

	for i := 0; i < MINES; i++ {
		x := rand.Intn(GRIDLENGTH)
		y := rand.Intn(GRIDLENGTH)

		for grid[x][y] == -1 {
			x = rand.Intn(GRIDLENGTH)
			y = rand.Intn(GRIDLENGTH)
		}

		grid[x][y] = -1

		for ix := Max(0, x - 1); ix < Min(GRIDLENGTH, x + 1); ix++ {
			for iy := Max(0, y - 1); iy < Min(GRIDLENGTH, y + 1); iy++ {
				if grid[ix][iy] != - 1 {
					grid[ix][iy]++
				}
			}
		}
	}
}