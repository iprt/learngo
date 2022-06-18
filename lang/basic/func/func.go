package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil

	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported option %s", op)
	}
}

func div(a, b int) (m, n int) {
	return a / b, a * b
}

func apply(op func(int, int) int, a, b int) int {
	// 反射打印内容
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+
		"(%d, %d)\n", opName, a, b)

	return op(a, b)
}

func sum(numbers ...int) int {
	r := 0
	for i := range numbers {
		r += numbers[i]
	}
	return r
}

func swap(a, b int) (x, y int) {
	return b, a
}

func swapp(a, b *int) {
	fmt.Printf("a变量本身的地址=%p\n", &a)
	fmt.Printf("b变量本身的地址=%p\n", &b)
	fmt.Printf("a交换前的x赋予的地址=%p\n", a)
	fmt.Printf("b交换前的y赋予的地址=%p\n", b)
	fmt.Printf("a交换前的值=%d\n", *a)
	fmt.Printf("b交换前的值=%d\n", *b)
	*b, *a = *a, *b
	fmt.Printf("a变量本身的地址=%p\n", &a)
	fmt.Printf("b变量本身的地址=%p\n", &b)
	fmt.Printf("a交换后的x赋予的地址=%p\n", a)
	fmt.Printf("b交换后的y赋予的地址=%p\n", b)
	fmt.Printf("a交换后的值=%d\n", *a)
	fmt.Printf("b交换后的值=%d\n", *b)
}

func main() {
	if r, e := eval(10, 20, "+"); e != nil {
		fmt.Println("Error", e)
	} else {
		fmt.Printf("a + b = %d\n", r)
	}

	var _, n int = div(10, 10)
	fmt.Println(n)

	opv := apply(func(i int, i2 int) int {
		return i + i2
	}, 3, 4)
	fmt.Println(opv)

	fmt.Printf("sum 1 2 3 4 5 = %d\n", sum(1, 2, 3, 4, 5))

	x, y := 1, 2
	swapp(&x, &y)
}
