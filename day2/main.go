package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var selection map[string]int = map[string]int{"X": 1, "Y": 2, "Z": 3, "A": 1, "B": 2, "C": 3}

func result(theirInput string, ourInput string) int {
	ourSelection := selection[ourInput]
	theirSelection := selection[theirInput]
	total := ourSelection
	diff := theirSelection - ourSelection
	//This method is similar to a DFA, if the number of possibilities expands this isn't tenable
	if diff == 0 {
		total += 3
	} else if diff == 1 {
	} else if diff == -1 {
		total += 6
	} else if diff == 2 {
		total += 6
	} else if diff == -2 {
	}
	return total
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
		score += result(parsed[0], parsed[1])
	}

	fmt.Println(score)
}
