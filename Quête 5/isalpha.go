package piscine

func IsAlpha(s string) bool {
	for _, i := range s {
		if (i <= 96 || i >= 123) && (i < 65 || i > 90) && (i < 48 || i > 58) {
			return false
		}
	}
	return true
}
