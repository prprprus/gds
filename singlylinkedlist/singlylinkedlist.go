package singlylinkedlist

import (
	"encoding/json"
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
func New(value ...interface{}) *List {
	list := &List{}

	if len(value) != 0 {
		list.Append(value...)
	}

	return list
}

// List Interface

// Append i.e. [] -> Append(1, 2, 3) -> [1, 2, 3]
func (list *List) Append(value ...interface{}) {
	if len(value) == 0 {
		return
	}

	// if size is equal to 0, it is a new singly linked list
	if list.size == 0 {
		for i, v := range value {
			newElement := &element{value: v}
			if i == 0 {
				list.frist = newElement
				list.last = newElement
				list.last.next = nil
			}
			list.last.next = newElement
			list.last = newElement
			list.last.next = nil
			list.size++
		}
	} else {
		for _, v := range value {
			newElement := &element{value: v}
			list.last.next = newElement
			list.last = newElement
			list.last.next = nil
			list.size++
		}
	}
}

// PreAppend i.e. [1, 2] -> Append(3, 4) -> [3, 4, 1, 2]
func (list *List) PreAppend(value ...interface{}) {
	if len(value) == 0 {
		return
	}

	for i := len(value) - 1; i >= 0; i-- {
		newElement := &element{value: value[i]}
		newElement.next = list.frist
		list.frist = newElement
		list.size++
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

	return values
}

// Iterator Interface

// Iterator iterator
type Iterator struct {
	list    *List    // list of iterator
	index   int      // current index
	element *element // current element
}

// Iterator new iterator
func (list *List) Iterator() *Iterator {
	return &Iterator{
		list:    list,
		index:   -1,
		element: nil,
	}
}

// Next
func (iterator *Iterator) Next() bool {
	if iterator.index == -1 {
		iterator.index++
		iterator.element = iterator.list.frist
		return true
	}

	if !iterator.list.indexInRange(iterator.index + 1) {
		return false
	}

	iterator.index++
	iterator.element = iterator.element.next
	return true
}

// Index
func (iterator *Iterator) Index() int {
	return iterator.index
}

// Value
func (iterator *Iterator) Value() interface{} {
	return iterator.element.value
}

// Begin
func (iterator *Iterator) Begin() {
	iterator.index = -1
	iterator.element = nil
}

// Serialization

// ToJSON
func (list *List) ToJSON() ([]byte, error) {
	return json.Marshal(list.Values())
}

// FromJSON
func (list *List) FromJSON(data []byte) error {
	elements := make([]interface{}, 0)
	err := json.Unmarshal(data, &elements)
	if err == nil {
		list.Clear()
		list.Append(elements...)
	}
	return err
}
