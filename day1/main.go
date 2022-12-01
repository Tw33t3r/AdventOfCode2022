package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type ElfAmount struct {
	elfId  int
	amount int
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	total := 0
	var amounts []ElfAmount
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		text := reader.Text()
		if text == "" {
			amounts = append(amounts, ElfAmount{elfId: len(amounts), amount: total})
			total = 0
		} else {
			amount, err := strconv.Atoi(text)
			check(err)
			total += amount
		}
	}

	sort.Slice(amounts, func(i int, j int) bool {
		return amounts[i].amount < amounts[j].amount
	})

	for _, value := range amounts {
		fmt.Println(value)
	}
}
