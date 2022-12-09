package main

import (
	"io"
	"os"
	"strconv"
	"strings"
)

func FindNumberOfCharactersUntilEndOfFirstStartOfPacketMarker(input *os.File) (string, error) {
	return findNumberOfCharactersUntilEndOfFirstMarker(input, 4)
}

func FindNumberOfCharactersUntilEndOfFirstStartOfMessageMarker(input *os.File) (string, error) {
	return findNumberOfCharactersUntilEndOfFirstMarker(input, 14)
}

func findNumberOfCharactersUntilEndOfFirstMarker(input *os.File, expectDistinctSymbols int) (string, error) {
	b := new(strings.Builder)
	_, err := io.Copy(b, input)
	if err != nil {
		return "", err
	}
	result := doFindNumberOfCharactersUntilEndOfFirstMarker(b.String(), expectDistinctSymbols)
	return strconv.Itoa(result), nil
}

func doFindNumberOfCharactersUntilEndOfFirstMarker(signal string, expectDistinctSymbols int) int {
	lastNSymbols := make([]rune, expectDistinctSymbols)
	firstSymbol := rune(signal[0])
	for i := range lastNSymbols {
		lastNSymbols[i] = firstSymbol
	}
	for i, symbol := range signal {
		lastNSymbols[i%expectDistinctSymbols] = symbol
		if allDifferent(lastNSymbols) {
			return i + 1
		}
	}
	return -1
}

func allDifferent(symbols []rune) bool {
	m := make(map[rune]int, len(symbols))
	for _, symbol := range symbols {
		m[symbol] += 1
	}
	return len(m) == len(symbols)
}
