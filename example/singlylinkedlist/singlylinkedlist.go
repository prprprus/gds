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

	_, _ = list.IndexOf(1)   // 0, nil
	_, _ = list.IndexOf(8)   // 5, nil
	_, _ = list.IndexOf(100) // -1, ErrIndexOf
	list.Reverse()           // [3, 8, 7, -1, 5, 4, 1]
	_ = list.Empty()         // false
	_ = list.Size()          // 7
	list.Clear()             // []
}
