package arraylist

import "errors"

var (
	// ErrIndex is returned when the index is out of the list
	ErrIndex = errors.New("index is out of the list")
	// ErrIndexOf is returned when the index of value can not found
	ErrIndexOf = errors.New("index of value can not found")
)

const (
	growthFactor = float32(2.0)
	shrinkFactor = float32(0.25)
)

// List
type List struct {
	elements []interface{}
	size     int
}

func (list *List) growth(n int) {
	currCap := cap(list.elements)
	if list.size+n >= currCap {
		newCap := int(float32(currCap) * growthFactor)
		list.resize(newCap)
	}

	list.size += n
}

func (list *List) shrink() {
	if shrinkFactor == 0.0 {
		return
	}

	currCap := cap(list.elements)
	if list.size <= int(float32(currCap)*shrinkFactor) {
		list.resize(list.size)
	}

	list.size--
}

func (list *List) resize(cap int) {
	newElements := make([]interface{}, cap, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}

// List Interface

// New
func New(values ...interface{}) *List {
	list := &List{}
	if len(values) != 0 {
		list.Append(values...)
	}
	return list
}

// Append
func (list *List) Append(values ...interface{}) {
	list.growth(len(values))
	index := list.size
	for _, v := range values {
		list.elements[index] = v
		index++
	}
}

func (list *List) indexInRange(index int) bool {
	if index >= 0 && index < list.size {
		return true
	}
	return false
}

// Get
func (list *List) Get(index int) (interface{}, error) {
	if !list.indexInRange(index) {
		return nil, ErrIndex
	}

	return list.elements[index], nil
}

// Remove
func (list *List) Remove(index int) error {
	if !list.indexInRange(index) {
		return ErrIndex
	}

	list.elements[index] = nil
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.shrink()

	return nil
}

// Contains
func (list *List) Contains(values ...interface{}) bool {
	if len(values) == 0 {
		return true
	}
	if list.size == 0 {
		return false
	}

	for _, value := range values {
		found := false
		for _, element := range list.elements {
			if element == value {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Swap
func (list *List) Swap(i, j int) error {
	if !list.indexInRange(i) || !list.indexInRange(j) {
		return ErrIndex
	}
	if i == j {
		return nil
	}

	list.elements[i], list.elements[j] = list.elements[j], list.elements[i]

	return nil
}

// Insert
func (list *List) Insert(index int, values ...interface{}) error {
	if !list.indexInRange(index) {
		return ErrIndex
	}
	if len(values) == 0 {
		return nil
	}
	if index == list.size-1 {
		list.Append(values...)
		return nil
	}

	l := len(values)
	list.growth(l)
	copy(list.elements[index+1+l:], list.elements[index+1:])
	copy(list.elements[index+1:], values)

	return nil
}

// Set
func (list *List) Set(index int, value interface{}) error {
	if !list.indexInRange(index) {
		return ErrIndex
	}

	list.elements[index] = value

	return nil
}

// IndexOf
func (list *List) IndexOf(value interface{}) (int, error) {
	for index, element := range list.elements {
		if element == value {
			return index, nil
		}
	}
	return -1, ErrIndexOf
}

// Container Interface

// Empty
func (list *List) Empty() bool {
	return list.size == 0
}

// Size
func (list *List) Size() int {
	return list.size
}

// Clear
func (list *List) Clear() {
	list.size = 0
	list.elements = []interface{}{}
}

// Values
func (list *List) Values() []interface{} {
	newElements := make([]interface{}, list.size, list.size)
	copy(newElements, list.elements)
	return newElements
}
