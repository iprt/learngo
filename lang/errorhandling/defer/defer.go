package main

import (
	"bufio"
	"fmt"
	"os"
)

func fib() func() int {
	var a, b = 1, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)

	panic("error occurred")

	fmt.Println(4)
}

func tryDefer2() {
	for i := 0; i < 100; i++ {
		// 更加深刻的理解defer式栈
		defer fmt.Println(i)
		if i == 30 {
			panic("error occurred")
		}
	}
}

/*
真正的对资源的操作
*/
func writeFile(filename string) {
	file, err := os.Create(filename)

	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	// 这种写法一般式想到的时候就可以写上去了
	defer writer.Flush()

	f := fib()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	// tryDefer()
	// writeFile("fib.txt")

	tryDefer2()

}
