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

// Here is an example engine schematic:
// ```
// 467..114..
// ...*......
// ..35..633.
// ......#...
// 617*......
// .....+.58.
// ..592.....
// ......755.
// ...$.*....
// .664.598..
// ```
// In this schematic, two numbers are not part numbers because they are not adjacent to a symbol: 114 (top right) and 58 (middle right). Every other number is adjacent to a symbol and so is a part number; their sum is 4361.

func isDigit(char string) bool {
	re := regexp.MustCompile(`[0-9]`)
	return re.MatchString(char)
}

func isNotDigitOrPeriod(char string) bool {
	re := regexp.MustCompile(`[^0-9.]`)
	return re.MatchString(char)
}

func isPartNumber(number string, currentIndex int, puzzleIndex int, puzzle []string) bool {
	leftIndex := currentIndex - len(number) - 1
	topPuzzleIndex := puzzleIndex - 1
	bottomPuzzleIndex := puzzleIndex + 1
	passed := false
	for i := leftIndex; i <= currentIndex; i++ {
		if passed {
			break
		}
		for j := topPuzzleIndex; j <= bottomPuzzleIndex; j++ {
			if passed {
				break
			}
			if i >= 0 && j >= 0 && j < len(puzzle) && i < len(puzzle[j]) {
				result := isNotDigitOrPeriod(string(puzzle[j][i]))
				if result {
					// fmt.Printf("i %d j %d char %s result %t\n", i, j, string(puzzle[j][i]), result)
					passed = result
				}
			}
		}

	}
	return passed
}

func printNeighborhood(puzzle []string, currentIndex int, puzzleIndex int, length int) {
	leftIndex := currentIndex - 1 - length
	topPuzzleIndex := puzzleIndex - 1
	bottomPuzzleIndex := puzzleIndex + 1
	str := ""
	for j := topPuzzleIndex; j <= bottomPuzzleIndex; j++ {
		if j >= 0 && j < len(puzzle) {
			for i := leftIndex; i <= currentIndex; i++ {
				if i >= 0 && i < len(puzzle[j]) {
					str += string(puzzle[j][i])
				}
			}
			str += "\n"
		}
	}
	fmt.Print(str)
}

func part1(input []string) int {
	output := 0
	for i, line := range input {
		current := ""
		for j, char := range line {
			if isDigit(string(char)) {
				current += string(char)
			} else if current != "" {
				if isPartNumber(current, j, i, input) {
					// fmt.Println("part: line", i)
					// printNeighborhood(input, j, i, len(current))
					number, err := strconv.Atoi(current)
					if err != nil {
						panic(err)
					}
					// fmt.Printf("string %s number %d\n", current, number)
					output += number
				} else {
					// fmt.Printf("not part: line %d\n", i)
					// printNeighborhood(input, j, i, len(current))
				}
				current = ""
			}
		}
		// END OF STRING CASE!!!
		if current != "" && isPartNumber(current, len(line), i, input) {
			number, err := strconv.Atoi(current)
			if err != nil {
				panic(err)
			}
			// fmt.Printf("string %s number %d\n", current, number)
			output += number
		}
	}

	return output
}

// for every * character that is adjacent to exactly 2 numbers
// multiply those 2 numbers together and add the result to the output
func part2(input []string) int {
	output := 0
	fullString := strings.Join(input, "")
	rowLength := len(input[0])
	starRegex := regexp.MustCompile(`[*]`)
	numberRegex := regexp.MustCompile(`\d+`)
	starIndices := starRegex.FindAllStringIndex(fullString, -1)
	numberIndices := numberRegex.FindAllStringIndex(fullString, -1)
	// fmt.Printf("starIndices: %v\n", starIndices)
	// fmt.Printf("numberIndices: %v\n", numberIndices)
	for _, starIndex := range starIndices {
		// time.Sleep(time.Second / 2)
		adjacentNumbers := []int{}
		topLeft := starIndex[0] - rowLength - 1
		top := starIndex[0] - rowLength
		topRight := starIndex[1] - rowLength
		left := starIndex[0] - 1
		self := starIndex[0]
		right := starIndex[1]
		bottomLeft := starIndex[0] + rowLength - 1
		bottom := starIndex[0] + rowLength
		bottomRight := starIndex[1] + rowLength

		for _, numberIndex := range numberIndices {
			start := numberIndex[0]
			end := numberIndex[len(numberIndex)-1] - 1
			// fmt.Printf("start: %d end: %d\n", start, end)
			// fmt.Printf("topLeft: %d top: %d topRight: %d\nleft: %d right: %d\nbottomLeft: %d bottom: %d bottomRight: %d\n", topLeft, top, topRight, left, right, bottomLeft, bottom, bottomRight)
			if start == topLeft || start == top || start == topRight || start == left || start == self || start == right || start == bottomLeft || start == bottom || start == bottomRight || end == topLeft || end == top || end == topRight || end == left || end == self || end == right || end == bottomLeft || end == bottom || end == bottomRight {
				// fmt.Printf("found: %s\n", fullString[start:end+1])
				number, err := strconv.Atoi(fullString[start : end+1])
				if err != nil {
					panic(err)
				}
				adjacentNumbers = append(adjacentNumbers, number)
			}
		}
		if len(adjacentNumbers) == 2 {
			// printNeighborhood(input, (starIndex[1]%rowLength)+2, starIndex[0]/rowLength, 6)
			// fmt.Printf("adjacentNumbers: %v\n", adjacentNumbers)
			ratio := adjacentNumbers[0] * adjacentNumbers[1]
			// fmt.Printf("ratio: %d\n", ratio)
			output += ratio
			// fmt.Printf("output: %d\n\n", output)
		}
		// else {
		// 	println("\n\n ******* ERROR *******")
		// 	printNeighborhood(input, (starIndex[1]%rowLength)+2, starIndex[0]/rowLength, 6)
		// 	println(" ******* ERROR *******\n\n")
		// }
	}

	return output
}

func main() {
	year := 2023
	day := 3
	startTime := time.Now()
	input := utils.GetInputArray(year, day)
	postFetchTime := time.Now()
	result1 := part1(input)
	endTime1 := time.Now()
	result2 := part2(input)
	endTime2 := time.Now()

	printWidth := 25
	title := fmt.Sprintf("Advent of Code %d - Day %d", year, day)

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
