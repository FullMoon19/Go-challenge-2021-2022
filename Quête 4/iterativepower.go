package piscine

func IterativePower(nb int, power int) int {
	value := 1
	if power < 0 {
		return 0
	} else {
		for i := 1; i <= power; i++ {
			value *= nb
		}
	}
	return value
}
