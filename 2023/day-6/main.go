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
var day = "6"

// Time:      7  15   30
// Distance:  9  40  200

// the above is 3 race records
// 9mm in 7ms
// 40mm in 15ms
// 200mm in 30ms

// our boat starts at
// 0mm
// 0ms
// 0mm/ms

// for each ms holding go increase speed
// by 1mm/ms for the rest of the ms of the race
// but don't move while holding go

//              7ms race
// throttle time    |   travel time      |   distance
// 0ms              |   7ms - 0ms = 7ms  |   0mm/ms * 7ms = 0mm
// 1ms              |   7ms - 1ms = 6ms  |   1mm/ms * 6ms = 6mm
// 2ms              |   7ms - 2ms = 5ms  |   2mm/ms * 5ms = 10mm
// 3ms              |   7ms - 3ms = 4ms  |   3mm/ms * 4ms = 12mm
// 4ms              |   7ms - 4ms = 3ms  |   4mm/ms * 3ms = 12mm
// 5ms              |   7ms - 5ms = 2ms  |   5mm/ms * 2ms = 10mm
// 6ms              |   7ms - 6ms = 1ms  |   6mm/ms * 1ms = 6mm
// 7ms              |   7ms - 7ms = 0ms  |   7mm/ms * 0ms = 0mm

type Race struct {
	distance int
	time     int
}

func parseRaces(input string, part2 bool) []Race {
	races := []Race{}
	numberRegex := regexp.MustCompile(`\d+`)
	lines := strings.Split(input, "\n")
	timeStrings := numberRegex.FindAllString(lines[0], -1)
	distanceStrings := numberRegex.FindAllString(lines[1], -1)
	if part2 {
		time, err := strconv.Atoi(strings.Join(timeStrings, ""))
		if err != nil {
			panic(err)
		}
		distance, err := strconv.Atoi(strings.Join(distanceStrings, ""))
		if err != nil {
			panic(err)
		}
		race := Race{distance: distance, time: time}
		races = append(races, race)
		return races
	}
	for i := 0; i < len(timeStrings); i++ {
		time, err := strconv.Atoi(timeStrings[i])
		if err != nil {
			panic(err)
		}
		distance, err := strconv.Atoi(distanceStrings[i])
		if err != nil {
			panic(err)
		}
		races = append(races, Race{distance: distance, time: time})
	}
	return races
}

func findNumberOfWins(race Race) int {
	for throttleMs := 0; throttleMs < race.time/2; throttleMs++ {
		travelMs := race.time - throttleMs
		speed := throttleMs * 1 // 1mm/ms
		distance := speed * travelMs
		if distance > race.distance {
			// distance is symettrical around time
			losses := (throttleMs - 1) * 2
			wins := race.time - losses - 1
			return wins
		}
	}
	return 0
}

func Part1(input string) int {
	output := 1
	for _, race := range parseRaces(input, false) {
		output = output * findNumberOfWins(race)
	}
	return output
}

func Part2(input string) int {
	return findNumberOfWins(parseRaces(input, true)[0])
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
