package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// assumes ascii
func encodeText(text rune) int {
	if unicode.IsUpper(rune(text)) {
		return int(text - 38)
	} else {
		return int(text - 96)
	}
}

func splitString(text string) (string, string) {
	midway := (len(text) / 2)
	return text[0:midway], text[midway:]
}

func firstMatch(text1 string, text2 string) rune {
	charMapping := make(map[rune]bool)
	for _, character := range text1 {
		charMapping[character] = true
	}
	for _, character := range text2 {
		if charMapping[character] {
			return character
		}
	}
	panic(errors.New("no matching character"))
}

// We expand to all three match for this easily, but if we're expecting to keep growing in partitions,then we would need to re-consider the structure of this
func allThreeMatch(text1 string, text2 string, text3 string) rune {
	charMapping := make(map[rune]bool)
	firstTwoMatches := make(map[rune]bool)
	for _, character := range text1 {
		charMapping[character] = true
	}
	for _, character := range text2 {
		if charMapping[character] {
			firstTwoMatches[character] = true
		}
	}
	for _, character := range text3 {
		if firstTwoMatches[character] {
			return character
		}
	}
	panic(errors.New("no matching character"))
}

func part1(file *os.File) {
	reader := bufio.NewScanner(file)
	total := 0
	for reader.Scan() {
		first, second := splitString(reader.Text())
		match := firstMatch(first, second)
		priority := encodeText(match)
		total += priority
	}
	fmt.Println(total)
}

func part2(file *os.File) {
	reader := bufio.NewScanner(file)
	total := 0
	positionInLine := 0
	//If we were to expect growth on amount of lines parsed I wouldn't do it this way, using modulus would be much better
	var first string
	var second string
	var third string
	for reader.Scan() {
		switch positionInLine {
		case 0:
			first = reader.Text()
			positionInLine++
		case 1:
			second = reader.Text()
			positionInLine++
		case 2:
			third = reader.Text()
			match := allThreeMatch(first, second, third)
			priority := encodeText(match)
			total += priority
			positionInLine = 0
		}
	}
	fmt.Println(total)
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	//part1(file)
	part2(file)
}
