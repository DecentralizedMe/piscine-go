package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func printFileContent(file string) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		z01.PrintRune('E')
		z01.PrintRune('R')
		z01.PrintRune('R')
		z01.PrintRune('O')
		z01.PrintRune('R')
		z01.PrintRune(':')
		z01.PrintRune(' ')
		for _, char := range err.Error() {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
		os.Exit(1)
	}
	for _, char := range data {
		z01.PrintRune(rune(char))
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		data := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(data)
			if err != nil {
				if err.Error() == "EOF" {
					break
				}
				z01.PrintRune('E')
				z01.PrintRune('R')
				z01.PrintRune('R')
				z01.PrintRune('O')
				z01.PrintRune('R')
				z01.PrintRune(':')
				z01.PrintRune(' ')
				for _, char := range err.Error() {
					z01.PrintRune(char)
				}
				z01.PrintRune('\n')
				os.Exit(1)
			}
			if n == 0 {
				break
			}
			for _, char := range data[:n] {
				z01.PrintRune(rune(char))
			}
		}
	} else {
		for _, arg := range args {
			printFileContent(arg)
		}
	}
}
