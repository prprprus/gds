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
