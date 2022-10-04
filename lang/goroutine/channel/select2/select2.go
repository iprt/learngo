// Package main
// @Description: 对比 select.go 进行修改
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
		// 模拟消费延迟
		time.Sleep(time.Second)
		fmt.Printf("Worker %d receive %d \n", id, n)
	}
}

func createWorker(id int) chan int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {

	// 产生数据
	var c1, c2 = generator(), generator()

	// 消费者
	worker := createWorker(0)

	var values []int

	tm := time.After(10 * time.Second)

	// 每隔1秒钟送一个值过来
	tick := time.Tick(time.Second)

	for {
		// 利用 nil channel的特点
		var activeWorker chan int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			// 取出第一个元素
			activeValue = values[0]
		}
		select {
		// 下面3个case是系统的调度
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			// 重置
			values = values[1:]

		// 下面3个case是系统的监控
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("queue len ", len(values))
		case <-tm:
			fmt.Println("bye")
			return
		}
	}

}
