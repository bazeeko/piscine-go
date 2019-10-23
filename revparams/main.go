package main

import "os"
import "github.com/01-edu/z01"

func main() {
	length := 0
	for range os.Args {
		length++
	}
	for index := length - 1; index > 0; index-- {
		for _, char := range os.Args[index] {
			z01.PrintRune(char)
		}
		z01.PrintRune(10)
	}
}
