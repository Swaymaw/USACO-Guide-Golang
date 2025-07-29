package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestGetMaxRange(t *testing.T) {
	type testCase struct {
		diamonds []int
		k        int
		want     int
	}

	tt := make(map[string]testCase)

	baseDir := "./testdata/"

	entries, _ := os.ReadDir(baseDir)

	for _, e := range entries {
		name := e.Name()

		if strings.HasSuffix(name, ".in") {
			var want int

			diamonds, k := getData(baseDir + name)
			file, _ := os.Open(baseDir + name[:len(name)-3] + ".out")

			scanner := bufio.NewScanner(file)

			if scanner.Scan() {
				parts := strings.Fields(scanner.Text())
				want, _ = strconv.Atoi(parts[0])
			}

			tt[name] = testCase{diamonds: diamonds, k: k, want: want}

		}
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := getMaxRange(tc.diamonds, tc.k)

			if got != tc.want {
				t.Fatalf("TestCase %s failed where got: %d and expected: %d", name, got, tc.want)
			}
		})
	}

}
