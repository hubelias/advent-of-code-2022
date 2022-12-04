package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOneRangeFullyContainsTheOther(t *testing.T) {
	assert.True(t, oneRangeFullyContainsTheOther(sectionRange{min: 0, max: 2}, sectionRange{min: 0, max: 2}))
	assert.True(t, oneRangeFullyContainsTheOther(sectionRange{min: 0, max: 2}, sectionRange{min: 1, max: 2}))
	assert.True(t, oneRangeFullyContainsTheOther(sectionRange{min: 0, max: 2}, sectionRange{min: 0, max: 1}))
	assert.False(t, oneRangeFullyContainsTheOther(sectionRange{min: 1, max: 3}, sectionRange{min: 0, max: 1}))
	assert.False(t, oneRangeFullyContainsTheOther(sectionRange{min: 1, max: 3}, sectionRange{min: 0, max: 2}))
	assert.False(t, oneRangeFullyContainsTheOther(sectionRange{min: 1, max: 3}, sectionRange{min: 2, max: 4}))
	assert.True(t, oneRangeFullyContainsTheOther(sectionRange{min: 0, max: 2}, sectionRange{min: 0, max: 2}))
	assert.True(t, oneRangeFullyContainsTheOther(sectionRange{min: 1, max: 2}, sectionRange{min: 0, max: 2}))
	assert.True(t, oneRangeFullyContainsTheOther(sectionRange{min: 0, max: 1}, sectionRange{min: 0, max: 2}))
	assert.False(t, oneRangeFullyContainsTheOther(sectionRange{min: 0, max: 1}, sectionRange{min: 1, max: 3}))
	assert.False(t, oneRangeFullyContainsTheOther(sectionRange{min: 0, max: 2}, sectionRange{min: 1, max: 3}))
	assert.False(t, oneRangeFullyContainsTheOther(sectionRange{min: 2, max: 4}, sectionRange{min: 1, max: 3}))
}

func TestOneRangeOverlapsTheOther(t *testing.T) {
	assert.True(t, oneRangeOverlapsTheOther(sectionRange{min: 1, max: 2}, sectionRange{min: 0, max: 1}))
	assert.True(t, oneRangeOverlapsTheOther(sectionRange{min: 1, max: 2}, sectionRange{min: 2, max: 3}))
	assert.True(t, oneRangeOverlapsTheOther(sectionRange{min: 0, max: 1}, sectionRange{min: 1, max: 2}))
	assert.True(t, oneRangeOverlapsTheOther(sectionRange{min: 2, max: 3}, sectionRange{min: 1, max: 2}))
	assert.False(t, oneRangeOverlapsTheOther(sectionRange{min: 0, max: 1}, sectionRange{min: 2, max: 3}))
	assert.False(t, oneRangeOverlapsTheOther(sectionRange{min: 2, max: 3}, sectionRange{min: 0, max: 1}))
}
