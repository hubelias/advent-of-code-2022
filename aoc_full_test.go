package main

import (
	"testing"
)

func TestAdventOfCode(t *testing.T) {
	testAoCDay(t, AdventOfCodeDay{
		title:       "--- Day 1: Calorie Counting ---",
		url:         "https://adventofcode.com/2022/day/1",
		examplesDir: "inputs/day01",
		partOne: AdventOfCodePuzzlePart{
			testFunc: CountMaxCaloriesByElf,
			examples: []Example{
				{
					comment:  "default example",
					filename: "default.txt",
					expected: "24000",
				},
				{
					comment:  "actual puzzle input",
					filename: "full.txt",
					expected: "64929",
				},
			},
		},
		partTwo: AdventOfCodePuzzlePart{
			testFunc: CountCaloriesByTopThreeElves,
			examples: []Example{
				{
					comment:  "default example",
					filename: "default.txt",
					expected: "45000",
				},
				{
					comment:  "default reversed",
					filename: "default_reversed.txt",
					expected: "45000",
				},
				{
					comment:  "actual puzzle input",
					filename: "full.txt",
					expected: "?",
				},
			},
		},
	})
}
