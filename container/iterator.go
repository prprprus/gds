package container

// ValueIterator Interface
type ValueIterator interface {
	Next() bool
	Begin()
	Value() interface{}
}

// ReverseValueIterator Interface
type ReverseValueIterator interface {
	ValueIterator

	Prev() bool
	End()
}

// IndexIterator Interface
type IndexIterator interface {
	ValueIterator

	Index() int
}

// ReverseIndexIterator Interface
type ReverseIndexIterator interface {
	IndexIterator

	Prev() bool
	End()
}

// KeyIterator Interface
type KeyIterator interface {
	ValueIterator

	Key() interface{}
}

// ReverseKeyIterator Interface
type ReverseKeyIterator interface {
	KeyIterator

	Prev() bool
	End()
}
