package main

import "testing"

func compareGrids(got, want []string) bool {
	if len(got) != len(want) {
		return false
	}

	for idx := range len(got) {
		if got[idx] != want[idx] {
			return false
		}
	}
	return true
}

func TestExpand(t *testing.T) {
	type testCase struct {
		k    int
		grid []string
		want []string
	}

	tt := map[string]testCase{
		"sample input": {
			k: 2,
			grid: []string{
				"XXX.",
				"X..X",
				"XXX.",
				"X..X",
				"XXX.",
			},
			want: []string{
				"XXXXXX..",
				"XXXXXX..",
				"XX....XX",
				"XX....XX",
				"XXXXXX..",
				"XXXXXX..",
				"XX....XX",
				"XX....XX",
				"XXXXXX..",
				"XXXXXX..",
			},
		},
	}

	for name, tc := range tt {
		t.Run(
			name, func(t *testing.T) {
				got := expand(tc.grid, tc.k)
				if !compareGrids(got, tc.want) {
					t.Fatalf("Failed for testCase: %s", name)
				}
			},
		)
	}

}
