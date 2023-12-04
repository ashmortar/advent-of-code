package main

import (
	utils "ashmortar/advent-of-code/utilities"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/TwiN/go-color"
)

var year = "2023"
var day = "4"

// example input:
// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
// Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
// Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
// Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
// Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
// Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11

func printLineResult(line string, matchedNumbers []string, value int) {
	fmt.Println(line)
	left, right, _ := strings.Cut(line, "|")
	for _, number := range matchedNumbers {
		withSpaces := " " + number + " "
		leftIndex := strings.Index(left, withSpaces)
		rightIndex := strings.Index(right, withSpaces)
		length := len(number)
		leftSpace := strings.Repeat(" ", leftIndex+1)
		betweenSpace := strings.Repeat(" ", len(left)-leftIndex-length+rightIndex+1)
		fmt.Println(leftSpace + color.Green + number + color.Reset + betweenSpace + color.Green + number + color.Reset)
	}
	fmt.Printf("Matched: %d\tValue: ", len(matchedNumbers))
	fmt.Print(color.Red + strconv.Itoa(value) + color.Reset)
	fmt.Println()
	// time.Sleep(1 * time.Second)
}

func getGroups(line string) (string, []string, []string) {
	numbersRegex := regexp.MustCompile(`\d+`)
	name, values, found := strings.Cut(line, ":")
	if !found {
		panic("Error parsing card:" + line)
	}
	left, right, found := strings.Cut(values, "|")
	if !found {
		panic("Error parsing card:" + line)
	}
	leftNumbers := numbersRegex.FindAllString(left, -1)
	rightNumbers := numbersRegex.FindAllString(right, -1)
	return name, leftNumbers, rightNumbers
}

func findMatches(leftNumbers []string, rightNumbers []string) []string {
	matches := []string{}
	for _, leftNumber := range leftNumbers {
		for _, rightNumber := range rightNumbers {
			if leftNumber == rightNumber {
				matches = append(matches, leftNumber)
			}
		}
	}
	return matches
}

func scoreCard(matches []string) int {
	value := 0
	if len(matches) > 0 {
		value = 1 << (len(matches) - 1)
	}
	return value
}

func getCards(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

// in each line add to the output value of the card
// a cards value is equal to 2 to the power of the number of
// matches -1.  A match is when a number on the left side
// of the | is also on the right side of the |
func Part1(input string) int {
	output := 0
	cards := getCards(input)

	for _, card := range cards {
		_, leftNumbers, rightNumbers := getGroups(card)
		matches := findMatches(leftNumbers, rightNumbers)
		value := scoreCard(matches)
		output += value
		// printLineResult(card, matches, value)
	}
	return output
}

// for part 2 we win copies of the cards below the winning
// card equal to the number of matches in that card
// so if card 10 has 5 matches we get one additional copy
// of cards 11-15.  This process continues until none of
// the copies cause you to win any more cards.
func Part2(input string) int {
	cards := getCards(input)
	cardCounts := make([]int, len(cards))
	output := 0
	for i, card := range cards {
		_, leftNumbers, rightNumbers := getGroups(card)
		// increment the count of this card to count as the first copy
		cardCounts[i]++
		matches := findMatches(leftNumbers, rightNumbers)
		// fmt.Printf("Card %d: %v\n", i+1, matches)
		// increment the counts of the next n cards
		for j := 1; j < len(matches)+1; j++ {
			cardCounts[i+j] += cardCounts[i]
		}
		output += cardCounts[i]
	}
	return output

}

func main() {
	year, err := strconv.Atoi(year)
	if err != nil {
		fmt.Println("Error converting year to int:", err)
		return
	}
	day, err := strconv.Atoi(day)
	if err != nil {
		fmt.Println("Error converting day to int:", err)
		return
	}
	startTime := time.Now()
	input := utils.GetInputString(year, day)
	puzzleParseTime := time.Now()
	result1 := Part1(input)
	part1Time := time.Now()
	result2 := Part2(input)
	part2Time := time.Now()

	printWidth := 25
	title := fmt.Sprintf("Advent of Code %d - Day %d", year, day)

	fmt.Println(color.Bold + color.Purple + title + color.Reset)
	fmt.Println(color.Bold + color.Blue + strings.Repeat("=", 55) + color.Reset)
	fmt.Println(color.Bold + color.Blue + "Input:" + color.Reset)
	fmt.Printf(color.Bold+"Lines: %-*d\tParse: %v\n"+color.Reset, printWidth, len(input), puzzleParseTime.Sub(startTime))
	fmt.Println(color.Bold + color.Blue + "Part 1:" + color.Reset)
	fmt.Printf(color.Bold+"Result: %-*d\tTime: %v\n"+color.Reset, printWidth, result1, part1Time.Sub(puzzleParseTime))
	fmt.Println(color.Bold + color.Blue + "Part 2:" + color.Reset)
	fmt.Printf(color.Bold+"Result: %-*d\tTime: %v\n"+color.Reset, printWidth, result2, part2Time.Sub(part1Time))
	fmt.Println(color.Bold+color.Blue+"Total Time:"+color.Reset, part2Time.Sub(startTime))
	fmt.Print("\n")
}
