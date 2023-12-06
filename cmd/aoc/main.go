package main

import (
	utils "ashmortar/advent-of-code/utilities"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
)

func run(year string, day string) {
	cmd := exec.Command("go", "run", "main.go")
	cmd.Dir = filepath.Join(utils.RootDir(), year, "day-"+day)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func main() {
	flagYear := flag.String("year", "", "Year of the puzzle")
	flagDay := flag.String("day", "", "Day of the puzzle")
	flag.Parse()
	year := string(*flagYear)
	day := string(*flagDay)
	//
	if year != "" {
		// if we have year and day run the main.go file in that day's directory
		if day != "" {
			run(year, day)
			fmt.Println("Done")
			return
		}
		// if year was provided but not day run the main.go file in each day's
		// directory in chronological order
		for i := 1; i <= 25; i++ {
			day := strconv.Itoa(i)
			run(year, day)
		}
		fmt.Println("Done")
		return
	}
	// if year and day weren't provided loop through all year directories and
	// run the main.go file in each day's directory in chronological order
	contents, err := os.ReadDir(utils.RootDir())
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}
	years := []string{}
	numberReg := regexp.MustCompile(`\d+`)
	for _, f := range contents {
		if f.IsDir() && numberReg.MatchString(f.Name()) {
			years = append(years, f.Name())
		}
	}
	sort.Slice(years, func(i, j int) bool {
		return years[i] < years[j]
	})

	for _, year := range years {
		days, err := os.ReadDir(filepath.Join(utils.RootDir(), year))
		if err != nil {
			fmt.Println("Error reading directory:", err)
			return
		}
		for _, dayDir := range days {
			day := numberReg.FindString(dayDir.Name())
			run(year, day)
		}
	}
	fmt.Println("Done")
}
