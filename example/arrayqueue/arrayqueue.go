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
