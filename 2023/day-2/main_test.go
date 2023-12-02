package main

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed puzzle.txt
var testInput string

func TestPart1(t *testing.T) {
	input := strings.Split(testInput, "\n")
	assert.Equal(t, 2447, part1(input))
}

func BenchmarkPart1(b *testing.B) {
	input := strings.Split(testInput, "\n")
	for n := 0; n < b.N; n++ {
		part1(input)
	}
}

func TestPart2(t *testing.T) {
	input := strings.Split(testInput, "\n")
	assert.Equal(t, 56322, part2(input))
}

func BenchmarkPart2(b *testing.B) {
	input := strings.Split(testInput, "\n")
	for n := 0; n < b.N; n++ {
		part2(input)
	}
}
