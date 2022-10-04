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
		fmt.Printf("Worker %d receive %c \n", id, n)
	}
}

func main() {
	channels := make([]chan int, 10)

	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	// 因为程序过快，所以这里 go func 还没执行完就结束了
	time.Sleep(time.Second)
}
