package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func getCommands(filename string) []Command {
	file, err := os.Open(filename)
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
	return commands
}

func TestGetTime(t *testing.T) {
	type testCase struct {
		commands []Command
		want     int
	}

	tt := make(map[string]testCase)

	dir := "./testdata/"

	entries, _ := os.ReadDir(dir)

	for _, e := range entries {
		name := e.Name()
		if strings.HasSuffix(name, ".in") {
			commands := getCommands(dir + name)
			file, _ := os.Open(dir + name[:len(name)-3] + ".out")

			scanner := bufio.NewScanner(file)

			res := -1

			if scanner.Scan() {
				res, _ = strconv.Atoi(scanner.Text())
			}

			tt[e.Name()] = testCase{commands: commands, want: res}
		}
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := getTime(tc.commands)

			if got != tc.want {
				t.Fatalf("TestCase %s failed with got: %d and want %d", name, got, tc.want)
			}
		})
	}

}
