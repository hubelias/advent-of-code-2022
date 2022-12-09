package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type crate struct {
	content rune
}

func readInitialLayout(drawingRows []string) []Stack[crate] {
	stacks := make([]Stack[crate], 0)
	for rowIdx := len(drawingRows) - 1; rowIdx >= 0; rowIdx-- {
		row := drawingRows[rowIdx]
		for i := 0; i < len(row); i += 4 {
			var currentStack Stack[crate]
			if len(stacks) <= i/4 {
				currentStack = NewStack[crate]()
				stacks = append(stacks, currentStack)
			} else {
				currentStack = stacks[i/4]
			}

			crateAsText := strings.TrimSpace(row[i:min(len(row), i+4)])
			if len(crateAsText) == 3 {
				nextCrate := crate{content: []rune(crateAsText)[1]}
				currentStack.Push(&nextCrate)
			}
		}
	}
	return stacks
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func MoveCratesWithCrateMover9000(input *os.File) (string, error) {
	return moveCratesAccordingToInstructions(input, func(from Stack[crate], to Stack[crate], count int) {
		for i := 0; i < count; i++ {
			to.Push(from.Pop())
		}
	})
}

func MoveCratesWithCrateMover9001(input *os.File) (string, error) {
	return moveCratesAccordingToInstructions(input, func(from Stack[crate], to Stack[crate], count int) {
		cratesToMove := []*crate{}
		for i := 0; i < count; i++ {
			cratesToMove = append(cratesToMove, from.Pop())
		}
		for i := count - 1; i >= 0; i-- {
			to.Push(cratesToMove[i])
		}
	})
}

func moveCratesAccordingToInstructions(input *os.File, mover CrateMover) (string, error) {
	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	drawingRows := []string{}
	//instructions := []string{}

	for fileScanner.Scan() {
		nextLine := fileScanner.Text()
		drawingRows = append(drawingRows, nextLine)
		if len(strings.TrimSpace(nextLine)) == 0 {
			drawingRows = drawingRows[:len(drawingRows)-2] // remove row with legend
			break
		}
	}

	allStacks := readInitialLayout(drawingRows)
	for fileScanner.Scan() {
		moveCrates(allStacks, fileScanner.Text(), mover)
	}

	result := ""
	for _, stack := range allStacks {
		result += string(stack.Peek().content)
	}

	return result, nil
}

// sample instruction: "move 1 from 2 to 1"
// move 1 create from stack number 2 (idx=1) to stack number 1 (idx=0)
func moveCrates(allStacks []Stack[crate], instruction string, mover CrateMover) {
	tokens := strings.Split(instruction, " ")
	count, _ := strconv.Atoi(tokens[1])
	from, _ := strconv.Atoi(tokens[3])
	to, _ := strconv.Atoi(tokens[5])

	mover(allStacks[from-1], allStacks[to-1], count)
}

type CrateMover = func(from Stack[crate], to Stack[crate], count int)
