package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			// 随机 1500 ms 内生成一个数据
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d receive %d \n", id, n)
	}
}

func createWorker(id int) chan int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func demo() {
	var c1, c2 = generator(), generator()

	// 消费者
	worker := createWorker(0)

	n := 0
	hasValue := false

	for {
		// 利用 nil channel的特点
		var activeWorker chan int
		if hasValue {
			activeWorker = worker
		}
		select {
		case n = <-c1:
			hasValue = true
		case n = <-c2:
			hasValue = true
		case activeWorker <- n:
			// 重置
			hasValue = false
		}
	}
}

func main() {

	demo()

}
