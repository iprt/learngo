package main

import (
	"fmt"
	"time"
)

func subscriber(id int, c <-chan int) {
	for {
		read := <-c
		fmt.Printf("subscriber %d received %c\n", id, read)
	}
}

// 创建消费者， 返回 channel 返回的 chan 是往里面写东西的
func createSubscriber(id int) chan<- int {
	ch := make(chan int)
	go subscriber(id, ch)
	return ch
}

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

func channelDemo1() {
	ch := make(chan int)
	go subscriber(0, ch)

	go func() {
		for i := 0; i < 20; i++ {
			ch <- 'A' + i
		}
	}()

	time.Sleep(time.Second)
}

func channelDemo2() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createSubscriber(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func main() {

	channelDemo1()

}
