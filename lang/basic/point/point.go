package main

import (
	"fmt"
)

func demo() {
	address := "Shanghai,China"

	ptr := &address
	// 打印指针的类型
	fmt.Printf("ptr type ：%T\n", ptr)
	// 打印指针的值，其实就是指针的地址
	fmt.Printf("ptr value: %p\n", ptr)

	var value = *ptr
	fmt.Printf("value type :%T\n", value)
	fmt.Printf("value : %s\n", value)
}

// &是取地址，*是取值

/**
思考下为什么不可以

在go语言中，只有值传递
*/
func exchange2(a, b *int) {
	a, b = b, a
}

func exchange(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

func main() {
	demo()

	a, b := 6, 8
	exchange(&a, &b)
	fmt.Printf("a = %d\n", a)
	fmt.Printf("b = %d\n", b)
}
