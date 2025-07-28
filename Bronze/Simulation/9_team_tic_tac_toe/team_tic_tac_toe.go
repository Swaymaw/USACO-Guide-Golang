package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func getPossibleWins(board [3][3]rune) (int, int) {
	one_player := 0
	two_player := 0

	teams_won := map[string]struct{}{}

	// check win in rows
	for idx := range 3 {
		account := map[rune]struct{}{}

		for _, val := range board[idx] {
			account[val] = struct{}{}
		}
		string_account := ""

		for key, _ := range account {
			string_account += string(key)
		}

		splitted := strings.Split(string_account, "")
		sort.Strings(splitted)
		string_account = strings.Join(splitted, "")

		_, exists := teams_won[string_account]

		if len(account) == 1 && !exists {
			one_player += 1
			teams_won[string_account] = struct{}{}
		} else if len(account) == 2 && !exists {
			two_player += 1
			teams_won[string_account] = struct{}{}
		}
	}
	// check win in cols
	for idx := range 3 {
		account := map[rune]struct{}{}

		for iter := range 3 {
			account[board[iter][idx]] = struct{}{}
		}
		string_account := ""

		for key, _ := range account {
			string_account += string(key)
		}

		splitted := strings.Split(string_account, "")
		sort.Strings(splitted)
		string_account = strings.Join(splitted, "")

		_, exists := teams_won[string_account]

		if len(account) == 1 && !exists {
			one_player += 1
			teams_won[string_account] = struct{}{}
		} else if len(account) == 2 && !exists {
			two_player += 1
			teams_won[string_account] = struct{}{}
		}
	}
	// check win in diagnols
	account := map[rune]struct{}{}
	account[board[0][0]] = struct{}{}
	account[board[1][1]] = struct{}{}
	account[board[2][2]] = struct{}{}

	string_account := ""

	for key, _ := range account {
		string_account += string(key)
	}
	splitted := strings.Split(string_account, "")
	sort.Strings(splitted)
	string_account = strings.Join(splitted, "")

	_, exists := teams_won[string_account]

	if len(account) == 1 && !exists {
		one_player += 1
		teams_won[string_account] = struct{}{}
	} else if len(account) == 2 && !exists {
		two_player += 1
		teams_won[string_account] = struct{}{}
	}

	account = map[rune]struct{}{}
	account[board[2][0]] = struct{}{}
	account[board[1][1]] = struct{}{}
	account[board[0][2]] = struct{}{}

	string_account = ""

	for key, _ := range account {
		string_account += string(key)
	}
	splitted = strings.Split(string_account, "")
	sort.Strings(splitted)
	string_account = strings.Join(splitted, "")

	_, exists = teams_won[string_account]

	if len(account) == 1 && !exists {
		one_player += 1
		teams_won[string_account] = struct{}{}
	} else if len(account) == 2 && !exists {
		two_player += 1
		teams_won[string_account] = struct{}{}
	}

	return one_player, two_player
}

func main() {
	board := [3][3]rune{}
	scanner := bufio.NewScanner(os.Stdin)
	for row := range 3 {
		if scanner.Scan() {
			line := scanner.Text()
			for col, char := range line {
				board[row][col] = char
			}
		}
	}

	one_p, two_p := getPossibleWins(board)
	fmt.Println(one_p)
	fmt.Println(two_p)
}
