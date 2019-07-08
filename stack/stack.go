package stack

// Stack Interface
type Stack interface {
	Push(value interface{})
	Pop() (interface{}, error)
	Peek() (interface{}, error)
	indexInRange(index int) bool
}
