package main

import (
	utils "ashmortar/advent-of-code/utilities"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	year := flag.String("year", "", "Year of the puzzle")
	day := flag.String("day", "", "Day of the puzzle")
	flag.Parse()

	if *year == "" || *day == "" {
		fmt.Println("Year and day are required")
		return
	}

	// Create the directory structure if it doesn't exist
	dirPath := filepath.Join(*year, "day-"+*day)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		fmt.Println("Error creating directories:", err)
		return
	}

	// List of template files to copy and replace placeholders
	templates := []struct {
		sourcePath string
		destPath   string
	}{
		{"template/main_test.go", filepath.Join(dirPath, "main_test.go")},
		{"template/main.go", filepath.Join(dirPath, "main.go")},
	}

	for _, tpl := range templates {
		sourceData, err := os.ReadFile(tpl.sourcePath)
		if err != nil {
			fmt.Println("Error reading template file:", err)
			return
		}

		// Replace placeholders with year and day values
		content := strings.ReplaceAll(string(sourceData), "{{Year}}", *year)
		content = strings.ReplaceAll(content, "{{Day}}", *day)

		// Write the modified content to the destination file
		if err := os.WriteFile(tpl.destPath, []byte(content), os.ModePerm); err != nil {
			fmt.Println("Error writing file:", err)
			return
		}

		fmt.Printf("Generated file: %s\n", tpl.destPath)
	}

	// cache the problem
	yearInt, err := strconv.Atoi(*year)
	if err != nil {
		fmt.Println("Error converting year to int:", err)
		return
	}
	dayInt, err := strconv.Atoi(*day)
	if err != nil {
		fmt.Println("Error converting day to int:", err)
		return
	}
	utils.CacheProblem(yearInt, dayInt)
	fmt.Println("Done!")
}
