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
