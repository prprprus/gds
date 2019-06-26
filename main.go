package main

import (
	"fmt"

	"github.com/prprprus/ds/singlylinkedlist"
	"github.com/prprprus/ds/util"
)

func main() {
	list := singlylinkedlist.New(1, 2, 3, 4, 5, 6, 7, 8, 9)

	list.Append(10, 11, 12)
	list.PreAppend(13, 14, 15, 16)

	v, ok := list.Get(7)
	fmt.Println(v, ok)

	err := list.Remove(0)
	fmt.Println(err)
	v, ok = list.Get(0)
	fmt.Println(v)

	fmt.Println(util.Sunday("12345", "15"))
	fmt.Println(util.Sunday("12345", "4"))

	_ = list.Swap(0, 2)
	v, ok = list.Get(0)
	fmt.Println(v)
	v, ok = list.Get(2)
	fmt.Println(v)
	// 16,15,14,1,2,3,4,5,6,7,8,9,10,11,12
	fmt.Println("--->")

	err = list.Insert(8, -4, -3, -1, -7, -9)
	fmt.Println(err)
	for i := 0; i < list.Size(); i++ {
		v, ok = list.Get(i)
		fmt.Println(v)
	}
	fmt.Println("--->")

	list.Set(11, -11)
	for i := 0; i < list.Size(); i++ {
		v, ok = list.Get(i)
		fmt.Println(v)
	}
	fmt.Println("--->")

	values := list.Values()
	fmt.Println(values)
	fmt.Println("--->")

	iterator := list.Iterator()
	for iterator.Next() {
		fmt.Println(iterator.Index(), iterator.Value())
	}
	fmt.Println("--->")

	fmt.Println(list.Values())
	data, err := list.ToJSON()
	fmt.Println(data)
	err = list.FromJSON(data)
	fmt.Println(list.Values())
}
