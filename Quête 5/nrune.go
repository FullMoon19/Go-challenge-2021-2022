package piscine

func NRune(s string, n int) rune {
	lastrune := []rune(s)
	if n > 0 && n < len(s)+1 {
		return lastrune[n-1]
	}
	return 0
}
