package main

import "os"
import "github.com/01-edu/z01"

func main() {
	programname := os.Args[0]
	for _, char := range programname {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}
