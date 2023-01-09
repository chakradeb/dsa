package main

import (
	"fmt"
	"github.com/chakradeb/dsa/array"
)

func main() {
	arr := array.NewArray()
	arr.Add(5)
	arr.Add(9)
	arr.Add(3)
	fmt.Println("index 1 is:", arr.Get(1))
	fmt.Println("last element removed:", arr.RemoveLast())
	arr.Add(4)
	arr.Add(7)
	arr.Add(1)
	arr.Add(6)
	fmt.Println(arr)
	fmt.Println("removed index 2 was:", arr.Remove(2))
	fmt.Println(arr)
}
