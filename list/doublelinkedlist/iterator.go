package doublelinkedlist

// Iterator
type Iterator struct {
	list    *List
	index   int
	element *element
}

// Iterator
func (list *List) Iterator() *Iterator {
	return &Iterator{
		list:    list,
		index:   -1,
		element: nil,
	}
}

// Begin
func (iterator *Iterator) Begin() {
	iterator.index = -1
	iterator.element = nil
}

// Next
func (iterator *Iterator) Next() bool {
	if !iterator.list.indexInRange(iterator.index + 1) {
		return false
	}

	if iterator.index == -1 {
		iterator.index++
		iterator.element = iterator.list.first
		return true
	}

	iterator.index++
	iterator.element = iterator.element.next
	return true
}

// End
func (iterator *Iterator) End() {
	iterator.index = iterator.list.size
	iterator.element = iterator.list.last
}

// Prev
func (iterator *Iterator) Prev() bool {
	if !iterator.list.indexInRange(iterator.index - 1) {
		return false
	}

	if iterator.index == iterator.list.size {
		iterator.index--
		iterator.element = iterator.list.last
		return true
	}

	iterator.index--
	iterator.element = iterator.element.prev
	return true
}

// Value
func (iterator *Iterator) Value() interface{} {
	return iterator.element.value
}

// Index
func (iterator *Iterator) Index() int {
	return iterator.index
}
