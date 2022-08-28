package main

import (
	"fmt"
	"io/ioutil"
	"iproute.org/learngo/lang/interface/downloader/infra"
	"iproute.org/learngo/lang/interface/downloader/testing"
	"net/http"
)

// retrieve ： 取回
func retrieve(url string) string {
	resp, err := http.Get("https://ipcrystal.com")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	return string(bytes)
}

func getRetriever() infra.Retriever {
	return infra.Retriever{}
}

// type 后面接的是名字， 理解一下， type name what

// go 接口定义抽象方法不需要 func
// 就像 java 接口定义 方法不需要public

// 抽象的描述得只要名字和方法

type retriever interface {
	Get(string) string
}

func main() {
	// 1.
	// resp, err := http.Get("https://ipcrystal.com")
	//
	// if err != nil {
	// 	panic(err)
	// }
	//
	// defer resp.Body.Close()
	//
	// bytes, _ := ioutil.ReadAll(resp.Body)

	// 2.
	// fmt.Printf("%s", retrieve("https://ipcrystal.com"))

	r1 := getRetriever()

	fmt.Printf("%s", r1.Get("https://ipcrystal.com"))

	r2 := testing.Retriever{}
	fmt.Printf("%s", r2.Get("https://ipcrystal.com"))

	var r11 retriever = infra.Retriever{}
	fmt.Printf("%s", r11.Get("https://ipcrystal.com"))

	var r22 retriever = testing.Retriever{}
	fmt.Printf("%s", r22.Get("https://ipcrystal.com"))

}
