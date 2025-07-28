package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMaxFill(x, y, M int) int {
	maxFill := 0

	possible := M / y

	if possible == 0 {
		return (M / x) * x
	}

	for mul := range possible + 1 {
		val := y * mul
		val = val + (((M - val) / x) * x)
		maxFill = max(val, maxFill)
	}

	return maxFill
}

func main() {
	var x, y, M int
	file, _ := os.Open("testdata/4.in")
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		x, _ = strconv.Atoi(parts[0])
		y, _ = strconv.Atoi(parts[1])
		M, _ = strconv.Atoi(parts[2])
	}

	res := getMaxFill(x, y, M)

	fmt.Println(res)

}
