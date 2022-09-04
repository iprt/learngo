package main

import (
	"fmt"
	"time"
)

func channelDemo0() {
	ch := make(chan int)
	go func() {
		for {
			i := <-ch
			fmt.Println("channel demo0 :", i)
		}
	}()
	ch <- 1
	ch <- 2
	ch <- 3
	time.Sleep(time.Second)
}

func main() {
	channelDemo0()
}
