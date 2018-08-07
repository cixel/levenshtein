package levenshtein

// Ratio is the % difference between two strings
func Ratio(s, t string) int {
	sRunes := []rune(s)
	m := len(sRunes)

	tRunes := []rune(t)
	n := len(tRunes)

	length := m + n
	dist := ComputeDistance(s, t)
	return int((1 - (float32(dist) / float32(length))) * 100)
}

// PartialRatio finds the ratio of the most similar substring,
// within the longer string, of the length of the shorter string
func PartialRatio(s, t string) int {
	// normalize, s always <= t
	if len(s) > len(t) {
		temp := s
		s = t
		t = temp
	}

	var highScore int
	for i := 0; i < len(t)-len(s)+1; i++ {
		substring := t[i : i+len(s)]
		ratio := Ratio(s, substring)
		if ratio > highScore {
			highScore = ratio
		}
	}

	return highScore
}
