package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for a := 48; a <= '7'; a++ {
		for b := a + 1; b <= '8'; b++ {
			for c := b + 1; c <= '9'; c++ {
				if a < b {
					if b < c {
						z01.PrintRune(rune(a))
						z01.PrintRune(rune(b))
						z01.PrintRune(rune(c))
						if !(a == '7' && b == '8' && c == '9') {
							z01.PrintRune(',')
							z01.PrintRune(' ')

						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
