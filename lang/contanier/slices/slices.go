package main

import (
	"fmt"
)

func updateSlice(s []int) {
	s[0] = 100
}

func study() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6] =", arr[2:6])
	fmt.Println("arr[:6] =", arr[:6])

	s1 := arr[2:]
	fmt.Println("s1 =", s1)
	s2 := arr[:]
	fmt.Println("s2 =", s2)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println("s1 = ", s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2
	fmt.Println("arr =", arr)
	s1 = arr[2:6]
	s2 = s1[3:5] // [s1[3], s1[4]]
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",
		s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n",
		s2, len(s2), cap(s2))

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)

	fmt.Println("s3, s4, s5 =", s3, s4, s5)
	// s4 and s5 no longer view arr.
	fmt.Println("arr =", arr)

	// Uncomment to run sliceOps demo.
	// If we see undefined: sliceOps
	// please try go run slices.go sliceops.go
	fmt.Println("Uncomment to see sliceOps demo")
	// sliceOps()
}

func reSlice() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr)
	arr1 := arr[2:8]
	fmt.Println("basic arr1 is: ", arr1)
	arr1[0] = 888
	fmt.Println("after modify 1 arr is", arr)
	fmt.Println("after modify 1 arr1 is", arr1)

	// 堆arr1 再切片
	arr2 := arr1[2:4]
	fmt.Println("basic arr2 is: ", arr2)
	arr2[0] = 999
	fmt.Println("after modify 2 arr is:", arr)
	fmt.Println("after modify 2 arr2 is :", arr2)
}

func reSlice2() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr)

	arr1 := arr[2:6]
	arr2 := arr[4:7]

	fmt.Println("arr1 is: ", arr1)
	fmt.Println("arr2 is: ", arr2)

	arr1[2] = 888
	arr2[0] = 999

	fmt.Println("arr is", arr)

	arr1 = append(arr1, 10086)
	fmt.Println("after append")
	fmt.Println("arr is:", arr)
	fmt.Println("arr1 is", arr1)
	fmt.Println("arr2 is", arr2)

}

/*
func main() {
	reSlice2()
}
*/
