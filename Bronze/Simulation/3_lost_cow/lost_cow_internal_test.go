package main

import "testing"

func TestGetSteps(t *testing.T) {
	type testCase struct {
		x    int
		y    int
		want int
	}

	tt := map[string]testCase{
		"sample input": {
			x:    3,
			y:    6,
			want: 9,
		},
		"custom input": {
			x:    4,
			y:    -2,
			want: 20,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := getSteps(tc.x, tc.y)
			if got != tc.want {
				t.Fatalf("Test case %s for input x: %d, y: %d failed with got: %d and expected: %d", name, tc.x, tc.y, got, tc.want)
			}
		})
	}

}
