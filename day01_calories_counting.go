package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func CountMaxCaloriesByElf(input *os.File) (string, error) {
	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	maxElfCalories := 0
	currentElfCalories := 0

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()
		if currentLine == "" {
			if currentElfCalories > maxElfCalories {
				maxElfCalories = currentElfCalories
			}
			currentElfCalories = 0
			continue
		}
		foodCalories, err := strconv.Atoi(currentLine)
		if err != nil {
			return "", err
		}

		currentElfCalories += foodCalories
	}

	if currentElfCalories > maxElfCalories {
		maxElfCalories = currentElfCalories
	}

	return strconv.Itoa(maxElfCalories), nil
}

func CountCaloriesByTopThreeElves(input *os.File) (string, error) {
	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	topThreeElfCallories := [3]int{0, 0, 0}
	currentElfCalories := 0

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()
		if currentLine == "" {
			fmt.Printf("New elf found: %d \n", currentElfCalories)
			for i := range topThreeElfCallories {
				if currentElfCalories >= topThreeElfCallories[i] {
					prevElfCallories := topThreeElfCallories[i]
					topThreeElfCallories[i] = currentElfCalories
					currentElfCalories = prevElfCallories
				}
			}
			fmt.Printf("Top 3 elves: %v \n", topThreeElfCallories)
			currentElfCalories = 0
			continue
		}
		foodCalories, err := strconv.Atoi(currentLine)
		if err != nil {
			return "", err
		}

		currentElfCalories += foodCalories
	}

	for i := range topThreeElfCallories {
		if currentElfCalories >= topThreeElfCallories[i] {
			prevElfCallories := topThreeElfCallories[i]
			topThreeElfCallories[i] = currentElfCalories
			currentElfCalories = prevElfCallories
		}
	}

	topThreeElvesTotalCalories := 0
	for i := range topThreeElfCallories {
		topThreeElvesTotalCalories += topThreeElfCallories[i]
	}

	return strconv.Itoa(topThreeElvesTotalCalories), nil
}
