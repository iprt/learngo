package main

import (
	"fmt"
	"os"
)

func fib() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func listFiles(action func(f os.File)) {

}

func main() {

	next := fib()

	/**
	loop print
	*/
	for i := next(); i < 10000; i = next() {
		fmt.Println("value is ", i)
	}

}
