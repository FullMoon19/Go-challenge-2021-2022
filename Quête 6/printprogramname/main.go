package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := []rune(os.Args[0])[2:]
	for _, i := range argument {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
