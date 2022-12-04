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
					expected: "193697",
				},
			},
		},
	})
	testAoCDay(t, AdventOfCodeDay{
		title:       "--- Day 2: Rock Paper Scissors ---",
		url:         "https://adventofcode.com/2022/day/2",
		examplesDir: "inputs/day02",
		partOne: AdventOfCodePuzzlePart{
			testFunc: CalculateScoreWithGuessedStrategy,
			examples: []Example{
				{
					filename: "default.txt",
					expected: "15",
				},
				{
					comment:  "actual puzzle input",
					filename: "full.txt",
					expected: "15337",
				},
			},
		},
		partTwo: AdventOfCodePuzzlePart{
			testFunc: CalculateScoreWithCorrectStrategy,
			examples: []Example{
				{
					filename: "default.txt",
					expected: "12",
				},
				{
					comment:  "actual puzzle input",
					filename: "full.txt",
					expected: "11696",
				},
			},
		},
	})
	testAoCDay(t, AdventOfCodeDay{
		title:       "--- Day 3: Rucksack Reorganization ---",
		url:         "https://adventofcode.com/2022/day/3",
		examplesDir: "inputs/day03",
		partOne: AdventOfCodePuzzlePart{
			testFunc: CalculateSumOfPrioritiesOfItemsInBothRucksackCompartments,
			examples: []Example{
				{
					filename: "default.txt",
					expected: "157",
				},
				{
					comment:  "actual puzzle input",
					filename: "full.txt",
					expected: "7597",
				},
			},
		},
		partTwo: AdventOfCodePuzzlePart{
			testFunc: CalculateSumOfPrioritiesOfCommonItemsPerElfGroup,
			examples: []Example{
				{
					filename: "default.txt",
					expected: "70",
				},
				{
					comment:  "actual puzzle input",
					filename: "full.txt",
					expected: "2607",
				},
			},
		},
	})
	testAoCDay(t, AdventOfCodeDay{
		title:       "--- Day 4: Camp Cleanup ---",
		url:         "https://adventofcode.com/2022/day/4",
		examplesDir: "inputs/day04",
		partOne: AdventOfCodePuzzlePart{
			testFunc: CountPairsWhereOneRangeFullyContainsTheOther,
			examples: []Example{
				{
					filename: "default.txt",
					expected: "2",
				},
				{
					comment:  "actual puzzle input",
					filename: "full.txt",
					expected: "651",
				},
			},
		},
		partTwo: AdventOfCodePuzzlePart{
			testFunc: CountPairsWhereOneRangeOverlapsTheOther,
			examples: []Example{
				{
					filename: "default.txt",
					expected: "4",
				},
				{
					comment:  "actual puzzle input",
					filename: "full.txt",
					expected: "956",
				},
			},
		},
	})
}
