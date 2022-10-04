package main

import (
	"fmt"
)

type worker struct {
	in   chan int
	done chan bool
}

func doWork(id int, w worker) {
	// 更简单的做法
	for n := range w.in {
		fmt.Printf("Worker %d Receive %c\n", id, n)
		go func() {
			// 阻塞的，需要等待人接收 ,所以需要继续改成异步的
			w.done <- true
		}()
	}
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w)
	return w
}

func doneDemo() {
	workers := make([]worker, 10)

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	for _, worker := range workers {
		<-worker.done
		<-worker.done
	}
}

func main() {
	doneDemo()
}
