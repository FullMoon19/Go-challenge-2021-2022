package piscine

func AlphaCount(s string) int {
	count := 0
	for _, r := range s {
		for compteur := 97; compteur <= 122; compteur++ {
			if string(r) == string(rune(compteur)) {
				count++
			}
		}
		for compteur := 65; compteur <= 90; compteur++ {
			if string(r) == string(rune(compteur)) {
				count++
			}
		}
	}
	return count
}
