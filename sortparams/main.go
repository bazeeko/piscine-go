package main

import "os"
import "github.com/01-edu/z01"

func SortSlice(slice []string, length int) {
	changed := false
	for index := 1; index < length-1; index++ {
		if slice[index] > slice[index+1] {
			temp := slice[index]
			slice[index] = slice[index+1]
			slice[index+1] = temp
			changed = true
		}
	}
	if changed == true {
		SortSlice(slice, length)
	}
}

func main() {
	length := 0
	for range os.Args {
		length++
	}
	slice := os.Args
	SortSlice(slice, length)
	for index := 1; index < length; index++ {
		for _, char := range slice[index] {
			z01.PrintRune(char)
			z01.PrintRune(10)
		}
	}
}
