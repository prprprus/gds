package main

import (
	"fmt"

	"github.com/prprprus/ds/skiplist"
)

func main() {
	// ---> singly linked list

	// list := singlylinkedlist.New(1, 2, 3, 4, 5, 6, 7, 8, 9)

	// list.Append(10, 11, 12)
	// list.PreAppend(13, 14, 15, 16)

	// v, ok := list.Get(7)
	// fmt.Println(v, ok)

	// err := list.Remove(0)
	// fmt.Println(err)
	// v, ok = list.Get(0)
	// fmt.Println(v)

	// fmt.Println(util.Sunday("12345", "15"))
	// fmt.Println(util.Sunday("12345", "4"))

	// _ = list.Swap(0, 2)
	// v, ok = list.Get(0)
	// fmt.Println(v)
	// v, ok = list.Get(2)
	// fmt.Println(v)
	// // 16,15,14,1,2,3,4,5,6,7,8,9,10,11,12
	// fmt.Println("--->")

	// err = list.Insert(8, -4, -3, -1, -7, -9)
	// fmt.Println(err)
	// for i := 0; i < list.Size(); i++ {
	// 	v, ok = list.Get(i)
	// 	fmt.Println(v)
	// }
	// fmt.Println("--->")

	// list.Set(11, -11)
	// for i := 0; i < list.Size(); i++ {
	// 	v, ok = list.Get(i)
	// 	fmt.Println(v)
	// }
	// fmt.Println("--->")

	// values := list.Values()
	// fmt.Println(values)
	// fmt.Println("--->")

	// iterator := list.Iterator()
	// for iterator.Next() {
	// 	fmt.Println(iterator.Index(), iterator.Value())
	// }
	// fmt.Println("--->")

	// fmt.Println(list.Values())
	// data, err := list.ToJSON()
	// fmt.Println(data)
	// err = list.FromJSON(data)
	// fmt.Println(list.Values())

	// ---> double linked list

	// list := doublelinkedlist.New(1, 2, 3, 4, 5, 6, 7, 8, 9)
	// iterator := list.Iterator()
	// iterator.Begin()
	// for iterator.Next() {
	// 	fmt.Println(iterator.Index(), iterator.Value())
	// }
	// fmt.Println("--->")
	// iterator.End()
	// for iterator.Prev() {
	// 	fmt.Println(iterator.Index(), iterator.Value())
	// }

	// ---> skip list
	skiplist := skiplist.New()
	skiplist.Set(777)
	skiplist.Set(999)
	skiplist.Set(99)
	skiplist.Set(1)
	skiplist.Set(12)
	skiplist.Set(3)
	skiplist.Set(42)
	skiplist.Set(9)
	skiplist.Set(15)
	skiplist.Set(6)
	skiplist.Set(79)
	skiplist.Set(18)
	skiplist.Set(63)
	skiplist.Set(81)
	skiplist.Set(-1)
	skiplist.Show()
	fmt.Println("size:", skiplist.Size)
	fmt.Println(skiplist.Get(777))
	fmt.Println(skiplist.Get(42))
	fmt.Println(skiplist.Get(3))
	fmt.Println(skiplist.Get(18))
	fmt.Println(skiplist.Get(15))
	fmt.Println(skiplist.Get(99))
	fmt.Println(skiplist.Get(999))

	fmt.Println(skiplist.Get(11))
	fmt.Println(skiplist.Get(111))
	fmt.Println(skiplist.Get(1111))
	fmt.Println(skiplist.Get(1278))
	fmt.Println(skiplist.Get(11178))
}