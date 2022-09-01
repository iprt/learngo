package main

import (
	"fmt"
)

func main() {

	defer func() {
		err := recover()
		// if err != nil {
		// 	fmt.Println(err)
		// }

		if e, ok := err.(error); ok {
			fmt.Println("this is an error", e)
		} else {
			// panic(e)
			fmt.Println("I don't know how to do ,", err)
		}

	}()

	// panic(errors.New("test recover"))

	panic(123)
}
