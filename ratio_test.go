package levenshtein

import "testing"

// https://chairnerd.seatgeek.com/fuzzywuzzy-fuzzy-string-matching-in-python/
var strings = []string{
	"new york mets",
	"new york mets",
	"new YORK mets",
	"the wonderful new york mets",
}

var silly = []string{
	"{",
	"{",
	"{a",
	"{a",
	"a{",
	"{b",
}

func TestRatio(t *testing.T) {
	tests := []testCase{
		{strings[0], strings[1], 100},
		{silly[0], silly[1], 100},
		{silly[2], silly[3], 100},
	}

	for i, test := range tests {
		dist := Ratio(test.a, test.b)
		if dist != test.expected {
			t.Errorf("[%d]: Ratio(%q, %q) - returned: %v, expected: %v",
				i, test.a, test.b, dist, test.expected)
		}
	}
}

func TestPartialRatio(t *testing.T) {
	tests := []testCase{
		{strings[0], strings[3], 100},
		{strings[3], strings[0], 100},
		{"dog", "dog", 100},
		{"dog", "", 0},
		{"xo", "zo", 75},
	}

	for i, test := range tests {
		dist := PartialRatio(test.a, test.b)
		if dist != test.expected {
			t.Errorf("[%d]: PartialRatio(%q, %q) - returned: %v, expected: %v",
				i, test.a, test.b, dist, test.expected)
		}
	}
}
