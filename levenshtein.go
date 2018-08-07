package levenshtein

// ComputeDistance finds the Levenshtein distances between two strings a and b
// https://en.wikipedia.org/wiki/Levenshtein_distance#Iterative_with_two_matrix_rows
func ComputeDistance(s, t string) int {
	// equal strings
	if s == t {
		return 0
	}

	sRunes := []rune(s)
	m := len(sRunes)

	tRunes := []rune(t)
	n := len(tRunes)

	// 0-length strings
	if m == 0 {
		return n
	}

	if n == 0 {
		return m
	}

	v0 := make([]int, n+1)
	v1 := make([]int, n+1)

	for i := 0; i < n+1; i++ {
		v0[i] = i
	}

	for i := 0; i < m; i++ {
		v1[0] = i + 1

		for j := 0; j < n; j++ {
			deletionCost := v0[j+1] + 1
			insertionCost := v1[j] + 1

			var substitionCost int
			if sRunes[i] == tRunes[j] {
				substitionCost = v0[j]
			} else {
				substitionCost = v0[j] + 1
			}

			v1[j+1] = min(deletionCost, insertionCost, substitionCost)
		}

		for j := 0; j < n+1; j++ {
			v0[j] = v1[j]
		}
	}

	return v1[n]
}

func min(a, b, c int) int {
	if a <= b && b <= c {
		return a
	} else if b <= c {
		return b
	}

	return c

}
