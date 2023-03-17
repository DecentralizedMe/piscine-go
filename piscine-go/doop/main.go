package main

import (
	"os"
	"strconv"
)

func main() {
	var num1 int64
	var num2 int64
	var result int64
	var sign string

	if len(os.Args) != 4 {
		return
	}

	num1, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		return
	}

	num2, err = strconv.ParseInt(os.Args[3], 10, 64)
	if err != nil {
		return
	}

	sign = os.Args[2]

	switch sign {
	case "+":
		result = num1 + num2
		if (num1 > 0 && num2 > 0 && result < 0) || (num1 < 0 && num2 < 0 && result > 0) {
			return
		}
	case "-":
		result = num1 - num2
		if (num1 > 0 && num2 < 0 && result < 0) || (num1 < 0 && num2 > 0 && result > 0) {
			return
		}
	case "*":
		result = num1 * num2
		if num1 != 0 && result/num1 != num2 {
			return
		}
	case "/":
		if num2 == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		} else {
			result = num1 / num2
		}
	case "%":
		if num2 == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		} else {
			result = num1 % num2
		}
	default:
		return
	}
	bytes := PrintNbr(result)
	os.Stdout.Write(bytes)
}

func PrintNbr(n int64) []byte {
	var bytes []byte
	var runes []rune
	if n < 0 {
		bytes = append(bytes, '-')
	}
	if n == 0 {
		bytes = append(bytes, '0', '\n')
		return bytes
	} else {
		for {
			var remain int64 = n % 10
			if remain < 0 {
				remain = remain * -1
			}
			runes = append(runes, rune(remain))
			n = n / 10
			if n == 0 {
				break
			}
		}
	}
	for i := len(runes) - 1; i >= 0; i-- {
		bytes = append(bytes, byte(runes[i]+48))
	}
	bytes = append(bytes, '\n')
	return bytes
}
