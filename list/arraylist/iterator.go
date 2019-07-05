package arraylist

// Iterator
type Iterator struct {
	list  *List
	index int
}

// Iterator
func (list *List) Iterator() *Iterator {
	return &Iterator{
		list:  list,
		index: -1,
	}
}

// Begin
func (iterator *Iterator) Begin() {
	iterator.index = -1
}

// Next
func (iterator *Iterator) Next() bool {
	if !iterator.list.indexInRange(iterator.index + 1) {
		return false
	}

	iterator.index++
	return true
}

// End
func (iterator *Iterator) End() {
	iterator.index = iterator.list.size
}

// Prev
func (iterator *Iterator) Prev() bool {
	if !iterator.list.indexInRange(iterator.index - 1) {
		return false
	}

	iterator.index--
	return true
}

// Value
func (iterator *Iterator) Value() interface{} {
	return iterator.list.elements[iterator.index]
}

// Index
func (iterator *Iterator) Index() int {
	return iterator.index
}
