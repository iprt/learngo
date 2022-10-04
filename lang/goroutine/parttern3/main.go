package main

import (
	"fmt"
	"math/rand"
	"time"
)

func msgGen(name string, done chan struct{}) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {

			select {
			case <-time.After(time.Duration(rand.Intn(2000)) * time.Millisecond):
				c <- fmt.Sprintf("service %s message  %d", name, i)
			case <-done:
				fmt.Println("Cleaning up")
				time.Sleep(time.Second)
				fmt.Println("Cleaning up done")
				// 再发一个
				done <- struct{}{}
			}

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

func main() {
	done := make(chan struct{})
	m1 := msgGen("service1", done)
	for i := 0; i < 5; i++ {
		if m, ok := timeoutWait(m1, time.Second); ok {
			fmt.Println(m)
		} else {
			fmt.Println("timeout")
		}
	}
	done <- struct{}{}

	// 再收一个
	<-done

}
