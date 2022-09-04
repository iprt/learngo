package main

import (
	"fmt"
	"time"
)

func worker(ch chan int) {
	for {
		i := <-ch
		fmt.Println("channel demo0 :", i)
	}
}

func channelDemo() {
	ch := make(chan int)
	go worker(ch)
	ch <- 1
	ch <- 2
	ch <- 3
	time.Sleep(time.Second)
}
func main() {
	channelDemo()
}
