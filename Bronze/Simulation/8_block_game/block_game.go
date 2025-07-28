package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Block struct {
	front string
	back  string
}

func getNumWords(blocks []Block) map[rune]int {
	// get total words needed
	count := map[rune]int{}

	for _, val := range blocks {
		for _, char := range val.front {
			count[char] += 1
		}
		for _, char := range val.back {
			count[char] += 1
		}
	}

	// remove words from each block which are common in back and front
	for _, val := range blocks {
		count1 := map[rune]int{}
		count2 := map[rune]int{}

		for _, char := range val.front {
			count1[char] += 1
		}
		for _, char := range val.back {
			count2[char] += 1
		}

		common := map[rune]int{}

		for r, c1 := range count1 {
			if c1 >= 1 && count2[r] >= 1 {
				common[r] = min(c1, count2[r])
			}
		}

		for r, c := range common {
			count[r] -= c
		}
	}
	return count
}

func main() {
	var n int
	blocks := make([]Block, 0, n)

	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)
	for range n {
		if scanner.Scan() {
			line := scanner.Text()
			parts := strings.Fields(line)
			blocks = append(blocks, Block{front: parts[0], back: parts[1]})
		}
	}

	res := getNumWords(blocks)

	for _, val := range "abcdefghijklmnopqrstuvwxyz" {
		fmt.Printf("%d\n", res[val])
	}

}
