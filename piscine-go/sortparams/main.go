package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var list []string
	for _, arg := range os.Args[1:] {
		if len(list) == 0 {
			list = append(list, arg)
		} else {
			inserted := false
			for i := 0; i < len(list); i++ {
				if arg > list[i] {
					list = append(list[:i+1], list[i:]...)
					list[i] = arg
					inserted = true
					break
				}
			}
			if !inserted {
				list = append(list, arg)
			}
		}
	}
	for i := len(list) - 1; i >= 0; i-- {
		for _, letter := range list[i] {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}
