package main

import "testing"

func TestMaxOverSpeed(t *testing.T) {
	type testCase struct {
		track_segments []Segment
		cow_segments   []Segment
		want           int
	}

	tt := map[string]testCase{
		"sample input": {
			track_segments: []Segment{Segment{length: 40, limit: 75}, Segment{length: 50, limit: 35}, Segment{length: 10, limit: 45}},
			cow_segments:   []Segment{Segment{length: 40, limit: 76}, Segment{length: 20, limit: 30}, Segment{length: 40, limit: 40}},
			want:           5,
		},
		"my test case": {
			track_segments: []Segment{Segment{length: 25, limit: 10}, Segment{length: 25, limit: 20}, Segment{length: 25, limit: 25}, Segment{length: 25, limit: 30}},
			cow_segments:   []Segment{Segment{length: 33, limit: 5}, Segment{length: 34, limit: 35}, Segment{length: 33, limit: 40}},
			want:           15,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := maxOverSpeed(tc.track_segments, tc.cow_segments)

			if got != tc.want {
				t.Fatalf("Test failed for %s. got - %d expected - %d", name, got, tc.want)
			}

		})
	}
}
