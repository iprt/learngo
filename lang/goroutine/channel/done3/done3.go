package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in   chan int
	done func()
}

func doWork(id int, w worker) {
	// 更简单的做法
	for n := range w.in {
		fmt.Printf("Worker %d Receive %c\n", id, n)
		w.done()
	}
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			// 函数式改造
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

func doneDemo() {
	//
	var wg = sync.WaitGroup{}
	wg.Add(20)

	workers := make([]worker, 10)

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()

}

func main() {
	doneDemo()
}
