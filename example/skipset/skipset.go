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
