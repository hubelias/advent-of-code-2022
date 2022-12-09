package main

import (
	"errors"
	"os"
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
	testAoCDay(t, AdventOfCodeDay{
		title:       "--- Day 5: Supply Stacks ---",
		url:         "https://adventofcode.com/2022/day/5",
		examplesDir: "inputs/day05",
		partOne: AdventOfCodePuzzlePart{
			testFunc: MoveCratesWithCrateMover9000,
			examples: []Example{
				{
					filename: "default.txt",
					expected: "CMZ",
				},
				{
					comment:  "actual puzzle input",
					filename: "full.txt",
					expected: "RFFFWBPNS",
				},
			},
		},
		partTwo: AdventOfCodePuzzlePart{
			testFunc: MoveCratesWithCrateMover9001,
			examples: []Example{
				{
					filename: "default.txt",
					expected: "MCD",
				},
				{
					comment:  "actual puzzle input",
					filename: "full.txt",
					expected: "CQQBBJFCS",
				},
			},
		},
	})
	testAoCDay(t, AdventOfCodeDay{
		title:       "--- Day 6: Tuning Trouble ---",
		url:         "https://adventofcode.com/2022/day/6",
		examplesDir: "inputs/day06",
		partOne: AdventOfCodePuzzlePart{
			testFunc: FindNumberOfCharactersUntilEndOfFirstStartOfPacketMarker,
			examples: []Example{
				{filename: "ex1.txt", expected: "5"},
				{filename: "ex2.txt", expected: "6"},
				{filename: "ex3.txt", expected: "10"},
				{filename: "ex4.txt", expected: "11"},
				{
					comment:  "actual puzzle input",
					filename: "full.txt",
					expected: "1142",
				},
			},
		},
		partTwo: AdventOfCodePuzzlePart{
			testFunc: FindNumberOfCharactersUntilEndOfFirstStartOfMessageMarker,
			examples: []Example{
				{filename: "ex1.txt", expected: "23"},
				{filename: "ex2.txt", expected: "23"},
				{filename: "ex3.txt", expected: "29"},
				{filename: "ex4.txt", expected: "26"},
				{
					comment:  "actual puzzle input",
					filename: "full.txt",
					expected: "?",
				},
			},
		},
	})
}

var TODO PuzzleSolver = func(input *os.File) (string, error) {
	return "", errors.New("not implemented")
}
