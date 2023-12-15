package main

import (
	utils "ashmortar/advent-of-code/utilities"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/TwiN/go-color"
)

var year = "2023"
var day = "5"

// example input:
// seeds: 79 14 55 13

// seed-to-soil map:
// 50 98 2
// 52 50 48

// soil-to-fertilizer map:
// 0 15 37
// 37 52 2
// 39 0 15

// fertilizer-to-water map:
// 49 53 8
// 0 11 42
// 42 0 7
// 57 7 4

// water-to-light map:
// 88 18 7
// 18 25 70

// light-to-temperature map:
// 45 77 23
// 81 45 19
// 68 64 13

// temperature-to-humidity map:
// 0 69 1
// 1 0 69

// humidity-to-location map:
// 60 56 37
// 56 93 4

// return the lowest location number
// map rows  [x y z]
// where:
// - y is the source range start
// - x is the target range start
// - z is the range length

func getNumbers(input string) []int {
	numberRegex := regexp.MustCompile(`\d+`)
	numbers := numberRegex.FindAllString(input, -1)
	output := make([]int, len(numbers))
	for i, number := range numbers {
		output[i], _ = strconv.Atoi(number)
	}
	return output
}

func getSeeds(input string, part2 bool) [][]int {
	lines := strings.Split(input, "\n")
	seedString, found := strings.CutPrefix(lines[0], "seeds: ")
	if !found {
		fmt.Println("Error parsing input")
		return nil
	}
	result := [][]int{}
	numbers := getNumbers(seedString)
	for i, seed := range numbers {
		if part2 {
			if i%2 == 0 {
				result = append(result, []int{seed, numbers[i+1]})
			} else {
				continue
			}
		} else {
			result = append(result, []int{seed, 0})
		}
	}
	return result
}

func parseMapRow(input string) (sourceRangeStart, targetRangeStart, rangeLength int) {
	numbers := getNumbers(input)
	sourceRangeStart = numbers[1]
	targetRangeStart = numbers[0]
	rangeLength = numbers[2]
	return sourceRangeStart, targetRangeStart, rangeLength
}

func mapSeeds(options [][]int, input string) int {

	// fmt.Printf("Seeds are: %v\n", options)
	mappedThisRound := make([]bool, len(options))
	lines := strings.Split(input, "\n")
	// begin at the first line of the first map
	for i := 2; i < len(lines); i++ {
		// if this is an empty line continue
		if len(lines[i]) == 0 {
			fmt.Printf("results: %v\n\n", options)
			continue
		}
		// if this is the line identifying the map
		// then print the map name and continue
		if strings.Contains(lines[i], "map") {
			mappedThisRound = make([]bool, len(options))
			fmt.Printf("Processing %s\n", lines[i])
			continue
		}
		// we are in a map line

		sourceRangeStart, targetRangeStart, rangeLength := parseMapRow(lines[i])
		sourceRangeEnd := sourceRangeStart + rangeLength
		for j := 0; j < len(options); j++ {
			currentOption := options[j]
			currentStart := options[j][0]
			currentEnd := currentStart + options[j][1]
			alreadyMapped := mappedThisRound[j]
			if alreadyMapped {
				fmt.Printf("option %d has already been mapped\n", currentOption)
				continue
			}

			fmt.Printf("is %d in range %d-%d?\t", currentOption, sourceRangeStart, sourceRangeEnd)
			currentInRange := currentStart <= sourceRangeEnd && currentEnd >= sourceRangeStart
			if currentInRange && !mappedThisRound[j] {
				fmt.Printf("yes, ")

				offset := currentStart - sourceRangeStart
				fmt.Printf("offset from source range start is %d, ", offset)

				result := []int{targetRangeStart + offset, options[j][1]}
				fmt.Printf("result is %d + %d = %d, ", targetRangeStart, offset, result)
				options[j] = result
				wasMapped := true
				mappedThisRound[j] = wasMapped
				fmt.Printf("mapped to %d\n", result)
			} else {
				fmt.Printf("no, ")
				if mappedThisRound[j] {
					fmt.Printf("option has already been mapped\n")
				} else if currentEnd < sourceRangeStart {
					fmt.Printf("option is less than source range start\n")
				} else if currentStart >= sourceRangeEnd {
					fmt.Printf("option is greater than source range end\n")
				}
			}
		}
		fmt.Printf("mappedThisRound: %v\n", mappedThisRound)
		println()
	}
	fmt.Printf("\nresults: %v\n", options)

	slices.SortFunc(options, func(a, b []int) int {
		return a[0] - b[0]
	})
	fmt.Printf("sorted options: %v\n", options)
	return options[0][0]

}

func Part1(input string) int {
	// return mapSeeds(getSeeds(input, false), input)
	return 0
}

func Part2(input string) int {

	return mapSeeds(getSeeds(input, true), input)
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
