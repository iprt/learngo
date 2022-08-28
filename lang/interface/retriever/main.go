package main

import (
	"fmt"
	"iproute.org/learngo/lang/interface/retriever/mock"
	"iproute.org/learngo/lang/interface/retriever/really"
)

type Retriever interface {
	Get(string) string
}

func download(r Retriever) string {
	return r.Get("https://ipcrystal.com")
}

func main() {
	var r Retriever = mock.Retriever{
		Content: "Hello World",
	}
	// fmt.Println("download is :", download(r))

	// %T 打印出类型  %v 打印出值
	fmt.Printf("%T %v\n", r, r)

	r = really.HttpRetriever{}
	fmt.Printf("%T %v \n", r, r)

	r = &really.HttpRetriever{}

	fmt.Println("download is :", download(r))

	fmt.Printf("%T %v \n", r, r)
}
