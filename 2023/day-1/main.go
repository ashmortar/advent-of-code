package main

import (
	utils "ashmortar/advent-of-code/utilities"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/TwiN/go-color"
)

func isDigit(char string) bool {
	return char == "0" || char == "1" || char == "2" || char == "3" || char == "4" || char == "5" || char == "6" || char == "7" || char == "8" || char == "9"
}

func part1(input []string) int {
	output := 0

	for _, line := range input {
		if line == "" {
			continue
		}
		firstDigit := ""
		lastDigit := ""
		loopLength := len(line)
		for i := 0; i < int(loopLength); i++ {
			if firstDigit == "" && isDigit(string(line[i])) {
				firstDigit = string(line[i])
			}
			if lastDigit == "" && isDigit(string(line[len(line)-i-1])) {
				lastDigit = string(line[len(line)-i-1])
			}
			if firstDigit != "" && lastDigit != "" {
				break
			}
		}

		if firstDigit == lastDigit && firstDigit == "" {
			panic("No digits found")
		}
		strNum := firstDigit + lastDigit
		number, err := strconv.Atoi(strNum)
		if err != nil {
			panic(err)
		}
		output += number
	}
	return output
}

var digitMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getDigitIfContainsAsString(tail string) int {
	for digit, num := range digitMap {
		if strings.Contains(tail, digit) {
			return num
		}
	}
	return 0
}

func part2(input []string) int {
	output := 0

	for _, line := range input {
		if line == "" {
			continue
		}
		firstDigit := ""
		frontTail := ""
		lastDigit := ""
		endTail := ""
		loopLength := len(line)
		for i := 0; i < int(loopLength); i++ {
			if firstDigit == "" {
				if isDigit(string(line[i])) {
					firstDigit = string(line[i])
				} else {
					frontTail += string(line[i])
					containedNumber := getDigitIfContainsAsString(frontTail)
					if containedNumber != 0 {
						firstDigit = strconv.Itoa(containedNumber)
					}
				}
			}
			if lastDigit == "" {
				if isDigit(string(line[len(line)-i-1])) {
					lastDigit = string(line[len(line)-i-1])
				} else {
					endTail = string(line[len(line)-i-1]) + endTail
					containedNumber := getDigitIfContainsAsString(endTail)
					if containedNumber != 0 {
						lastDigit = strconv.Itoa(containedNumber) + lastDigit
					}
				}
			}
			if firstDigit != "" && lastDigit != "" {
				break
			}
		}

		if firstDigit == lastDigit && firstDigit == "" {
			panic("No digits found")
		}
		strNum := firstDigit + lastDigit
		number, err := strconv.Atoi(strNum)
		if err != nil {
			panic(err)
		}
		output += number
	}
	return output
}

func main() {
	startTime := time.Now()
	input := utils.GetInput(2023, 1)
	postFetchTime := time.Now()
	result1 := part1(input)
	endTime1 := time.Now()
	result2 := part2(input)
	endTime2 := time.Now()

	printWidth := 25
	title := "\n\t\tAdvent of Code 2023 Day 1"
	// fmt.Println(title)
	// fmt.Println(strings.Repeat("=", 55))
	// fmt.Println("Input fetched in", postFetchTime.Sub(startTime))
	// fmt.Println("Part 1:")
	// fmt.Printf("Result: %-*d\tTime: %v\n", printWidth, result1, endTime1.Sub(postFetchTime))
	// fmt.Println("Part 2:")
	// fmt.Printf("Result: %-*d\tTime: %v\n", printWidth, result2, endTime2.Sub(endTime1))
	// fmt.Println("Total Time:", endTime2.Sub(startTime))

	fmt.Println(color.Bold + color.Purple + title + color.Reset)
	fmt.Println(color.Bold + color.Blue + strings.Repeat("=", 55) + color.Reset)
	fmt.Println(color.Bold + color.Blue + "Input:" + color.Reset)
	fmt.Printf(color.Bold+"Lines: %-*d\tParse: %v\n"+color.Reset, printWidth, len(input), postFetchTime.Sub(startTime))
	fmt.Println(color.Bold + color.Blue + "Part 1:" + color.Reset)
	fmt.Printf(color.Bold+"Result: %-*d\tTime: %v\n"+color.Reset, printWidth, result1, endTime1.Sub(postFetchTime))
	fmt.Println(color.Bold + color.Blue + "Part 2:" + color.Reset)
	fmt.Printf(color.Bold+"Result: %-*d\tTime: %v\n"+color.Reset, printWidth, result2, endTime2.Sub(endTime1))
	fmt.Println(color.Bold+color.Blue+"Total Time:"+color.Reset, endTime2.Sub(startTime))
	fmt.Print("\n")
}
