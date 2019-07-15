![](https://raw.githubusercontent.com/prprprus/picture/master/ds8.png)

![build status](https://travis-ci.org/prprprus/ds.svg?branch=master)
[![go report](https://goreportcard.com/badge/github.com/prprprus/ds)](https://goreportcard.com/report/github.com/prprprus/ds)
[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/prprprus/ds)
[![license](https://img.shields.io/badge/license-license-yellow.svg)](https://github.com/prprprus/ds/blob/master/LICENSE)

[中文文档](https://github.com/prprprus/ds/blob/master/README-zh.md)

# Introduction

Implement some data structures with go.

## Data Structures

- [Iterator]()
- [Container]()
    - [List]()
        - [SinglyLinkedList]()
        - [DoubleLinkedList]()
        - [ArrayList]()
    - [Stack]()
        - [LinkedListStack]()
        - [ArrayStack]()
    - [Queue]()
        - [LinkedListQueue]()
        - [ArrayQueue]()
    - [SkipList]()
    - [Map]()
        - [HashMap]()
        - [LinkedHashMap]()
        - [SkipMap]()
    - [Set]()
        - [HashSet]()
        - [LinkedHashSet]()
        - [SkipSet]()
- [Util]()
    - [Comparator]()

## Iterator

The package provides six iterators as following.

`ValueIterator` will traverse the value backwards.

```go
type ValueIterator interface {
	Next() bool
	Begin()
	Value() interface{}
}
```

`ReverseValueIterator` will traverse the value pair backwards or forwards.

```go
type ReverseValueIterator interface {
	ValueIterator

	Prev() bool
	End()
}
```

`IndexIterator` will traverse the index-value pair backwards.

```go
type IndexIterator interface {
	ValueIterator

	Index() int
}
```

`ReverseIndexIterator` will traverse the index-value pair backwards or forwards.

```go
type ReverseIndexIterator interface {
	IndexIterator

	Prev() bool
	End()
}
```

`KeyIterator` will traverse the key-value pair backwards.

```go
type KeyIterator interface {
	ValueIterator

	Key() interface{}
}
```

`ReverseKeyIterator` will traverse the key-value pair backwards or forwards.

```go
type ReverseKeyIterator interface {
	KeyIterator

	Prev() bool
	End()
}
```

Different data structures have different support for iterator as following.

![](https://raw.githubusercontent.com/prprprus/picture/master/ds1.png)

## Container

All data structures will implement the Container interface.

```go
type Container interface {
	Empty() bool
	Size() int
	Clear()
	Values() []interface{}
}
```

### List

List is an abstraction of a linear data structure, ordered and value repeatable.

Implements [Container]() interface.

```go
type List interface {
	Append(values ...interface{})
	Get(index int) (interface{}, error)
	Remove(index int) error
	Contains(values ...interface{}) bool
	Swap(i, j int) error
	Insert(index int, values ...interface{}) error
	Set(index int, value interface{}) error
	IndexOf(value interface{}) (int, error)

	container.Container
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
}
```

#### SinglyLinkedList

SinglyLinkedList is a singly linked list, the previous element points to the next element.

Implements [List]() Iterator, 