package main

import (
	"fmt"
)

/**
打印 slice 的各个属性
*/
func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n",
		s, len(s), cap(s))
}

func sliceOps() {

	fmt.Println("Create Slice")
	var s []int

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, i)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

}
