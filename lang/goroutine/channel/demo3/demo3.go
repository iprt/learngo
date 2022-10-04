package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	// for {
	// 	n, ok := <-c
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Printf("Worker %d Receive %d\n", id, n)
	// }

	// 更简单的做法
	for n := range c {
		fmt.Printf("Worker %d Receive %d\n", id, n)
	}
}

//
// createWorker
//  @Description: 从 worker 变成 createWorker
//  @param id
//  @return chan
//
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func channelDemo() {
	channels := make([]chan<- int, 10)

	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	// 因为程序过快，所以这里 go func 还没执行完就结束了
	time.Sleep(time.Second)
}

func bufferedChannel() {
	c := make(chan int, 3)

	go worker(0, c)

	c <- 1
	c <- 2
	c <- 3
	c <- 4

	time.Sleep(time.Second)

}

func channelClose() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4

	// 关闭channel 注意： worker 会一直接收到  0 (即 int的默认值)
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	// channelDemo()

	// bufferedChannel()

	channelClose()
}
