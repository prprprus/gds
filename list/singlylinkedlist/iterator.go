package singlylinkedlist

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
	if !iterator.list.indexInRange(iterator.index + 1) {
		return false
	}

	if iterator.index == -1 {
		iterator.index++
		iterator.element = iterator.list.frist
		return true
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
