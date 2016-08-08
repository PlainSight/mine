package main

import (
	"math/rand"
)

func main() {
	startGame()
	setup()
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
	win bool
	lose bool
	lastState bool
)

func uMin(a uint, b uint) uint {
	if a < b {
		return a
	} else {
		return b
	}
}

func uDec(a uint) uint {
	if a == 0 {
		return 0
	} else {
		return uint(a - 1)
	}
}

func startGame() {
	lastState = win && !lose
	win = false
	lose = false

	for x := uint(0); x < GRIDLENGTH; x++ {
		for y := uint(0); y < GRIDLENGTH; y++ {
			revealed[x][y] = false
			flagged[x][y] = false
			grid[x][y] = 0
		}
	}

	for i := uint(0); i < MINES; i++ {
		x := uint(rand.Intn(int(GRIDLENGTH)))
		y := uint(rand.Intn(int(GRIDLENGTH)))
		for grid[x][y] == -1 {
			x = uint(rand.Intn(int(GRIDLENGTH)))
			y = uint(rand.Intn(int(GRIDLENGTH)))
		}

		grid[x][y] = -1

		for ix := uDec(x); ix < uMin(GRIDLENGTH, x + 2); ix++ {
			for iy := uDec(y); iy < uMin(GRIDLENGTH, y + 2); iy++ {
				if grid[ix][iy] != -1 {
					grid[ix][iy]++
				}
			}
		}
	}
}

func reveal(x uint, y uint) {
	if !revealed[x][y] && !flagged[x][y] {
		revealed[x][y] = true

		if grid[x][y] == -1 {
			lose = true
		} else {
			if grid[x][y] == 0 {
				find(x, y)
			}
		}
	}
}

func find(x uint, y uint) {
	for ix := uDec(x); ix < uMin(GRIDLENGTH, x + 2); ix++ {
		for iy := uDec(y); iy < uMin(GRIDLENGTH, y + 2); iy++ {
			if ix != x || iy != y {
				reveal(ix, iy)
			}
		}
	}
}

func flagClick(x uint, y uint) {
	flagged[x][y] = !flagged[x][y]
}

func revealClick(x uint, y uint) {
	reveal(x, y)
	winCheck()
}

func winCheck() {
	counter := uint(0)
	for x := uint(0); x < GRIDLENGTH; x++ {
		for y := uint(0); y < GRIDLENGTH; y++ {
			if !revealed[x][y] {
				counter++
			}
		}
	}

	if counter == MINES {
		win = true
	}

	if lose {
		//show loss
		startGame()
	}

	if !lose && win {
		//show win
		startGame()
	}
}