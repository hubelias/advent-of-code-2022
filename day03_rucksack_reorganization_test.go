package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToRuneConversion(t *testing.T) {
	assert.Equal(t, []int32{97, 98, 99, 120, 121, 122}, []rune("abcxyz"))
	assert.Equal(t, []int32{65, 66, 67, 88, 89, 90}, []rune("ABCXYZ"))
}

func TestCalculatePriority(t *testing.T) {
	assert.Equal(t, 1, calculatePriority('a'))
	assert.Equal(t, 26, calculatePriority('z'))
	assert.Equal(t, 27, calculatePriority('A'))
	assert.Equal(t, 52, calculatePriority('Z'))
}

func TestFindItemPresentInBothCompartments(t *testing.T) {
	assert.Equal(t, []rune("p")[0], findItemPresentInBothCompartments("vJrwpWtwJgWrhcsFMMfFFhFp"))
	assert.Equal(t, []rune("L")[0], findItemPresentInBothCompartments("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"))
	assert.Equal(t, []rune("P")[0], findItemPresentInBothCompartments("PmmdzqPrVvPwwTWBwg"))
	assert.Equal(t, rune(-1), findItemPresentInBothCompartments("Aa"))
	assert.Equal(t, []rune("C")[0], findItemPresentInBothCompartments("ABCCba"))
}

func TestFindDuplicateItemAmongElves(t *testing.T) {
	assert.Equal(t, []rune("r")[0], findDuplicateItemAmongElves([]string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"}))
	assert.Equal(t, []rune("Z")[0], findDuplicateItemAmongElves([]string{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"}))
	assert.Equal(t, rune(-1), findDuplicateItemAmongElves([]string{"aa", "bb", "cc"}))
}
