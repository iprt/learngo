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

func fanIn(c1, c2 chan string) chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-c1
		}
	}()

	go func() {
		for {
			c <- <-c2
		}
	}()
	return c
}
func fanInMore(chs ...chan string) chan string {
	c := make(chan string)
	for _, ch := range chs {
		// chCopy := ch
		// go func() {
		// 	for {
		// 		c <- <-chCopy
		// 	}
		// }()

		// 函数式解决传参的问题
		go func(in chan string) {
			for {
				c <- <-in
			}
		}(ch)
	}
	return c
}

func fanInBySelect(c1, c2 chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case msg := <-c1:
				c <- msg
			case msg := <-c2:
				c <- msg
			}
		}
	}()
	return c
}

func main() {
	// handle
	m1 := msgGen("service1")
	m2 := msgGen("service2")

	// m := fanIn(m1, m2)
	// m := fanInBySelect(m1, m2)
	m := fanInMore(m1, m2)
	for {
		fmt.Println(<-m)
	}

}
