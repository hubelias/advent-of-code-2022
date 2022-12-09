package main

type Stack[T any] interface {
	Push(*T)
	Pop() *T
	Peek() *T
	Depth() int
}

type stack[T any] struct {
	items []*T
}

func NewStack[T any]() Stack[T] {
	return &stack[T]{items: make([]*T, 0)}
}

func (this *stack[T]) Push(element *T) {
	this.items = append(this.items, element)
}

func (this *stack[T]) Peek() *T {
	idx := len(this.items) - 1
	if idx < 0 {
		return nil
	} else {
		return this.items[idx]
	}
}

func (this *stack[T]) Pop() *T {
	idx := len(this.items) - 1
	if idx < 0 {
		return nil
	} else {
		lastElement := this.items[idx]
		this.items = this.items[:idx]
		return lastElement
	}
}

func (this *stack[T]) Depth() int {
	return len(this.items)
}
