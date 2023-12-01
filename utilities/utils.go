package utils

import (
	"fmt"
	"io"
	"net/http"

	"os"
	"strconv"
	"strings"
)

func getCookie() string {
	content, err := os.ReadFile("./aoc_cookie")
	if err != nil {
		panic(err)
	}
	return "session=" + string(content)
}

func fetchInput(year int, day int) string {
	url := "https://adventofcode.com/" + strconv.Itoa(year) + "/day/" + strconv.Itoa(day) + "/input"

	fmt.Printf("Fetching input for year %d day %d\n", year, day)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Cookie", getCookie())

	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func createFilepath(year int, day int) string {
	return "./" + strconv.Itoa(year) + "/day-" + strconv.Itoa(day) + "/puzzle.txt"
}

func cacheInput(year int, day int) string {
	input := fetchInput(year, day)
	if strings.Contains(input, "Please log in") {
		fmt.Printf("Error fetching input for year %d day %d\n", year, day)
		fmt.Println(input)
		panic("Please log in to advent of code website and set AOC_SESSION environment variable")
	}
	fmt.Printf("Caching input for year %d day %d\n", year, day)
	filepath := createFilepath(year, day)
	err := os.WriteFile(filepath, []byte(input), 0644)

	if err != nil {
		panic(err)
	}

	return input
}

// If file has already been cached it will return the cached version
// parsed as an array of strings.
// if not it will fetch the file from the advent of code website
// and cache it locally.
func GetInput(year int, day int) []string {
	filepath := createFilepath(year, day)
	bytes, err := os.ReadFile(filepath)

	if err != nil {
		if os.IsNotExist(err) {
			return strings.Split(cacheInput(year, day), "\n")
		}

		panic(err)
	}

	str := string(bytes)

	return strings.Split(str, "\n")

}
