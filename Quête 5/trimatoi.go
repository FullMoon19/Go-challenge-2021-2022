package piscine

func TrimAtoi(s string) int {
	val := 0
	minus := 0
	for _, i := range s {
		if i == '-' && val == 0 {
			minus = 1
		}
		if i >= 48 && i <= 57 {
			val = (int(i) - 48) + val*10
		}
	}
	if minus == 1 {
		return -val
	}
	return val
}
