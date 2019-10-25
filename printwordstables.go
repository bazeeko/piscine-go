package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(table []string) {
	for i := range table {
		word := []rune(table[i])
		for j := range word {
			z01.PrintRune(word[j])
		}
		z01.PrintRune(10)
	}
}
