package main

import (
	"fmt"
	"math/rand"
	"time"
)

func msgGen(name string) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("service %s message  %d", name, i)
			i++
		}
	}()
	return c
}

func nonBlockingWait(c chan string) (string, bool) {
	select {
	case m := <-c:
		return m, true
	default:
		return "", false
	}
}

func timeoutWait(c chan string, timeout time.Duration) (string, bool) {
	select {
	case m := <-c:
		return m, true
	case <-time.After(timeout):
		return "", false
	}
}

func nbwDemo() {
	// handle
	m1 := msgGen("service1")
	m2 := msgGen("service2")
	for {
		fmt.Println(<-m1)
		if msg, ok := nonBlockingWait(m2); ok {
			fmt.Println(msg)
		}
	}
}
func timeoutDemo() {
	m1 := msgGen("service1")
	for {
		if msg, ok := timeoutWait(m1, 1*time.Second); ok {
			fmt.Println(msg)
		} else {
			fmt.Println("timeout")
		}
	}
}

func main() {

	// nbwDemo()
	timeoutDemo()
}
