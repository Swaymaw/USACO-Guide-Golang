package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Segment struct {
	length int
	limit  int
}

func maxOverSpeed(track_segments, cow_segments []Segment) int {
	ans := 0

	var i, j int

	var track_length, cow_length int

	for i < len(track_segments) && j < len(cow_segments) {
		ans = max(cow_segments[j].limit-track_segments[i].limit, ans)
		if cow_length+cow_segments[j].length < track_length+track_segments[i].length {
			cow_length += cow_segments[j].length
			j++
		} else if cow_length+cow_segments[j].length > track_length+track_segments[i].length {
			track_length += track_segments[i].length
			i++
		} else {
			cow_length += cow_segments[j].length
			track_length += track_segments[i].length
			i++
			j++
		}

	}
	return ans
}

func main() {
	var n, m int
	var track_segments, cow_segments []Segment

	fmt.Scan(&n, &m)
	scanner := bufio.NewScanner(os.Stdin)

	for range n {
		if scanner.Scan() {
			line := scanner.Text()
			parts := strings.Fields(line)

			length, _ := strconv.Atoi(parts[0])
			limit, _ := strconv.Atoi(parts[1])

			track_segments = append(track_segments, Segment{length: length, limit: limit})
		}
	}
	for range m {
		if scanner.Scan() {
			line := scanner.Text()
			parts := strings.Fields(line)

			length, _ := strconv.Atoi(parts[0])
			limit, _ := strconv.Atoi(parts[1])

			cow_segments = append(cow_segments, Segment{length: length, limit: limit})
		}
	}

	fmt.Println("Track Segments")
	fmt.Println(track_segments)
	fmt.Println("Cow Segments")
	fmt.Println(cow_segments)
	ans := maxOverSpeed(track_segments, cow_segments)
	fmt.Println(ans)
}
