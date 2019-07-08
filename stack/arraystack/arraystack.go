package arraystack

import (
	"errors"

	"github.com/prprprus/ds/list/arraylist"
)

var (
	ErrPop = errors.New("stack is empty")
)

type Stack struct {
	list *arraylist.List
}

func New() *Stack {
	return &Stack{
		list: arraylist.New(),
	}
}

// Stack Interface

func (stack *Stack) Push(value interface{}) {
	stack.list.Append(value)
}

func (stack *Stack) Pop() (interface{}, error) {
	if stack.Size() == 0 {
		return nil, ErrPop
	}

	element, _ := stack.list.Get(stack.list.Size() - 1)
	stack.list.Remove(stack.list.Size() - 1)
	return element, nil
}

func (stack *Stack) Peek() (interface{}, error) {
	if stack.Size() == 0 {
		return nil, ErrPop
	}

	element, _ := stack.list.Get(stack.list.Size() - 1)
	return element, nil
}

// Check if the index is within the length of the list.
func (stack *Stack) indexInRange(index int) bool {
	if index >= 0 && index < stack.list.Size() {
		return true
	}
	return false
}

// Container Interface

func (stack *Stack) Empty() bool {
	return stack.list.Size() == 0
}

func (stack *Stack) Size() int {
	return stack.list.Size()
}

func (stack *Stack) Clear() {
	stack.list.Clear()
}

func (stack *Stack) Values() []interface{} {
	return stack.list.Values()
}
