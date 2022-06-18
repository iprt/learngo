package main

import (
	"fmt"
	"io/ioutil"
)

// 分支操作

func grade(score int) string {
	g := ""

	switch {
	case score < 0 || score > 100:
		// 调用 panic 后会立刻停止执行当前函数的剩余代码
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {

	fmt.Println(grade(10))

	const filename = "abc.txt"

	// If "abc.txt" is not found,
	// please check what current directory is,
	// and change filename accordingly.
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("contents is %s \n", file)
	}

	// simplify
	if f, e := ioutil.ReadFile(filename); e != nil {
		fmt.Println(e)
	} else {
		fmt.Printf("content is %s\n", f)
	}

}
