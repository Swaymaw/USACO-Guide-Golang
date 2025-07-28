package main

import "testing"

func TestGetStartEndEstimate(t *testing.T) {
	type testCase struct {
		sensors       []SensorInfo
		want_start_lb int
		want_start_ub int
		want_end_lb   int
		want_end_ub   int
	}

	tt := map[string]testCase{
		"sample input": {
			sensors:       []SensorInfo{SensorInfo{loc: "on", lb: 1, ub: 1}, SensorInfo{loc: "none", lb: 10, ub: 14}, SensorInfo{loc: "none", lb: 11, ub: 15}, SensorInfo{loc: "off", lb: 2, ub: 3}},
			want_start_lb: 10,
			want_start_ub: 13,
			want_end_lb:   8,
			want_end_ub:   12,
		},
		"example": {
			sensors:       []SensorInfo{SensorInfo{loc: "on", lb: 5, ub: 5}, SensorInfo{loc: "none", lb: 3, ub: 10}},
			want_start_lb: 0,
			want_start_ub: 5,
			want_end_lb:   5,
			want_end_ub:   10,
		},
		"minimal": {
			sensors:       []SensorInfo{SensorInfo{loc: "none", lb: 0, ub: 1000}},
			want_start_lb: 0,
			want_start_ub: 1000,
			want_end_lb:   0,
			want_end_ub:   1000,
		},
		"last": {
			sensors:       []SensorInfo{SensorInfo{loc: "none", lb: 0, ub: 0}, SensorInfo{loc: "on", lb: 1, ub: 1}, SensorInfo{loc: "off", lb: 1, ub: 1}, SensorInfo{loc: "none", lb: 0, ub: 0}},
			want_start_lb: 0,
			want_start_ub: 0,
			want_end_lb:   0,
			want_end_ub:   0,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			start_lb, start_ub, end_lb, end_ub := getStartEndEstimate(tc.sensors)

			if start_lb != tc.want_start_lb || start_ub != tc.want_start_ub || end_lb != tc.want_end_lb || end_ub != tc.want_end_ub {
				t.Fatalf("Test case %s failed with got:\n%d, %d\n%d, %d\nand\nexpected:\n%d, %d\n%d, %d", name, start_lb, start_ub, end_lb, end_ub, tc.want_start_lb, tc.want_start_ub, tc.want_end_lb, tc.want_end_ub)
			}
		})
	}
}
