package main

import (
	"bufio"
	"fmt"
	"os"
)

func fib() func() int {
	var a, b = 1, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)

	panic("error occurred")

	fmt.Println(4)
}

func tryDefer2() {
	for i := 0; i < 100; i++ {
		// 更加深刻的理解defer式栈
		defer fmt.Println(i)
		if i == 30 {
			panic("error occurred")
		}
	}
}

/*
真正的对资源的操作
*/
func writeFile(filename string) {
	file, err := os.Create(filename)

	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	// 这种写法一般式想到的时候就可以写上去了
	defer writer.Flush()

	f := fib()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

/**
错误处理的示例
*/

/*
func openFileNolog(name string, flag int, perm FileMode) (*File, error) {
	if name == "" {
		return nil, &PathError{Op: "open", Path: name, Err: syscall.ENOENT}
	}
	r, errf := openFile(name, flag, perm)
	if errf == nil {
		return r, nil
	}
	r, errd := openDir(name)
	if errd == nil {
		if flag&O_WRONLY != 0 || flag&O_RDWR != 0 {
			r.Close()
			return nil, &PathError{Op: "open", Path: name, Err: syscall.EISDIR}
		}
		return r, nil
	}
	return nil, &PathError{Op: "open", Path: name, Err: errf}
}


看看 error 到底是

*/
func writeFile2(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	// if err != nil {
	// 	// 原来的写法：panic(err)
	// 	fmt.Println("file already exists")
	// 	return
	// }

	if err != nil {
		if pathError, ok := err.(*os.PathError); ok {
			fmt.Println("file already exists: ", pathError.Path)
		} else {
			// fmt.Println("other error")
			panic(err)
		}
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	// tryDefer()

	// tryDefer2()

	// writeFile("fib.txt")
	writeFile2("fib.txt")

}
