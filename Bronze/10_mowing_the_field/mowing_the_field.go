package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	direction string
	distance  int
}

func getTime(commands []Command) int {
	lastVisit := [2001][2001]int{}

	for i := range lastVisit {
		for j := range lastVisit[i] {
			lastVisit[i][j] = -1
		}
	}

	minTime := math.MaxInt

	curx := 1000
	cury := 1000
	curTime := 0

	lastVisit[curx][cury] = 0

	for _, command := range commands {

		for range command.distance {
			curTime += 1

			if command.direction == "N" {
				cury += 1
			} else if command.direction == "E" {
				curx += 1
			} else if command.direction == "W" {
				curx -= 1
			} else if command.direction == "S" {
				cury -= 1
			}

			if lastVisit[curx][cury] != -1 {
				minTime = min(minTime, curTime-lastVisit[curx][cury])
			}

			lastVisit[curx][cury] = curTime
		}
	}

	if minTime == math.MaxInt {
		return -1
	}

	return minTime
}

func main() {
	file, err := os.Open("testdata/2.in")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	N := 0
	if scanner.Scan() {
		N, _ = strconv.Atoi(scanner.Text())
	}

	commands := make([]Command, 0, N)

	for range N {
		if scanner.Scan() {
			parts := strings.Fields(scanner.Text())
			distance, _ := strconv.Atoi(parts[1])

			commands = append(commands, Command{direction: parts[0], distance: distance})
		}
	}

	res := getTime(commands)

	fmt.Println(res)
}
