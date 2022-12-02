package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type result struct {
	resultsInLoss int
	resultsInWin  int
}

var selection map[string]int = map[string]int{"X": 1, "Y": 2, "Z": 3, "A": 1, "B": 2, "C": 3}

var necessaryInput map[int]result = map[int]result{1: {3, 2}, 2: {1, 3}, 3: {2, 1}}

func matchResult(theirInput string, necessaryOutcome string) int {
	theirSelection := selection[theirInput]
	switch necessaryOutcome {
	case "X":
		return necessaryInput[theirSelection].resultsInLoss
	case "Y":
		return theirSelection + 3
	case "Z":
		return necessaryInput[theirSelection].resultsInWin + 6
	default:
		panic(errors.New("invalid outcome string encountered"))
	}
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	var score = 0
	var reader = bufio.NewScanner(file)
	for reader.Scan() {
		unparsed := reader.Text()
		parsed := strings.Fields(unparsed)
		score += matchResult(parsed[0], parsed[1])
	}

	fmt.Println(score)
}
