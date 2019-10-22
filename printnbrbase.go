package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	var negative bool
	if nbr < 0 {
		nbr = -nbr
		negative = true
	} else {
		negative = false
	}
	baseC := []rune(base)
	length := RuneArrayLength(baseC)
	for i := range baseC {
		for j := i + 1; j < length; j++ {
			if baseC[j] == '+' || baseC[j] == '-' || baseC[i] == baseC[j] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}
	if length < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	result := ""
	for nbr > 0 {
		temp := nbr % length
		result = string(baseC[temp]) + result
		nbr = nbr / length
	}
	answer := []rune(result)
	if negative == true {
		z01.PrintRune('-')
	}
	for _, char := range answer {
		z01.PrintRune(char)
	}
}
