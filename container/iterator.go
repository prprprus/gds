package container

// IndexIterator Interface
type IndexIterator interface {
	Next() bool
	Begin()
	Value() interface{}
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
	Next() bool
	Begin()
	Value() interface{}
	Key() interface{}
}

// ReverseKeyIterator Interface
type ReverseKeyIterator interface {
	KeyIterator

	Prev() bool
	End()
}
