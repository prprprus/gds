package singlylinkedlist

// Iterator represents an iterable structure of list.
type Iterator struct {
	list    *List    // list of iterator
	index   int      // current index
	element *element // current element
}

// Iterator returns the iterator object of list.
func (list *List) Iterator() *Iterator {
	return &Iterator{
		list:    list,
		index:   -1,
		element: nil,
	}
}

// Next returns true if there are still iterable objects, false otherwise.
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

// Index returns the current index of the iterator.
func (iterator *Iterator) Index() int {
	return iterator.index
}

// Value returns the current value of the element of the iterator.
func (iterator *Iterator) Value() interface{} {
	if iterator.element != nil {
		return iterator.element.value
	}
	return nil
}

// Begin reset the iterator to the initial status.
func (iterator *Iterator) Begin() {
	iterator.index = -1
	iterator.element = nil
}
