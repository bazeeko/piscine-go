package main

import "os"
import "github.com/01-edu/z01"

func chartoint(c rune) int {
	if c == '0' {
		return 0
	} else if c == '1' {
		return 1
	} else if c == '2' {
		return 2
	} else if c == '3' {
		return 3
	} else if c == '4' {
		return 4
	} else if c == '5' {
		return 5
	} else if c == '6' {
		return 6
	} else if c == '7' {
		return 7
	} else if c == '8' {
		return 8
	} else if c == '9' {
		return 9
	} else {
		return -1
	}
}

func BasicAtoi(s string) int {
	str := []rune(s)
	answer := 0
	length := 0
	for range s {
		length++
	}
	for i := range str {
		if chartoint(str[i]) == -1 {
			return -1
		}
		if answer == 0 && str[i] == '0' {
			answer = 0
		} else {
			ten := 1
			for j := i; j < length-1; j++ {
				ten = ten * 10
			}
			answer = answer + chartoint(str[i])*ten
		}
	}
	return answer
}

func nbrtoalpha(nbr int) rune {
	var letter rune
	letter = 96 + rune(nbr)
	return letter
}

func main() {
	slice := os.Args
	length := 0
	for range os.Args {
		length++
	}
	dif := 0
	if slice[1] == "--upper" {
		dif = -32
	}
	for i := 1; i < length; i++ {
		if BasicAtoi(slice[i]) >= 1 && BasicAtoi(slice[i]) <= 26 {
			z01.PrintRune(nbrtoalpha(BasicAtoi(slice[i]) + dif))
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune(10)
}
