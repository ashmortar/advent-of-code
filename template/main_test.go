package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed puzzle.txt
var testInput string

func TestPart1(t *testing.T) {
	assert.Equal(t, 532428, Part1(testInput))
}

func BenchmarkPart1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Part1(testInput)
	}
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 84051670, Part2(testInput))
}

func BenchmarkPart2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Part2(testInput)
	}
}
