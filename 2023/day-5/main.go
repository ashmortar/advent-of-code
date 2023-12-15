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

type MapRange struct {
	name   string
	ranges [][]int
}

func getMapName(input string, index int) string {
	mapNamePattern := regexp.MustCompile(`\w+-to-\w+`)
	mapNames := mapNamePattern.FindAllString(input, -1)
	return mapNames[index]
}

func getMaps(input string) []MapRange {
	orderedMapRanges := []MapRange{}
	mapPattern := regexp.MustCompile(`map:\n`)
	mapRowPattern := regexp.MustCompile(`\d+ \d+ \d+`)
	mapIndices := mapPattern.FindAllStringIndex(input, -1)
	for i, index := range mapIndices {
		mapRange := MapRange{
			name:   getMapName(input, i),
			ranges: [][]int{},
		}

		start := index[1]
		var end int
		if i == len(mapIndices)-1 {
			end = len(input)
		} else {
			end = mapIndices[i+1][0]
		}
		mapString := input[start:end]
		mapLines := strings.Split(mapString, "\n")
		for _, line := range mapLines {
			isMapRow := mapRowPattern.MatchString(line)
			if !isMapRow {
				continue
			}
			sourceRangeStart, targetRangeStart, rangeLength := parseMapRow(line)
			mapRange.ranges = append(mapRange.ranges, []int{sourceRangeStart, targetRangeStart, rangeLength})
		}
		orderedMapRanges = append(orderedMapRanges, mapRange)
	}
	return orderedMapRanges
}

func processMap(inputRanges [][]int, mapRangeStruct MapRange, noisy bool) [][]int {
	if noisy {
		fmt.Printf("\n\n~~~~~~~~~Processing map %s~~~~~~~~~~\n", mapRangeStruct.name)
	}
	results := [][]int{}
	inputsToProcess := [][]int{}
	for _, inputRange := range inputRanges {
		inputsToProcess = append(inputsToProcess, []int{inputRange[0], inputRange[1]})
	}

	for len(inputsToProcess) > 0 {
		inputRange := inputsToProcess[0]
		inputsToProcess = inputsToProcess[1:]
		if noisy {
			time.Sleep(1 * time.Second)
			fmt.Printf("checking input range: %d-%d\n", inputRange[0], inputRange[0]+inputRange[1])
		}

		mapped := false
		inputRangeStart := inputRange[0]
		inputRangeEnd := inputRangeStart + inputRange[1]

		for _, mapRange := range mapRangeStruct.ranges {
			if noisy {
				fmt.Printf("\tagainst map range: %d-%d\n", mapRange[0], mapRange[0]+mapRange[2])
			}
			if mapped {
				if noisy {
					fmt.Printf("\t\talready mapped\n")
				}
				break
			}
			sourceRangeStart := mapRange[0]
			sourceRangeEnd := sourceRangeStart + mapRange[2]
			targetRangeStart := mapRange[1]

			rangesOverlap := inputRangeStart >= sourceRangeStart && inputRangeStart <= sourceRangeEnd ||
				inputRangeEnd >= sourceRangeStart && inputRangeEnd <= sourceRangeEnd

			if rangesOverlap {
				if noisy {
					fmt.Print("\tYES\n")
				}
				startOfOverLap := inputRangeStart
				if startOfOverLap < sourceRangeStart {
					if noisy {
						fmt.Printf("\t\tstart of overlap is less than target range start\n")
					}
					startOfOverLap = sourceRangeStart
				}
				endOfOverLap := inputRangeEnd
				if endOfOverLap > sourceRangeEnd {
					if noisy {
						fmt.Printf("\t\tend of overlap is greater than target range end\n")
					}
					endOfOverLap = sourceRangeEnd
				}
				if noisy {
					fmt.Printf("\t\toverlap: %d-%d\n", startOfOverLap, endOfOverLap)
				}
				offset := startOfOverLap - sourceRangeStart
				if noisy {
					fmt.Printf("\t\toffset: %d\n", offset)
				}
				mappedResult := []int{targetRangeStart + offset, endOfOverLap - startOfOverLap}
				if noisy {
					fmt.Printf("\t\tmapped result: %v\n", mappedResult)
				}

				results = append(results, mappedResult)

				hasLeftOverStart := startOfOverLap > inputRangeStart
				if hasLeftOverStart {
					leftOverStart := []int{inputRangeStart, startOfOverLap - inputRangeStart - 1}
					if noisy {
						fmt.Printf("\t\tleft over start: %v\n", leftOverStart)
					}
					inputsToProcess = append(inputsToProcess, leftOverStart)
				}
				hasLeftOverEnd := endOfOverLap < inputRangeEnd
				if hasLeftOverEnd {
					leftOverEnd := []int{endOfOverLap + 1, inputRangeEnd - endOfOverLap - 1}
					if noisy {
						fmt.Printf("\t\tleft over end: %v\n", leftOverEnd)
					}
					inputsToProcess = append(inputsToProcess, leftOverEnd)
				}
				mapped = true
				if noisy {
					fmt.Println()
				}
				break
			}
			if noisy {
				fmt.Printf("\tNO\n\n")
			}
		}
		if !mapped {
			if noisy {
				fmt.Printf("\tNo map, range unchanged\n")
			}

			results = append(results, inputRange)
		}
	}
	return results
}

func mapSeeds(options [][]int, input string, noisy bool) int {
	if noisy {
		fmt.Printf("Seeds are: %v\n", options)
	}
	maps := getMaps(input)
	if noisy {
		fmt.Printf("Maps are: %v\n", maps)
	}
	ourOpts := [][]int{}
	for _, opt := range options {
		ourOpts = append(ourOpts, []int{opt[0], opt[1]})
	}
	for _, mapRanges := range maps {

		ourOpts = processMap(ourOpts, mapRanges, noisy)
		if noisy {
			fmt.Printf("options after map: %v\n", ourOpts)
		}
	}

	slices.SortFunc(ourOpts, func(a, b []int) int {
		return a[0] - b[0]
	})
	if noisy {
		fmt.Printf("sorted options: %v\n", ourOpts)
	}
	return ourOpts[0][0]
}

func Part1(input string) int {
	return mapSeeds(getSeeds(input, false), input, false)
}

func Part2(input string) int {
	return mapSeeds(getSeeds(input, true), input, false)
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
