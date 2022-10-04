package main

import (
	"fmt"
	"time"
)

//
// worker
//  @Description: channel 是可以作为参数传递的
//  @param c
//
func worker(id int, c chan int) {
	for {
		n := <-c
		// fmt.Println(n)
		fmt.Printf("Worker %d receive %d \n", id, n)
	}
}

func main() {
	c := make(chan int)

	// 必须要异步执行
	go worker(1, c)

	c <- 1
	c <- 2

	// 因为程序过快，所以这里 go func 还没执行完就结束了
	time.Sleep(time.Second)
}
