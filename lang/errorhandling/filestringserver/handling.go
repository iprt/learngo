package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func http1() {
	http.HandleFunc("/list/", func(writer http.ResponseWriter, request *http.Request) {

		// /list/fib.txt -> fib.txt
		path := request.URL.Path[len("/list/"):]

		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		all, err := ioutil.ReadAll(file)

		if err != nil {
			panic(err)
		}

		writer.Write(all)
	})

	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		panic(err)
	}

}

func http2() {

	http.HandleFunc("/list/", func(writer http.ResponseWriter, request *http.Request) {

		// /list/fib.txt -> fib.txt
		path := request.URL.Path[len("/list/"):]

		file, err := os.Open(path)
		if err != nil {
			// panic(err)
			http.Error(writer,
				err.Error(),
				http.StatusInternalServerError)
			return
		}
		defer file.Close()

		all, err := ioutil.ReadAll(file)

		if err != nil {
			panic(err)
		}

		writer.Write(all)
	})

	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		panic(err)
	}

}

func main() {
	http2()
}

/*

http1
网页访问 http://localhost:8888/list/fib.txt

再访问 http://localhost:8888/list/fib.txta
- 错误的访问，不希望panic

*/

/**
http2
访问 http://localhost:8888/list/fib.txta

*/
