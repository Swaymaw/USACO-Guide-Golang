package main

import "testing"

func TestGetMaxBuckets(t *testing.T) {
	type testCase struct {
		buckets map[int]int
		want    int
	}

	tt := map[string]testCase{
		"sample input": {
			buckets: map[int]int{4: 1, 10: -1, 8: 3, 13: -3, 2: 2, 6: -2},
			want:    4,
		},
		"one more": {
			buckets: map[int]int{2: 2, 15: -2, 4: 1, 6: -1, 3: 3, 16: -3},
			want:    6,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := getMaxBuckets(tc.buckets)

			if got != tc.want {
				t.Fatalf("Test Case Failed for %s with got: %d and expected: %d", name, got, tc.want)
			}
		})
	}
}
