package piscine

func IsNumeric(s string) bool {
	for _, i := range s {
		if i < 48 || i > 58 {
			return false
		}
	}
	return true
}
