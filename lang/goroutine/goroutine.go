package main

import (
	"fmt"
	"runtime"
	"time"
)

func test01() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				fmt.Println(i)
			}
		}(i)
	}

	time.Sleep(time.Second)
}

func test02() {
	var a [10]int
	fmt.Println("Running in", runtime.Version())
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

func main() {

	// test01()
	test02()
}
