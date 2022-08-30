package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/**
函数式编程
*/
func fibonacci() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

/**
为函数实现接口,传入进来就直接用
*/
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContent(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func testPrintFib() {

	/**
	对于type类型的变量的声明需要用 var param xxx
	*/
	var f intGen = fibonacci()

	printFileContent(f)
}

func testPrintFile() {
	file, err := os.Open("abc.txt")

	if err != nil {
		panic(err)
	}
	// fmt.Println(file)

	printFileContent(file)
}

func test() {
	f := fibonacci()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	testPrintFile()

	testPrintFib()
}

func main() {
	test()
	test()
}
