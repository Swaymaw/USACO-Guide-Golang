package main

import (
	"fmt"
)

func getSteps(x, y int) int {
	pos := x
	delta := 1
	distance := 0
	older_pos := x
	for {
		older_pos = pos
		pos = x + delta
		if pos*pos > y*y && pos*y > 0 {
			if older_pos-y > 0 {
				distance += older_pos - y
			} else {
				distance -= older_pos - y
			}
			break
		}
		if older_pos-pos > 0 {
			distance += older_pos - pos
		} else {
			distance -= older_pos - pos
		}
		delta *= -2
	}
	return distance
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	res := getSteps(x, y)

	fmt.Println("Steps Taken:", res)
}
