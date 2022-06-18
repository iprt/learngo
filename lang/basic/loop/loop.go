package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
数字转为二进制
*/
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n = n / 2 {
		lsb := n % 2
		// int 转成字符串 strconv.Itoa
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(convertToBin(8))
	fmt.Println(convertToBin(10))
	fmt.Println(convertToBin(32480932))

	printFile("abc.txt")

	// string的另外一种写法
	s := `abc"d"
	kkkk
	123
	
	p`

	printFileContents(strings.NewReader(s))

	//forever()
}
