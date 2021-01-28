package stack

import (
	"errors"
)

type ArrayStack struct {
	count    int
	capacity int
	items    []string
}

func (arrayStack *ArrayStack) Push(item string) error {
	if arrayStack.capacity == arrayStack.count {
		return errors.New("the stack is full")
	}

	arrayStack.items[arrayStack.count] = item
	arrayStack.count++
	return nil
}

func (arrayStack *ArrayStack) Pop() (error, string) {
	if arrayStack.count == 0 {
		return errors.New("the stack is empty"), ""
	}

	top := arrayStack.items[arrayStack.count-1]
	arrayStack.count--
	return nil, top
}

func (arrayStack *ArrayStack) Init(capacity int) error {
	arrayStack.capacity = capacity
	arrayStack.items = make([]string, capacity)
	return nil
}
