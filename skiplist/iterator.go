package skiplist

type Iterator struct {
	skiplist *SkipList
	element  *node
	index    int
}

func (skiplist *SkipList) New() *Iterator {
	return &Iterator{
		skiplist: skiplist,
		element:  nil,
		index:    -1,
	}
}

func (iterator *Iterator) Next() bool {
	if !iterator.skiplist.indexInRange(iterator.index + 1) {
		return false
	}

	if iterator.index == -1 {
		iterator.element = iterator.skiplist.head.next[0]
		iterator.index++
		return true
	}

	iterator.element = iterator.element.next[0]
	iterator.index++
	return true
}

func (iterator *Iterator) Begin() {
	iterator.element = nil
	iterator.index = -1
}

func (iterator *Iterator) Value() interface{} {
	return iterator.element.value
}

func (iterator *Iterator) Index() int {
	return iterator.index
}
