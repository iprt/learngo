package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()

	c <- 1
	c <- 2

	// 因为程序过快，所以这里 go func 还没执行完就结束了
	time.Sleep(time.Second)
}
