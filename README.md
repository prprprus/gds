![](https://raw.githubusercontent.com/prprprus/picture/master/ds8.png)

![build status](https://travis-ci.org/prprprus/ds.svg?branch=master)
[![go report](https://goreportcard.com/badge/github.com/prprprus/ds)](https://goreportcard.com/report/github.com/prprprus/ds)
[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/prprprus/ds)
[![license](https://img.shields.io/badge/license-license-yellow.svg)](https://github.com/prprprus/ds/blob/master/LICENSE)

[中文文档](https://github.com/prprprus/ds/blob/master/README-zh.md)

# Introduction

Implement some data structures with go.

## Data Structures

- [Iterator](#iterator)
    - [ValueIterator](#valueIterator)
    - [ReverseValueIterator](#reverseValueIterator)
    - [IndexIterator](#indexIterator)
    - [ReverseIndexIterator](#reverseIndexIterator)
    - [KeyIterator](#keyIterator)
    - [ReverseKeyIterator](#reverseKeyIterator)
- [Container](#container)
    - [List](#list)
        - [SinglyLinkedList](#singlylinkedlist)
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

### ValueIterator

The package provides six iterators as following.

`ValueIterator` will traverse the value backwards.

```go
type ValueIterator interface {
	Next() bool
	Begin()
	Value() interface{}
}
```

### ReverseValueIterator

`ReverseValueIterator` will traverse the value pair backwards or forwards.

```go
type ReverseValueIterator interface {
	ValueIterator

	Prev() bool
	End()
}
```

### IndexIterator

`IndexIterator` will traverse the index-value pair backwards.

```go
type IndexIterator interface {
	ValueIterator

	Index() int
}
```

### ReverseIndexIterator

`ReverseIndexIterator` will traverse the index-value pair backwards or forwards.

```go
type ReverseIndexIterator interface {
	IndexIterator

	Prev() bool
	End()
}
```

### KeyIterator

`KeyIterator` will traverse the key-value pair backwards.

```go
type KeyIterator interface {
	ValueIterator

	Key() interface{}
}
```

### ReverseKeyIterator

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

Implements [Container](#container) interface.

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

Implements [List](#list), [IndexIterator](#indexIterator).

```go
package main

import (
	"fmt"

	"github.com/prprprus/ds/list/singlylinkedlist"
)

func main() {
	list := singlylinkedlist.New() // []
	list.Append(1)                 // [1]
	list.Append(2)                 // [1, 2]
	list.Append(3)                 // [1, 2, 3]
	list.PreAppend(4)              // [4, 1, 2, 3]
	_, _ = list.Get(0)             // 4, nil
	_, _ = list.Get(999)           // nil, ErrIndex
	_ = list.Remove(2)             // [4, 1, 3]
	_ = list.Contains()            // true
	_ = list.Contains(4, 1)        // true
	_ = list.Contains(4, 3)        // true
	_ = list.Contains(4, 1, 3, 5)  // false
	_ = list.Swap(0, 1)            // [1, 4, 3]
	_ = list.Insert(1, 5, 6, 7, 8) // [1, 4, 5, 6, 7, 8, 3]
	_ = list.Set(3, -1)            // [1, 4, 5, -1, 7, 8, 3]

	// iterator
	it := list.Iterator()
	it.Begin()
	for it.Next() {
		fmt.Println(it.Index(), it.Value())
	}
	// output:
	// 0 1
	// 1 4
	// 2 5
	// 3 -1
	// 4 7
	// 5 8
	// 6 3

	_, _ = list.IndexOf(1)   // 0, nil
	_, _ = list.IndexOf(8)   // 5, nil
	_, _ = list.IndexOf(100) // -1, ErrIndexOf
	list.Reverse()           // [3, 8, 7, -1, 5, 4, 1]
	_ = list.Empty()         // false
	_ = list.Size()          // 7
	_ = list.Values()        // [3 8 7 -1 5 4 1]
	list.Clear()             // []
}
```