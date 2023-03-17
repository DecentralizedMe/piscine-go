package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func isPowerOfTwo(n int) bool {
	return (n != 0) && ((n & (n - 1)) == 0)
}

func main() {
	args := os.Args[1:]

	if len(args) == 1 {
		num, err := strconv.Atoi(args[0])
		if err != nil {
			return
		}
		if isPowerOfTwo(num) {
			z01.PrintRune('t')
			z01.PrintRune('r')
			z01.PrintRune('u')
			z01.PrintRune('e')
			z01.PrintRune('\n')
		} else {
			z01.PrintRune('f')
			z01.PrintRune('a')
			z01.PrintRune('l')
			z01.PrintRune('s')
			z01.PrintRune('e')
			z01.PrintRune('\n')
		}
	}
}
