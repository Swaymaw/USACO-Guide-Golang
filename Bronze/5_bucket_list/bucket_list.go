package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getMaxBuckets(buckets map[int]int) int {
	ans := 0
	current_buckets := 0

	keys := make([]int, 0, len(buckets))
	for k := range buckets {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, val := range keys {
		current_buckets += buckets[val]
		ans = max(ans, current_buckets)
	}
	return ans
}

func main() {
	var n int
	buckets := map[int]int{}

	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)
	for range n {
		if scanner.Scan() {
			line := scanner.Text()

			parts := strings.Fields(line)

			s, _ := strconv.Atoi(parts[0])
			t, _ := strconv.Atoi(parts[1])
			b, _ := strconv.Atoi(parts[2])

			buckets[s] = b
			buckets[t] = -b
		}
	}

	res := getMaxBuckets(buckets)

	fmt.Println("Max required buckets:", res)
}
