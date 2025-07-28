package main

import "testing"

func TestGetPossibleWins(t *testing.T) {
	type testCase struct {
		board    [3][3]rune
		want_one int
		want_two int
	}

	tt := map[string]testCase{
		"sample input": {
			board:    [3][3]rune{{'C', 'O', 'W'}, {'X', 'X', 'O'}, {'A', 'B', 'C'}},
			want_one: 0,
			want_two: 2,
		},
		"extreme": {
			board:    [3][3]rune{{'R', 'R', 'R'}, {'R', 'R', 'R'}, {'R', 'R', 'R'}},
			want_one: 1,
			want_two: 0,
		},
		"example": {
			board:    [3][3]rune{{'X', 'Y', 'Z'}, {'X', 'Y', 'Z'}, {'X', 'Y', 'X'}},
			want_one: 2,
			want_two: 2,
		},
		"example 2": {
			board:    [3][3]rune{{'Q', 'Q', 'Q'}, {'Q', 'P', 'Q'}, {'Q', 'Q', 'Q'}},
			want_one: 1,
			want_two: 1,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			one_p, two_p := getPossibleWins(tc.board)

			if one_p != tc.want_one || two_p != tc.want_two {
				t.Fatalf("Failed for test %s got %d, %d and expected %d, %d", name, one_p, two_p, tc.want_one, tc.want_two)
			}

		})
	}
}
