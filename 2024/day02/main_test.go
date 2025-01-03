package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInputFile = "testInput.txt"
var testInput string

var solutionPart1 = "2"
var solutionPart2 = "4"

func init() {
	readBytes, err := os.ReadFile(testInputFile)
	if err != nil {
		fmt.Printf("Failed to read %s with error %s", testInputFile, err)
	}

	if string(readBytes) == "" {
		fmt.Printf("Please update the %s file", testInputFile)
		os.Exit(1)
	}

	testInput = string(readBytes)
}

func TestSolvePart1(t *testing.T) {
	assert.NotEmpty(t, solutionPart1)
	assert.Equal(t, solutionPart1, solvePart1(testInput), "Please provide the test solution for part 1")
}

func TestSolvePart2(t *testing.T) {
	assert.NotEmpty(t, solutionPart2, "Please provide the test solution for part 2")
	assert.Equal(t, solutionPart2, solvePart2(testInput))
}
