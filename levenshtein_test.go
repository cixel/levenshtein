package levenshtein

import "testing"

type testCase struct {
	a, b     string
	expected int
}

func TestDistance(t *testing.T) {
	tests := []testCase{
		{"dog", "dog", 0},
		{"kitten", "mitten", 1},
		{"kitten", "sitting", 3},
		{"kitten", "", 6},
		{"", "kitten", 6},
		{"flaw", "lawn", 2},
		{"gumbo", "gambol", 2},
		{"ehden", "dehden", 1},
		{"💯", "💯", 0},
		{"💯", "💯💯", 1},
		{"💯", "", 1},
		{"î", "î", 0},
	}

	for i, test := range tests {
		dist := ComputeDistance(test.a, test.b)
		if dist != test.expected {
			t.Errorf("[%d]: ComputeDistance(%q, %q) - returned: %v, expected: %v",
				i, test.a, test.b, dist, test.expected)
		}
	}
}
