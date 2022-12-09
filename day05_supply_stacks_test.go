package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_readInitialLayout(t *testing.T) {
	result := readInitialLayout([]string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
	})

	assert.Equal(t, 2, result[0].Depth())
	assert.Equal(t, 3, result[1].Depth())
	assert.Equal(t, 1, result[2].Depth())
	assert.Equal(t, "N", string(result[0].Pop().content))
	assert.Equal(t, "Z", string(result[0].Pop().content))
}
