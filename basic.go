package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 变量的声明
var aa = 33

var (
	ss = "kk"
	bb = true
)

func helloWorld() {
	fmt.Println("hello world")
}

func variableZeroValue() {
	var a int
	var s string
	var b bool
	fmt.Println(a, s, b)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

// 复数的表现形式
func euler() {
	fmt.Printf("%.3f\n",
		cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

//func calcTriangle(a int, b int) int {
func calcTriangle(a, b int) int {
	sqrt := math.Sqrt(float64(a*a + b*b))
	return int(sqrt)
}

// 常量的表达
func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

// 枚举的表达
func enums() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	helloWorld()
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()

	euler()

	triangle()

	consts()

	enums()
}
