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
