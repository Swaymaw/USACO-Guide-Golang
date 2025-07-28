package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestGetMaxFill(t *testing.T) {
	type testCase struct {
		x    int
		y    int
		M    int
		want int
	}

	tt := make(map[string]testCase)
	dir := "./testdata/"
	entries, _ := os.ReadDir(dir)

	for _, e := range entries {
		name := e.Name()
		if strings.HasSuffix(name, ".in") {
			var x, y, M, want int
			inFile, _ := os.Open(dir + name)
			defer inFile.Close()
			inScanner := bufio.NewScanner(inFile)

			if inScanner.Scan() {
				parts := strings.Fields(inScanner.Text())

				x, _ = strconv.Atoi(parts[0])
				y, _ = strconv.Atoi(parts[1])
				M, _ = strconv.Atoi(parts[2])
			}

			outFile, _ := os.Open(dir + name[:len(name)-3] + ".out")
			defer outFile.Close()

			outScanner := bufio.NewScanner(outFile)

			if outScanner.Scan() {
				parts := strings.Fields(outScanner.Text())
				want, _ = strconv.Atoi(parts[0])
			}
			tt[name] = testCase{x: x, y: y, M: M, want: want}
		}
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := getMaxFill(tc.x, tc.y, tc.M)

			if got != tc.want {
				t.Fatalf("TestCase failed for %s where got: %d and expected %d", name, got, tc.want)
			}
		})
	}

}
