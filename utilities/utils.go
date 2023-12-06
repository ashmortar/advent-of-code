package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/yosssi/gohtml"
)

func getCookie() string {
	content, err := os.ReadFile(filepath.Join(RootDir(), "aoc_cookie"))
	if err != nil {
		println("Please log in to advent of code website and place cookie into aoc_cookie file at root level")
		panic(err)
	}
	return "session=" + string(content)
}

func fetchProblem(year int, day int) string {
	url := "https://adventofcode.com/" + strconv.Itoa(year) + "/day/" + strconv.Itoa(day)

	fmt.Printf("Fetching problem for year %d day %d\n", year, day)
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

	// problem is in the first <article> tag
	problem := string(bytes)
	return problem[strings.Index(problem, "<article") : strings.Index(problem, "</article>")+10]
}

func CacheProblem(year int, day int) string {
	problem := fetchProblem(year, day)
	problem = gohtml.Format(problem)
	fmt.Printf("Caching problem for year %d day %d\n", year, day)
	filepath := CreateFilepath(year, day, "problem.html")
	err := os.WriteFile(filepath, []byte(problem), 0644)

	if err != nil {
		panic(err)
	}

	return problem
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

func CreateFilepath(year int, day int, filename string) string {
	return filepath.Join(RootDir(), strconv.Itoa(year), "day-"+strconv.Itoa(day), filename)
}

func cacheInput(year int, day int) string {
	input := fetchInput(year, day)
	if strings.Contains(input, "Please log in") {
		fmt.Printf("Error fetching input for year %d day %d\n", year, day)
		fmt.Println(input)
		panic("Please log in to advent of code website and set AOC_SESSION environment variable")
	}
	fmt.Printf("Caching input for year %d day %d\n", year, day)
	filepath := CreateFilepath(year, day, "puzzle.txt")
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
func GetInputArray(year int, day int) []string {
	return strings.Split(GetInputString(year, day), "\n")
}

func GetInputString(year int, day int) string {
	filepath := CreateFilepath(year, day, "puzzle.txt")
	bytes, err := os.ReadFile(filepath)

	if err != nil {
		if os.IsNotExist(err) {
			return cacheInput(year, day)
		}

		panic(err)
	}

	return string(bytes)

}

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
