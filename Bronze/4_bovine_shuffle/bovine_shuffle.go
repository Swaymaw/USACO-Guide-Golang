package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getOriginalPos(shuffle, ids []int) []int {
	for idx := range shuffle {
		shuffle[idx] -= 1
	}

	original_mapping := make([]int, len(shuffle))

	for range 3 {
		for key, val := range shuffle {
			original_mapping[val] = key
		}
		_ = copy(shuffle, original_mapping)
	}

	new_ids := make([]int, len(ids))
	for idx, val := range original_mapping {
		new_ids[idx] = ids[val]
	}

	return new_ids
}

func main() {
	var n int
	var shuffle []int
	var ids []int

	fmt.Scan(&n)
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		text := scanner.Text()
		parts := strings.Fields(text)

		for _, part := range parts {
			val_int, _ := strconv.Atoi(string(part))
			shuffle = append(shuffle, val_int)
		}
	}
	if scanner.Scan() {
		text := scanner.Text()
		parts := strings.Fields(text)

		for _, part := range parts {
			val_int, _ := strconv.Atoi(string(part))
			ids = append(ids, val_int)
		}

	}

	originalPos := getOriginalPos(shuffle, ids)

	for _, id := range originalPos {
		fmt.Println(id)
	}
}
