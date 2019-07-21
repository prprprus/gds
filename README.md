![](https://raw.githubusercontent.com/prprprus/picture/master/ds19.png)

![build status](https://travis-ci.org/prprprus/ds.svg?branch=master)
[![go report](https://goreportcard.com/badge/github.com/prprprus/ds)](https://goreportcard.com/report/github.com/prprprus/ds)
[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/prprprus/ds)
[![license](https://img.shields.io/badge/license-license-yellow.svg)](https://github.com/prprprus/ds/blob/master/LICENSE)
[![](https://img.shields.io/badge/CN-%E4%B8%AD%E6%96%87-%09%23ff2121.svg)](./README-zh.md)

# Introduction

Implement data structures with Go.

## Table Of Content

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
        - [DoubleLinkedList](#DoubleLinkedList)
        - [ArrayList](#ArrayList)
    - [Stack](#Stack)
        - [LinkedListStack](#LinkedListStack)
        - [ArrayStack](#ArrayStack)
    - [Queue](#Queue)
        - [LinkedListQueue](#LinkedListQueue)
        - [ArrayQueue](#ArrayQueue)
    - [SkipList](#SkipList)
    - [Map](#Map)
        - [HashMap](#HashMap)
        - [LinkedHashMap](#LinkedHashMap)
        - [SkipMap](#SkipMap)
    - [Set](#Set)
        - [HashSet](#HashSet)
        - [LinkedHashSet](#LinkedHashSet)
        - [SkipSet](#SkipSet)
- [Util](#Util)
    - [Comparator](#Comparator)
- [Benchmarking](#Benchmarking)

## Iterator

### ValueIterator

The package provides six iterators as following.

`ValueIterator` traverses the value backward.

```go
type ValueIterator interface {
	Next() bool
	Begin()
	Value() interface{}
}
```

### ReverseValueIterator

`ReverseValueIterator` can traverse values forward or backward.

```go
type ReverseValueIterator interface {
	ValueIterator

	Prev() bool
	End()
}
```

### IndexIterator

`IndexIterator` traverses the index-value pair backward.

```go
type IndexIterator interface {
	ValueIterator

	Index() int
}
```

### ReverseIndexIterator

`ReverseIndexIterator` can traverse the index-value pair forward or backward.

```go
type ReverseIndexIterator interface {
	IndexIterator

	Prev() bool
	End()
}
```

### KeyIterator

`KeyIterator` traverses the key-value pair backward.

```go
type KeyIterator interface {
	ValueIterator

	Key() interface{}
}
```

### ReverseKeyIterator

`ReverseKeyIterator` can traverse the key-value pair forward or backward.

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

`List` is ordered and value repeatable.

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

The current element of `SinglyLinkedList` points to the next element.

Implements [List](#list), [ValueIterator](#ValueIterator) and [IndexIterator](#indexIterator) interface.

<img src="https://raw.githubusercontent.com/prprprus/picture/master/ds13.png" alt="drawing" width="300"/>

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

#### DoubleLinkedList

The current and next elements of the `DoubleLinkedList` point to each other.

Implements [List](#list), [ValueIterator](#ValueIterator), [ReverseValueIterator](#ReverseValueIterator), [IndexIterator](#IndexIterator) and [ReverseIndexIterator](#ReverseIndexIterator) interface.

<img src="https://raw.githubusercontent.com/prprprus/picture/master/ds14.png" alt="drawing" width="450"/>

```go
package main

import (
	"fmt"

	"github.com/prprprus/ds/list/doublelinkedlist"
)

func main() {
	list := doublelinkedlist.New() // []
	list.Append(1)                 // [1]
	list.Append(2)                 // [1 2]
	list.Append(3)                 // [1 2 3]
	list.PreAppend(4)              // [4 1 2 3]
	_, _ = list.Get(0)             // 4, nil
	_, _ = list.Get(999)           // nil, ErrIndex
	_ = list.Remove(2)             // [4 1 3]
	_ = list.Contains()            // true
	_ = list.Contains(4, 1)        // true
	_ = list.Contains(4, 3)        // true
	_ = list.Contains(4, 1, 3, 5)  // false
	_ = list.Swap(0, 1)            // [1 4 3]
	_ = list.Insert(1, 5, 6, 7, 8) // [1 4 5 6 7 8 3]
	_ = list.Set(3, -1)            // [1 4 5 -1 7 8 3]

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
	it.End()
	for it.Prev() {
		fmt.Println(it.Index(), it.Value())
	}
	// output:
	// 6 3
	// 5 8
	// 4 7
	// 3 -1
	// 2 5
	// 1 4
	// 0 1

	_, _ = list.IndexOf(1)   // 0, nil
	_, _ = list.IndexOf(8)   // 5, nil
	_, _ = list.IndexOf(100) // -1, ErrIndexOf
	list.Reverse()           // [3 8 7 -1 5 4 1]
	_ = list.Empty()         // false
	_ = list.Size()          // 7
	_ = list.Values()        // [3 8 7 -1 5 4 1]
	list.Clear()             // []
}
```

#### ArrayList

`ArrayList` is a dynamic array that can be dynamically scaled based on capacity and number of elements.

Implements [List](#list), [ValueIterator](#ValueIterator), [ReverseValueIterator](#ReverseValueIterator), [IndexIterator](#IndexIterator) and [ReverseIndexIterator](#ReverseIndexIterator) interface.

<img src="https://raw.githubusercontent.com/prprprus/picture/master/ds17.png" alt="drawing" width="300"/>

```go
package main

import (
	"fmt"

	"github.com/prprprus/ds/list/arraylist"
)

func main() {
	list := arraylist.New()        // []
	list.Append(1)                 // [1]
	list.Append(2)                 // [1 2]
	list.Append(3)                 // [1 2 3]
	_, _ = list.Get(0)             // 1, nil
	_, _ = list.Get(999)           // nil, ErrIndex
	_ = list.Remove(2)             // [1 2]
	_ = list.Contains()            // true
	_ = list.Contains(1, 2)        // true
	_ = list.Contains(2)           // true
	_ = list.Contains(1, 2, 3)     // false
	_ = list.Swap(0, 1)            // [2 1]
	_ = list.Insert(1, 5, 6, 7, 8) // [2 1 5 6 7 8]
	_ = list.Set(3, -1)            // [2 1 5 -1 7 8]

	// iterator
	it := list.Iterator()
	it.Begin()
	for it.Next() {
		fmt.Println(it.Index(), it.Value())
	}
	// output:
	// 0 2
	// 1 1
	// 2 5
	// 3 -1
	// 4 7
	// 5 8
	it.End()
	for it.Prev() {
		fmt.Println(it.Index(), it.Value())
	}
	// output:
	// 	5 8
	// 4 7
	// 3 -1
	// 2 5
	// 1 1
	// 0 2

	_, _ = list.IndexOf(1)   // 1, nil
	_, _ = list.IndexOf(8)   // 5, nil
	_, _ = list.IndexOf(100) // -1, ErrIndexOf
	_ = list.Empty()         // false
	_ = list.Size()          // 6
	_ = list.Values()        // [2 1 5 -1 7 8]
	list.Clear()             // []
}
```

### Stack

`Stack` is a FILO data structure.

Implements [Container](#container) interface.

<img src="https://raw.githubusercontent.com/prprprus/picture/master/ds15.png" alt="drawing" width="500"/>

```go
type Stack interface {
	Push(value interface{})
	Pop() (interface{}, error)
	Peek() (interface{}, error)

	container.Container
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
}
```

#### LinkedListStack

`LinkedListStack` is a stack based on [SinglyLinkedList](#SinglyLinkedList).

Implements [Stack](#Stack), [ValueIterator](#ValueIterator) and [IndexIterator](#IndexIterator) interface.

```go
package main

import (
	"fmt"

	"github.com/prprprus/ds/stack/linkedliststack"
)

func main() {
	stack := linkedliststack.New() // []
	stack.Push(1)                  // [1]
	stack.Push(2)                  // [2 1]
	stack.Push(3)                  // [3 2 1]

	// iterator
	it := stack.Iterator()
	it.Begin()
	for it.Next() {
		fmt.Println(it.Index(), it.Value())
	}
	// output:
	// 0 3
	// 1 2
	// 2 1

	_, _ = stack.Peek() // 3, nil
	_, _ = stack.Pop()  // 3, nil
	_, _ = stack.Peek() // 2, nil
	_ = stack.Empty()   // false
	_ = stack.Size()    // 2
	_ = stack.Values()  // [2 1]
	stack.Clear()       // []
}
```

#### ArrayStack

`ArrayStack` is a stack based on [ArrayList](#ArrayList).

Implements [Stack](#Stack), [ValueIterator](#ValueIterator), [ReverseValueIterator](#ReverseValueIterator), [IndexIterator](#IndexIterator) and [ReverseIndexIterator](#ReverseIndexIterator) interface.

```go
package main

import (
	"fmt"

	"github.com/prprprus/ds/stack/arraystack"
)

func main() {
	stack := arraystack.New() // []
	stack.Push(1)             // [1]
	stack.Push(2)             // [2 1]
	stack.Push(3)             // [3 2 1]

	// iterator
	it := stack.Iterator()
	it.Begin()
	for it.Next() {
		fmt.Println(it.Index(), it.Value())
	}
	// output:
	// 0 3
	// 1 2
	// 2 1

	_, _ = stack.Peek() // 3, nil
	_, _ = stack.Pop()  // 3, nil
	_, _ = stack.Peek() // 2, nil
	_ = stack.Empty()   // false
	_ = stack.Size()    // 2
	_ = stack.Values()  // [2 1]
	stack.Clear()       // []
}
```

### Queue

`Queue` is a FIFO data structure.

Implements [Container](#container) interface.

<img src="https://raw.githubusercontent.com/prprprus/picture/master/ds16.png" alt="drawing" width="300"/>

```go
type Queue interface {
	Put(value interface{})
	Get() (interface{}, error)

	container.Container
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
}
```

#### LinkedListQueue

`LinkedListQueue` is a stack based on [SinglyLinkedList](#SinglyLinkedList).

Implements [Queue](#Queue), [ValueIterator](#ValueIterator) and [IndexIterator](#IndexIterator) interface.

```go
package main

import (
	"fmt"

	"github.com/prprprus/ds/queue/linkedlistqueue"
)

func main() {
	queue := linkedlistqueue.New() // []
	queue.Put(1)                   // [1]
	queue.Put(2)                   // [1 2]
	queue.Put(3)                   // [1 2 3]
	queue.Put(4)                   // [1 2 3 4]

	// iterator
	it := queue.Iterator()
	it.Begin()
	for it.Next() {
		fmt.Println(it.Index(), it.Value())
	}
	// output:
	// 0 1
	// 1 2
	// 2 3
	// 3 4

	_, _ = queue.Get() // 1, nil
	_, _ = queue.Get() // 2, nil
	_ = queue.Empty()  // false
	_ = queue.Size()   // 2
	_ = queue.Values() // [3 4]
	queue.Clear()      // []
}
```

#### ArrayQueue

`ArrayQueue` is a stack based on [ArrayList](#ArrayList).

Implements [Queue](#Queue), [ValueIterator](#ValueIterator) and [IndexIterator](#IndexIterator) interface.

```go
package main

import (
	"fmt"

	"github.com/prprprus/ds/queue/arrayqueue"
)

func main() {
	queue := arrayqueue.New() // []
	queue.Put(1)              // [1]
	queue.Put(2)              // [1 2]
	queue.Put(3)              // [1 2 3]
	queue.Put(4)              // [1 2 3 4]

	// iterator
	it := queue.Iterator()
	it.Begin()
	for it.Next() {
		fmt.Println(it.Index(), it.Value())
	}
	// output:
	// 0 1
	// 1 2
	// 2 3
	// 3 4

	_, _ = queue.Get() // 1, nil
	_, _ = queue.Get() // 2, nil
	_ = queue.Empty()  // false
	_ = queue.Size()   // 2
	_ = queue.Values() // [3 4]
	queue.Clear()      // []
}
```

### SkipList

`SkipList` is a random data structure with performance comparable to that of red-black trees. It should be noted that the keys must be comparable types and element will be sorted by keys.

Implements [Container](#container), [ValueIterator](#ValueIterator) and [KeyIterator](#KeyIterator) interface.

<img src="https://raw.githubusercontent.com/prprprus/picture/master/ds2.png" alt="drawing" width="550"/>

```go
package main

import (
	"fmt"

	"github.com/prprprus/ds/skiplist"
	"github.com/prprprus/ds/util"
)

func main() {
	skiplist := skiplist.New(util.IntComparator) // []
	skiplist.Set(1, "a")                         // [{1: "a"}]
	skiplist.Set(2, "b")                         // [{1: "a"} {2: "b"}]
	skiplist.Set(3, "c")                         // [{1: "a"} {2: "b"} {3: "c"}]
	skiplist.Set(4, "d")                         // [{1: "a"} {2: "b"} {3: "c"} {4: "d"}]

	// iterator
	it := skiplist.Iterator()
	it.Begin()
	for it.Next() {
		fmt.Println(it.Key(), it.Value())
	}
	// output:
	// 1 a
	// 2 b
	// 3 c
	// 4 d

	_ = skiplist.Exists(1) // true
	_ = skiplist.Exists(9) // false
	_, _ = skiplist.Get(1) // "a", nil
	_, _ = skiplist.Get(3) // "c", nil
	_ = skiplist.Remove(2) // nil
	_ = skiplist.Keys()    // [1 3 4]
	_ = skiplist.Empty()   // false
	_ = skiplist.Size()    // 3
	_ = skiplist.Values()  // [a c d]
	skiplist.Clear()       // []
}
```

### Map

`Map` stores key-value pairs with excellent operational performance. It should be noted that the keys must be comparable types.

Implements [Container](#container) interface.

```go
type Map interface {
	Put(key, value interface{})
	Get(key interface{}) (interface{}, error)
	Remove(key interface{})

	container.Container
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
}
```

#### HashMap

`HashMap` is a map based on hash table.

Implements [Map](#Map) interface.

<img src="https://raw.githubusercontent.com/prprprus/picture/master/ds18.png" alt="drawing" width="400"/>

```go
package main

import (
	"github.com/prprprus/ds/maps/hashmap"
)

func main() {
	m := hashmap.New() // []
	m.Put(1, "a")      // [{1: "a"}]
	m.Put(2, "b")      // [{1: "a"} {2: "b"}]
	m.Put(3, "c")      // [{1: "a"} {2: "b"} {3: "c"}]
	m.Put(4, "d")      // [{1: "a"} {2: "b"} {3: "c"} {4: "d"}]
	_ = m.Keys()       // [1 2 3 4] (Note: order of random)
	_, _ = m.Get(1)    // "a", nil
	_, _ = m.Get(3)    // "c", nil
	_ = m.Remove(2)    // nil
	_ = m.Keys()       // [1 3 4]
	_ = m.Empty()      // false
	_ = m.Size()       // 3
	_ = m.Values()     // [a c d] (Note: order of random)
	m.Clear()          // []
}
```

#### LinkedHashMap

`LinkedHashMap` is a map based on hash table and [DoubleLinkedList](#DoubleLinkedList), it provides ordered key-value pairs.

Implements [Map](#Map), [ValueIterator](#ValueIterator), [ReverseValueIterator](#ReverseValueIterator), [KeyIterator](#KeyIterator) and [ReverseKeyIterator](#ReverseKeyIterator) interface.

```go
package main

import (
	"fmt"

	"github.com/prprprus/ds/maps/linkedhashmap"
)

func main() {
	m := linkedhashmap.New() // []
	m.Put(1, "a")            // [{1: "a"}]
	m.Put(2, "b")            // [{1: "a"} {2: "b"}]
	m.Put(3, "c")            // [{1: "a"} {2: "b"} {3: "c"}]
	m.Put(4, "d")            // [{1: "a"} {2: "b"} {3: "c"} {4: "d"}]

	// iterator
	it := m.Iterator()
	it.Begin()
	for it.Next() {
		fmt.Println(it.Key(), it.Value())
	}
	// output:
	// 1 a
	// 2 b
	// 3 c
	// 4 d
	it.End()
	for it.Prev() {
		fmt.Println(it.Key(), it.Value())
	}
	// output:
	// 4 d
	// 3 c
	// 2 b
	// 1 a

	_ = m.Keys()    // [1 2 3 4]
	_, _ = m.Get(1) // "a", nil
	_, _ = m.Get(3) // "c", nil
	_ = m.Remove(2) // nil
	_ = m.Keys()    // [1 3 4]
	_ = m.Empty()   // false
	_ = m.Size()    // 3
	_ = m.Values()  // [a c d]
	m.Clear()       // []
}
```

#### SkipMap

`SkipMap` is a map based on [SkipList](#SkipList).

Implements [Map](#Map), [ValueIterator](#ValueIterator) and [KeyIterator](#KeyIterator).

```go
package main

import (
	"fmt"

	"github.com/prprprus/ds/maps/linkedhashmap"
)

func main() {
	m := linkedhashmap.New() // []
	m.Put(1, "a")            // [{1: "a"}]
	m.Put(2, "b")            // [{1: "a"} {2: "b"}]
	m.Put(3, "c")            // [{1: "a"} {2: "b"} {3: "c"}]
	m.Put(4, "d")            // [{1: "a"} {2: "b"} {3: "c"} {4: "d"}]

	// iterator
	it := m.Iterator()
	it.Begin()
	for it.Next() {
		fmt.Println(it.Key(), it.Value())
	}
	// output:
	// 1 a
	// 2 b
	// 3 c
	// 4 d

	_ = m.Keys()    // [1 2 3 4]
	_, _ = m.Get(1) // "a", nil
	_, _ = m.Get(3) // "c", nil
	_ = m.Remove(2) // nil
	_ = m.Keys()    // [1 3 4]
	_ = m.Empty()   // false
	_ = m.Size()    // 3
	_ = m.Values()  // [a c d]
	m.Clear()       // []
}
```

### Set

`Set` is used to store non-repeating values, usually with good operational performance.

Implements [Container](#container) interface.

```go
type Set interface {
	Add(values ...interface{})
	Remove(values ...interface{}) error
	Contains(values ...interface{}) bool

	container.Container
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
}
```

#### HashSet

`HashSet` is a set based on hash table.

Implements [Set](#Set) interface.

```go
package main

import (
	"github.com/prprprus/ds/set/hashset"
)

func main() {
	s := hashset.New()      // []
	s.Add(1)                // [1]
	s.Add(2)                // [1 2]
	s.Add(3)                // [1 2 3]
	_ = s.Contains()        // true
	_ = s.Contains(1, 2, 3) // true
	_ = s.Contains(1, 3)    // true
	_ = s.Contains(2, 3, 4) // false
	_ = s.Remove(2)         // nil
	_ = s.Empty()           // false
	_ = s.Size()            // 2
	_ = s.Values()          // [1 3]
	s.Clear()               // []
}
```

#### LinkedHashSet

`LinkedHashSet` is a set based on hash table and [DoubleLinkedList](#DoubleLinkedList), it provides ordered value.

Implements [Set](#Set), [ValueIterator](#valueIterator) and [ReverseValueIterator](#reverseValueIterator) interface.

```go
package main

import (
	"fmt"

	"github.com/prprprus/ds/set/linkedhashset"
)

func main() {
	s := linkedhashset.New() // []
	s.Add(1)                 // [1]
	s.Add(2)                 // [1 2]
	s.Add(3)                 // [1 2 3]

	// iterator
	it := s.Iterator()
	it.Begin()
	for it.Next() {
		fmt.Println(it.Value())
	}
	// output:
	// 1
	// 2
	// 3
	it.End()
	for it.Prev() {
		fmt.Println(it.Value())
	}
	// output:
	// 3
	// 2
	// 1

	_ = s.Contains()        // true
	_ = s.Contains(1, 2, 3) // true
	_ = s.Contains(1, 3)    // true
	_ = s.Contains(2, 3, 4) // false
	_ = s.Remove(2)         // nil
	_ = s.Empty()           // false
	_ = s.Size()            // 2
	_ = s.Values()          // [1 3]
	s.Clear()               // []
}
```

#### SkipSet

`SkipSet` is a set based on [SkipList](#SkipList).

Implements [Set](#Set) and [ValueIterator](#valueIterator) interface.

```go
package main

import (
	"fmt"

	"github.com/prprprus/ds/set/skipset"
	"github.com/prprprus/ds/util"
)

func main() {
	s := skipset.New(util.IntComparator) // []
	s.Add(1)                             // [1]
	s.Add(2)                             // [1 2]
	s.Add(3)                             // [1 2 3]

	// iterator
	it := s.Iterator()
	it.Begin()
	for it.Next() {
		fmt.Println(it.Value())
	}
	// output:
	// 1
	// 2
	// 3

	_ = s.Contains()        // true
	_ = s.Contains(1, 2, 3) // true
	_ = s.Contains(1, 3)    // true
	_ = s.Contains(2, 3, 4) // false
	_ = s.Remove(2)         // nil
	_ = s.Empty()           // false
	_ = s.Size()            // 2
	_ = s.Values()          // [1 3]
	s.Clear()               // []
}
```

## Util

Contains some helper functions.

### Comparator

`Comparator` provides the following built-in type of comparator.

```go
func IntComparator(a, b interface{}) int

func Int8Comparator(a, b interface{}) int

func Int16Comparator(a, b interface{}) int

func Int32Comparator(a, b interface{}) int

func Int64Comparator(a, b interface{}) int

func UIntComparator(a, b interface{}) int

func UInt8Comparator(a, b interface{}) int

func UInt16Comparator(a, b interface{}) int

func UInt32Comparator(a, b interface{}) int

func UInt64Comparator(a, b interface{}) int

func Float32Comparator(a, b interface{}) int

func Float64Comparator(a, b interface{}) int

func ByteComparator(a, b interface{}) int

func RuneComparator(a, b interface{}) int

func StringComparator(a, b interface{}) int
```

The meaning of the return value is as follows.

```
-1 => a < b
0  => a == b
1  => a > b
```

For custom types, you can also create a corresponding comparator.

```go
package main

import (
    "fmt"

    "github.com/prprprus/ds/set/skipset"
)

type People struct {
    name string
    age  int
}

func AgeComparator(a, b interface{}) int {
    c1 := a.(People)
    c2 := b.(People)
    switch {
	case c1.age < c2.age:
		return -1
	case c1.age > c2.age:
		return 1
	default:
		return 0
	}
}

func main() {
    s := skipset.New(AgeComparator)
    s.Add(People{"Wade", 35})
    s.Add(People{"Simon", 32})
    s.Add(People{"yiyi", 22})

    fmt.Println(s.Values()) // [{"yiyi", 22}, {"Simon", 32}, {"Wade", 35}]
}
```

## Benchmarking

![]()