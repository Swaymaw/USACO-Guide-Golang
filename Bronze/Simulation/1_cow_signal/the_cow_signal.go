package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func expand(grid []string, k int) []string {
	expandedGrid := make([]string, len(grid)*k)
	idx := 0
	for _, row := range grid {
		for range k {
			expandedGrid[idx] = row
			idx++
		}
	}
	for idx, row := range expandedGrid {
		newString := ""

		for _, char := range row {
			newString += strings.Repeat(string(char), k)
		}
		expandedGrid[idx] = newString
	}

	return expandedGrid
}

func main() {
	var m, n, k int
	fmt.Scan(&m, &n, &k)

	scanner := bufio.NewScanner(os.Stdin)

	grid := make([]string, m)

	for i := 0; i < m; {
		if scanner.Scan() {
			line := scanner.Text()
			if line != "" {
				grid[i] = line
				i++
			}
		}
	}

	res := expand(grid, k)

	for _, row := range res {
		fmt.Println(row)
	}
}
