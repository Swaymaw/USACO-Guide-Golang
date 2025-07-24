package main

import "testing"

func TestGetMinDistance(t *testing.T) {
	type testCase struct {
		nums []int
		want int
	}

	tt := map[string]testCase{
		"sample input": {
			nums: []int{4, 7, 8, 6, 4},
			want: 48,
		},
		"minimal": {
			nums: []int{1, 1, 1},
			want: 3,
		},
		"one dominant": {
			nums: []int{1, 100, 1},
			want: 3,
		},
		"clustered": {
			nums: []int{10, 1, 10, 1},
			want: 24,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := getMinDistance(tc.nums)

			if got != tc.want {
				t.Fatalf("Failed for test case: %s with got: %d and expected: %d", name, got, tc.want)
			}
		})
	}

}
