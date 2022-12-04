package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type sectionRange struct {
	min int
	max int
}

func CountPairsWhereOneRangeFullyContainsTheOther(input *os.File) (string, error) {
	return countPairsWhere(input, oneRangeFullyContainsTheOther)
}

func CountPairsWhereOneRangeOverlapsTheOther(input *os.File) (string, error) {
	return countPairsWhere(input, oneRangeOverlapsTheOther)
}

func countPairsWhere(input *os.File, condition func(firstRange sectionRange, secondRange sectionRange) bool) (string, error) {
	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	result := 0

	for fileScanner.Scan() {
		ranges := strings.Split(fileScanner.Text(), ",")

		firstRange, err := parseRange(ranges[0])
		if err != nil {
			return "", err
		}
		secondRange, err := parseRange(ranges[1])
		if err != nil {
			return "", err
		}
		if condition(*firstRange, *secondRange) {
			result++
		}
	}

	return strconv.Itoa(result), nil
}

func parseRange(rep string) (*sectionRange, error) {
	minMax := strings.Split(rep, "-")
	min, err := strconv.Atoi(minMax[0])
	if err != nil {
		return nil, err
	}
	max, err := strconv.Atoi(minMax[1])
	if err != nil {
		return nil, err
	}
	return &sectionRange{min: min, max: max}, nil
}

func oneRangeFullyContainsTheOther(firstRange sectionRange, secondRange sectionRange) bool {
	return (firstRange.min >= secondRange.min && firstRange.max <= secondRange.max) ||
		(secondRange.min >= firstRange.min && secondRange.max <= firstRange.max)
}

func oneRangeOverlapsTheOther(firstRange sectionRange, secondRange sectionRange) bool {
	return (firstRange.max >= secondRange.min && firstRange.min <= secondRange.max)
}
