package main

import "os"
import "github.com/01-edu/z01"

func main() {
	length := 0
	for range os.Args {
		length++
	}
	for index := range os.Args {
		if index == length-1 {
			break
		}
		for _, char := range os.Args[index+1] {
			z01.PrintRune(char)
		}
		z01.PrintRune(10)
	}
}
