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
