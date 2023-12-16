package main

import (
	utils "ashmortar/advent-of-code/utilities"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/TwiN/go-color"
)

var year = "2023"
var day = "7"

// 32T3K 765
// T55J5 684
// KK677 28
// KTJJT 220
// QQQJA 483

type Hand struct {
	cards string
	bid   int
}

func parseHands(input string) []Hand {
	hands := []Hand{}

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		cards, strbid, found := strings.Cut(line, " ")
		if !found {
			fmt.Println("Error parsing line:", line)
			continue
		}
		bid, err := strconv.Atoi(strbid)
		if err != nil {
			fmt.Println("Error converting bid to int:", err)
			continue
		}

		hands = append(hands, Hand{
			cards: cards,
			bid:   bid,
		})
	}
	return hands
}

var cards = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}

func cardValue(card string) int {
	for i, c := range cards {
		if c == card {
			return i
		}
	}
	return -1
}

func compareHighCard(hand1, hand2 Hand) int {
	hand1Score := 0
	hand2Score := 0
	for i, card := range cards {
		if strings.Contains(hand1.cards, card) {
			hand1Score = i
		}
		if strings.Contains(hand2.cards, card) {
			hand2Score = i
		}
	}
	return hand1Score - hand2Score
}

// Every hand is exactly one type . From strongest to weakest, they are:

// Five of a kind , where all five cards have the same label: AAAAA
// Four of a kind , where four cards have the same label and one card has a different label: AA8AA
// Full house , where three cards have the same label, and the remaining two cards share a different label: 23332
// Three of a kind , where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
// Two pair , where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
// One pair , where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
// High card , where all cards' labels are distinct: 23456

func getHandType(hand Hand) (int, [][]int) {
	countMap := map[string]int{}
	for _, card := range strings.Split(hand.cards, "") {
		countMap[card]++
	}
	counts := [][]int{}
	for card, count := range countMap {
		counts = append(counts, []int{count, cardValue(card)})
	}
	slices.SortFunc(counts, func(count1, count2 []int) int {
		// num of cards first, then value of card for tie breaker
		if count1[0] == count2[0] {
			return count2[1] - count1[1]
		}
		return count2[0] - count1[0]
	})

	numUniqueCards := len(counts)

	switch numUniqueCards {
	case 1: // Five of a kind

		return 6, counts
	case 2: // Four of a kind or Full house
		if counts[0][0] == 4 {

			return 5, counts
		} else {

			return 4, counts
		}
	case 3: // Three of a kind or Two pair
		if counts[0][0] == 3 {

			return 3, counts
		} else {

			return 2, counts
		}
	case 4: // One pair

		return 1, counts
	case 5: // High card

		return 0, counts
	}
	fmt.Printf("Error getting hand type for hand: %v\n", hand)
	panic("Should not get here, illegal number of cards in hand")
}

func compareHands(hand1, hand2 Hand) int {
	handScore, cardScores := getHandType(hand1)
	hand2Score, cardScores2 := getHandType(hand2)

	if handScore == hand2Score {
		for i := 0; i < len(cardScores); i++ {
			hand1CardCount := cardScores[i][0]
			hand1CardScore := cardScores[i][1]
			hand2CardCount := cardScores2[i][0]
			hand2CardScore := cardScores2[i][1]
			if hand1CardCount == hand2CardCount {
				if hand1CardScore == hand2CardScore {
					continue
				}
				return hand1CardScore - hand2CardScore
			}
			return hand1CardCount - hand2CardCount
		}
	}
	return handScore - hand2Score
}

func Part1(input string) int {
	output := 0
	hands := parseHands(input)
	slices.SortFunc(hands, func(hand1, hand2 Hand) int {
		return compareHands(hand1, hand2)
	})
	for i, hand := range hands {
		ourHand := ""
		_, cardScores := getHandType(hand)
		for _, cardScore := range cardScores {
			ourHand += strings.Repeat(cards[cardScore[1]], cardScore[0])
		}
		fmt.Printf("%v %d\n", ourHand, hand.bid)
		score := hand.bid * (i + 1)
		// fmt.Printf("bid %d * rank %d = score %d\n\n", hand.bid, i+1, score)
		output += score
	}

	fmt.Printf("Output: %v\n", output)
	return output
}

func Part2(input string) int {
	output := 0
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
