package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var hello = "hello"
var world = "world"

func TestStack(t *testing.T) {
	newStack := NewStack[string]()

	assert.Equal(t, 0, newStack.Depth())
	assert.Nil(t, newStack.Peek())
	assert.Nil(t, newStack.Pop())

	newStack.Push(&hello)

	assert.Equal(t, 1, newStack.Depth())
	assert.Equal(t, hello, *(newStack.Peek()))
	assert.Equal(t, 1, newStack.Depth())
	assert.Equal(t, hello, *(newStack.Pop()))
	assert.Equal(t, 0, newStack.Depth())

	newStack.Push(&hello)
	newStack.Push(&world)

	assert.Equal(t, 2, newStack.Depth())
	assert.Equal(t, world, *(newStack.Peek()))
	assert.Equal(t, 2, newStack.Depth())
	assert.Equal(t, world, *(newStack.Pop()))
	assert.Equal(t, 1, newStack.Depth())
	assert.Equal(t, hello, *(newStack.Pop()))
	assert.Equal(t, 0, newStack.Depth())
}
