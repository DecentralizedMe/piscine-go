package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

const (
	empty = '.'
	size  = 9
)

type board [size][size]int

func main() {
	args := os.Args[1:]
	if len(args) != size {
		panic("Invalid number of arguments")
	}

	b := board{}
	for i := range b {
		for j, v := range args[i] {
			if v == empty {
				b[i][j] = 0
			} else {
				n, err := strconv.Atoi(string(v))
				if err != nil {
					panic("Invalid input")
				}
				b[i][j] = n
			}
		}
	}

	if !solve(&b) {
		panic("Invalid input")
	}

	for i := range b {
		for j := range b[i] {
			z01.PrintRune(rune(b[i][j] + '0'))
			if j < size-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func solve(b *board) bool {
	var row, col int
	var found bool

	for i := range b {
		for j := range b[i] {
			if b[i][j] == 0 {
				row, col = i, j
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	if !found {
		return true
	}

	for num := 1; num <= size; num++ {
		if isValid(b, row, col, num) {
			b[row][col] = num
			if solve(b) {
				return true
			}
			b[row][col] = 0
		}
	}

	return false
}

func isValid(b *board, row, col, num int) bool {
	for i := 0; i < size; i++ {
		if b[row][i] == num {
			return false
		}
		if b[i][col] == num {
			return false
		}
	}

	rowStart := (row / 3) * 3
	colStart := (col / 3) * 3
	for i := rowStart; i < rowStart+3; i++ {
		for j := colStart; j < colStart+3; j++ {
			if b[i][j] == num {
				return false
			}
		}
	}

	return true
}
