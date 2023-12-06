package main

import (
	utils "ashmortar/advent-of-code/utilities"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/TwiN/go-color"
)

// example line
// Game 3: 12 red, 1 blue; 6 red, 2 green, 3 blue; 2 blue, 5 red, 3 green

// sum each color for the game and if
//   - red > 12
//   - green > 13
//   - blue > 14
// then add the games id to the output

var colorMap = map[string]int{
	"red":   0,
	"green": 0,
	"blue":  0,
}
var colorLimit = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func part1(input []string) int {
	output := 0
	for _, line := range input {
		if line == "" {
			continue
		}
		// split the line into the game id and the colors
		game := strings.Split(line, ": ")
		id := game[0]
		colors := strings.Replace(game[1], "; ", ", ", -1)
		// reset the color map
		colorMap["red"] = 0
		colorMap["green"] = 0
		colorMap["blue"] = 0
		// loop through the colors and add them to the color map
		impossible := false
		colorSplit := strings.Split(colors, ", ")
		for _, color := range colorSplit {
			colorCount := strings.Split(color, " ")
			colorMap[colorCount[1]], _ = strconv.Atoi(colorCount[0])
			if colorMap[colorCount[1]] > colorLimit[colorCount[1]] {
				impossible = true
				break
			}
		}
		if !impossible {
			gameAndId := strings.Split(id, " ")
			idAsNumber, err := strconv.Atoi(gameAndId[1])
			if err != nil {
				panic(err)
			}
			output += idAsNumber
		}
	}
	return output
}

// example line
// Game 3: 12 red, 1 blue; 6 red, 2 green, 3 blue; 2 blue, 5 red, 3 green

// for each line find the highest count of each color
// multiply the highest count of each color together
// add the result to the output
func part2(input []string) int {
	output := 0

	for _, line := range input {
		if line == "" {
			continue
		}
		// split the line into the game id and the colors
		game := strings.Split(line, ": ")
		colors := strings.Replace(game[1], "; ", ", ", -1)
		// reset the color map
		colorMap["red"] = 0
		colorMap["green"] = 0
		colorMap["blue"] = 0
		// loop through the colors and add them to the color map
		colorSplit := strings.Split(colors, ", ")
		for _, colorAndCount := range colorSplit {
			slice := strings.Split(colorAndCount, " ")
			count, _ := strconv.Atoi(slice[0])
			if count > colorMap[slice[1]] {
				colorMap[slice[1]] = count
			}
		}
		// multiply the highest count of each color together
		output += colorMap["red"] * colorMap["green"] * colorMap["blue"]
	}
	return output
}

func main() {
	startTime := time.Now()
	input := utils.GetInputArray(2023, 2)
	postFetchTime := time.Now()
	result1 := part1(input)
	endTime1 := time.Now()
	result2 := part2(input)
	endTime2 := time.Now()

	printWidth := 25
	title := "\nAdvent of Code 2023 Day 2"

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
