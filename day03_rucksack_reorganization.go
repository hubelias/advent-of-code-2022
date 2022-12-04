package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calculatePriority(item rune) int {
	if item < 97 {
		return int(item - 38)
	} else {
		return int(item - 96)
	}
}

func findItemPresentInBothCompartments(rucksackContents string) rune {
	firstCompartment := rucksackContents[:(len(rucksackContents) / 2)]
	secondCompartment := rucksackContents[(len(rucksackContents) / 2):]

	itemsInFirstCompartment := make(map[rune]bool)
	for _, char := range firstCompartment {
		itemsInFirstCompartment[char] = true
	}

	for _, char := range secondCompartment {
		if itemsInFirstCompartment[char] {
			return char
		}
	}

	return -1
}

func CalculateSumOfPrioritiesOfItemsInBothRucksackCompartments(input *os.File) (string, error) {
	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	totalPriority := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		priorityForLine := calculatePriority(findItemPresentInBothCompartments(line))
		fmt.Printf("Priority for \"%s\" is %d\n", line, priorityForLine)
		totalPriority += priorityForLine
	}

	return strconv.Itoa(totalPriority), nil
}

func findDuplicateItemAmongElves(rucksacks []string) rune {
	elfCountPerItem := make(map[rune]int)
	for _, rucksackContent := range rucksacks {
		elfItems := make(map[rune]bool)
		for _, item := range rucksackContent {
			elfItems[item] = true
		}
		for item, _ := range elfItems {
			elfCountPerItem[item] = elfCountPerItem[item] + 1
		}
	}
	for item, itemCount := range elfCountPerItem {
		if itemCount == len(rucksacks) {
			return item
		}
	}
	return -1
}

func CalculateSumOfPrioritiesOfCommonItemsPerElfGroup(input *os.File) (string, error) {
	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	totalPriority := 0
	lines := []string{}

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
		if len(lines) == 3 {
			commonItem := findDuplicateItemAmongElves(lines)
			totalPriority += calculatePriority(commonItem)
			lines = []string{}
		}
	}

	return strconv.Itoa(totalPriority), nil
}
