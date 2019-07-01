package singlylinkedlist

import (
	"errors"
)

var (
	// ErrIndex
	ErrIndex = errors.New("index is not within the range of the list")
)

// List list
type List struct {
	frist *element
	last  *element
	size  int
}

type element struct {
	value interface{}
	next  *element
}

// New new singly linked list
func New(values ...interface{}) *List {
	list := &List{}

	if len(values) != 0 {
		list.Append(values...)
	}

	return list
}

// List Interface

// Append i.e. [] -> Append(1, 2, 3) -> [1, 2, 3]
func (list *List) Append(values ...interface{}) {
	if len(values) == 0 {
		return
	}

	// if size is equal to 0, it is a new singly linked list
	if list.size == 0 {
		for i, v := range values {
			newElement := &element{value: v}
			// head element
			if i == 0 {
				list.frist = newElement
				list.last = newElement
			} else {
				list.last.next = newElement
				list.last = newElement
			}
			list.size++
		}
	} else {
		for _, v := range values {
			newElement := &element{value: v}
			list.last.next = newElement
			list.last = newElement
			list.size++
		}
	}
}

// PreAppend i.e. [1, 2] -> Append(3, 4) -> [3, 4, 1, 2]
func (list *List) PreAppend(values ...interface{}) {
	if len(values) == 0 {
		return
	}

	if list.size == 0 {
		list.Append(values...)
	} else {
		for i := len(values) - 1; i >= 0; i-- {
			newElement := &element{value: values[i]}
			newElement.next = list.frist
			list.frist = newElement
			list.size++
		}
	}
}

// Check if the index is within the length of the list
func (list *List) indexInRange(index int) bool {
	if index >= 0 && index < list.size {
		return true
	}
	return false
}

// Get get value by index
func (list *List) Get(index int) (interface{}, error) {
	if !list.indexInRange(index) {
		return nil, ErrIndex
	}

	// find element by index
	foundElement := list.frist
	for i := 0; i != index; i++ {
		foundElement = foundElement.next
	}

	return foundElement.value, nil
}

// Remove remove element by index
func (list *List) Remove(index int) error {
	if !list.indexInRange(index) {
		return ErrIndex
	}

	foundElement := list.frist
	// if you want to remove the first element
	if index == 0 {
		list.frist = list.frist.next
	} else {
		preFoundElement := new(element)
		for i := 0; i != index; i++ {
			preFoundElement = foundElement
			foundElement = foundElement.next
		}
		preFoundElement.next = foundElement.next
	}

	foundElement = nil
	list.size--
	return nil
}

// Contains
func (list *List) Contains(value ...interface{}) bool {
	if len(value) == 0 {
		return true
	}

	// TODO

	return false
}

// Swap swap value by index
func (list *List) Swap(i, j int) error {
	if !list.indexInRange(i) || !list.indexInRange(j) {
		return ErrIndex
	}

	if i == j {
		return nil
	}

	// find element by i
	foundElementI := list.frist
	for index := 0; index != i; index++ {
		foundElementI = foundElementI.next
	}
	// find element by j
	foundElementJ := list.frist
	for index := 0; index != j; index++ {
		foundElementJ = foundElementJ.next
	}
	// swap
	foundElementI.value, foundElementJ.value = foundElementJ.value, foundElementI.value
	return nil
}

// Insert insert value (one or more) after index
func (list *List) Insert(index int, value ...interface{}) error {
	if len(value) == 0 {
		return nil
	}

	if !list.indexInRange(index) {
		return ErrIndex
	}

	// find element by index
	foundElement := list.frist
	for i := 0; i != index; i++ {
		foundElement = foundElement.next
	}
	foundElementNext := foundElement.next

	// insert
	for _, v := range value {
		element := &element{value: v}
		foundElement.next = element
		foundElement = element
	}
	foundElement.next = foundElementNext

	list.size += len(value)
	return nil
}

// Set set element by index
func (list *List) Set(index int, value interface{}) error {
	if !list.indexInRange(index) {
		return ErrIndex
	}

	// find element by index
	foundElement := list.frist
	for i := 0; i != index; i++ {
		foundElement = foundElement.next
	}
	foundElement.value = value
	return nil
}

// IndexOf get index by value
func (list *List) IndexOf(value interface{}) int {
	// TODO
	return 0
}

// Container Interface

// Empty
func (list *List) Empty() bool {
	return list.size == 0
}

// Size
func (list *List) Size() int {
	size := list.size
	return size
}

// Clear
func (list *List) Clear() {
	list.size = 0
	list.frist = nil
	list.last = nil
}

// Values
func (list *List) Values() interface{} {
	values := make([]interface{}, 0)

	iterator := list.Iterator()
	iterator.Begin()
	for iterator.Next() {
		values = append(values, iterator.Value())
	}

	// foundElement := list.frist
	// for foundElement != nil {
	// 	values = append(values, foundElement.value)
	// 	foundElement = foundElement.next
	// }

	return values
}
