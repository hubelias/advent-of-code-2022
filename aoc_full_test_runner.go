package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func testAoCDay(t *testing.T, aocDay AdventOfCodeDay) {
	t.Run(aocDay.title, func(t *testing.T) {
		testAoCPuzzlePart(t, "Part One", aocDay.partOne, aocDay.examplesDir)
		testAoCPuzzlePart(t, "Part Two", aocDay.partTwo, aocDay.examplesDir)
	})

}

func testAoCPuzzlePart(t *testing.T, partName string, partSpec AdventOfCodePuzzlePart, examplesDir string) {
	t.Run(partName, func(t *testing.T) {
		for _, example := range partSpec.examples {
			t.Run(testName(example), func(t *testing.T) {

				fullFileName := path.Join(examplesDir, example.filename)
				input, err := os.Open(fullFileName)
				if err != nil {
					t.Fatalf("Could not open file: %s", fullFileName)
				}
				//goland:noinspection GoUnhandledErrorResult
				defer input.Close()

				result, err := partSpec.testFunc(input)
				if err != nil {
					t.Fatalf("Failed due to error %s", err)
				}
				assert.Equal(t, example.expected, result)
			})
		}
	})
}

func testName(input Example) string {
	if input.comment == "" {
		return input.filename
	}
	return fmt.Sprintf("%s (%s)", input.comment, input.filename)
}

type AdventOfCodeDay struct {
	title       string
	url         string
	examplesDir string
	partOne     AdventOfCodePuzzlePart
	partTwo     AdventOfCodePuzzlePart
}

type AdventOfCodePuzzlePart struct {
	testFunc PuzzleSolver
	examples []Example
}

type PuzzleSolver func(input *os.File) (string, error)

type Example struct {
	comment  string
	filename string
	expected string
}
