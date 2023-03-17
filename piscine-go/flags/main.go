package main

import (
	"os"
	"strings"

	"github.com/01-edu/z01"
)

func printHelp() {
	z01.WriteString(os.Stdout, "--insert\n")
	z01.WriteString(os.Stdout, "  -i\n")
	z01.WriteString(os.Stdout, "   This flag inserts the string into the string passed as argument.\n")
	z01.WriteString(os.Stdout, "--order\n")
	z01.WriteString(os.Stdout, "  -o\n")
	z01.WriteString(os.Stdout, "   This flag will behave like a boolean, if it is called it will order the argument.\n")
}

func sort(strPrint string) {
	var list []string
	for _, arg := range strPrint {
		if len(list) == 0 {
			list = append(list, string(' '))
		}
		for index, word := range list {
			if string(arg) >= word {
				list = append(list[:index+1], list[index:]...)
				list[index] = string(arg)
				break
			}
		}
	}
	for i := len(list) - 2; i >= 0; i-- {
		for _, letter := range list[i] {
			z01.PrintRune(letter)
		}
	}
}

func main() {
	var newStr string
	switch len(os.Args) {
	case 2:
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			printHelp()
		} else {
			z01.WriteString(os.Stdout, os.Args[1])
		}
	case 3:
		if os.Args[1] == "--order" || os.Args[1] == "-o" {
			sort(os.Args[2])
		} else {
			z01.WriteString(os.Stdout, os.Args[2])
			for index, match := range os.Args[1] {
				if match == '=' {
					z01.WriteString(os.Stdout, os.Args[1][index+1:])
				}
			}
		}
	case 4:
		for index, match := range os.Args[1] {
			if match == '=' {
				newStr = os.Args[3] + os.Args[1][index+1:]
			}
		}
		result := strings.Replace(os.Args[2], " ", newStr, -1)
		z01.WriteString(os.Stdout, result)
	default:
		printHelp()
	}
	z01.PrintRune('\n')
}
