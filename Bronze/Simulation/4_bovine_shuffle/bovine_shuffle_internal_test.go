package main

import "testing"

func equalSlice(src, target []int) bool {
	if len(src) != len(target) {
		return false
	}
	for idx := range len(target) {
		if src[idx] != target[idx] {
			return false
		}
	}
	return true
}

func TestGetOriginalPos(t *testing.T) {
	type testCase struct {
		shuffles []int
		ids      []int
		want     []int
	}

	tt := map[string]testCase{
		"sample input": {
			shuffles: []int{1, 3, 4, 5, 2},
			ids:      []int{1234567, 2222222, 3333333, 4444444, 5555555},
			want:     []int{1234567, 5555555, 2222222, 3333333, 4444444},
		},
		"one more": {
			shuffles: []int{1, 3, 4, 5, 2, 7, 6},
			ids:      []int{1111111, 3333333, 4444444, 5555555, 2222222, 7777777, 6666666},
			want:     []int{1111111, 2222222, 3333333, 4444444, 5555555, 6666666, 7777777},
		},
		"last": {
			shuffles: []int{2, 3, 4, 1},
			ids:      []int{2345678, 3456789, 4567890, 1234567},
			want:     []int{1234567, 2345678, 3456789, 4567890},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := getOriginalPos(tc.shuffles, tc.ids)

			if !equalSlice(got, tc.want) {
				t.Fatalf("test case %s failed with got: %v and expected %v", name, got, tc.want)
			}
		})
	}
}
