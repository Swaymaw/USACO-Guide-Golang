package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getData(filename string) ([]int, int) {
	var N, K int

	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		parts := strings.Fields(scanner.Text())

		N, _ = strconv.Atoi(parts[0])
		K, _ = strconv.Atoi(parts[1])
	}

	diamonds := make([]int, 0, N)

	for range N {
		if scanner.Scan() {
			parts := strings.Fields(scanner.Text())
			val, _ := strconv.Atoi(parts[0])

			diamonds = append(diamonds, val)
		}
	}
	return diamonds, K
}

func getMaxRange(diamonds []int, K int) int {
	l := 0
	r := 1
	maxDiamonds := 1

	sort.Ints(diamonds)

	for l < r && r < len(diamonds) {
		if diamonds[r]-diamonds[l] <= K {
			maxDiamonds = max(maxDiamonds, r-l+1)
		} else {
			for l < r && diamonds[r]-diamonds[l] > K {
				l++
			}
		}
		r++
	}

	return maxDiamonds
}

func main() {
	diamonds, K := getData("testdata/2.in")

	res := getMaxRange(diamonds, K)

	fmt.Println(res)

}
