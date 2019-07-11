package skipmap

import "github.com/prprprus/ds/skiplist"

type Iterator struct {
	m                *skiplist.SkipList
	internalIterator *skiplist.Iterator
}

func (m *Map) Iterator() *Iterator {
	return &Iterator{
		m:                m.m,
		internalIterator: m.m.Iterator(),
	}
}

func (iterator *Iterator) Next() bool {
	return iterator.internalIterator.Next()
}

func (iterator *Iterator) Begin() {
	iterator.internalIterator.Begin()
}

func (iterator *Iterator) Key() interface{} {
	return iterator.internalIterator.Key()
}

func (iterator *Iterator) Value() interface{} {
	return iterator.internalIterator.Value()
}
