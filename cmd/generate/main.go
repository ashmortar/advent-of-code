package main

import (
	utils "ashmortar/advent-of-code/utilities"
	"regexp"

	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func removeDir(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Println("Error removing directory:", err)
	}
}
func main() {
	flagYear := flag.String("year", "", "Year of the puzzle")
	flagDay := flag.String("day", "", "Day of the puzzle")
	flagReset := flag.Bool("reset", false, "Reset the cache for the given year day or all")
	flag.Parse()
	numberReg := regexp.MustCompile(`\d+`)
	year := string(*flagYear)
	day := string(*flagDay)

	if *flagReset {
		if year == "" {
			// remove all year directories
			contents, err := os.ReadDir(utils.RootDir())
			if err != nil {
				fmt.Println("Error reading directory:", err)
				return
			}
			for _, f := range contents {
				if f.IsDir() && numberReg.MatchString(f.Name()) {
					removeDir(filepath.Join(utils.RootDir(), f.Name()))
				}
			}
			fmt.Println("Done!")
			return
		}

		if day == "" {
			//  remove the given ye
			removeDir(filepath.Join(utils.RootDir(), year))
		} else {
			// remove the given year day
			removeDir(filepath.Join(utils.RootDir(), year, "day-"+day))
		}

		fmt.Println("Done!")
		return
	}

	if year == "" || day == "" {
		// scan through the directories and find the latest year and day
		// and use the next day / year and day
		contents, err := os.ReadDir(utils.RootDir())
		if err != nil {
			fmt.Println("Error reading directory:", err)
			return
		}
		year = "0"
		day = "0"
		for _, f := range contents {
			if f.IsDir() {
				matches := numberReg.FindAllString(f.Name(), -1)
				if len(matches) == 1 && matches[0] > year {
					year = matches[0]
					days, err := os.ReadDir(filepath.Join(utils.RootDir(), year))
					if err != nil {
						fmt.Println("Error reading directory:", err)
						return
					}
					for _, d := range days {
						if d.IsDir() {
							matches := numberReg.FindAllString(d.Name(), -1)
							if len(matches) == 1 && matches[0] > day {
								day = matches[0]
							}
						}
					}
				}
			}
		}
		if day == "0" || year == "0" {
			fmt.Println("Error finding next latest year and day, please submit as flags")
			return
		}
		if day == "25" {
			yearInt, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println("Error converting year to int:", err)
				return
			}
			yearInt++
			year = strconv.Itoa(yearInt)
			day = "1"
		} else {
			dayInt, err := strconv.Atoi(day)
			if err != nil {
				fmt.Println("Error converting day to int:", err)
				return
			}
			dayInt++
			day = strconv.Itoa(dayInt)
		}
	}

	if len(year) != 4 || day == "0" {
		fmt.Println("Error: year and day are required")
		return
	}

	// Create the directory structure if it doesn't exist
	dirPath := filepath.Join(utils.RootDir(), year, "day-"+day)
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
		sourceData, err := os.ReadFile(filepath.Join(utils.RootDir(), tpl.sourcePath))
		if err != nil {
			fmt.Println("Error reading template file:", err)
			return
		}

		// Replace placeholders with year and day values
		content := strings.ReplaceAll(string(sourceData), "{{Year}}", year)
		content = strings.ReplaceAll(content, "{{Day}}", day)

		// Write the modified content to the destination file
		if err := os.WriteFile(tpl.destPath, []byte(content), os.ModePerm); err != nil {
			fmt.Println("Error writing file:", err)
			return
		}

		fmt.Printf("Generated file: %s\n", tpl.destPath)
	}

	// cache the problem
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		fmt.Println("Error converting year to int:", err)
		return
	}
	dayInt, err := strconv.Atoi(day)
	if err != nil {
		fmt.Println("Error converting day to int:", err)
		return
	}
	utils.CacheProblem(yearInt, dayInt)
	fmt.Println("Done!")
}
