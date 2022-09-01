package main

import (
	"iproute.org/learngo/lang/errorhandling/filestringserver2/filelisting"
	"log"
	"net/http"
	"os"
)

/**
第一步 type 一个 function 包装原来的function

第二步 原来的 function 错误直接返回
*/
type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

// 第三步，统一的错误处理
func errorWrapping(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// 业务处理
		err := handler(writer, request)

		// panic 保护

		log.Println("Error occurred ")

		// 统一的错误处理
		if err != nil {
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				// http.Error(writer, http.StatusText(http.StatusNotFound),
				// 	http.StatusNotFound)
				code = http.StatusNotFound
			case os.IsPermission(err):
				// 权限错误
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code),
				code)
		}
	}
}

func main() {

	// http.HandleFunc("/list/", filelisting.FileList)
	http.HandleFunc("/list/", errorWrapping(filelisting.FileList))

	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		panic(err)
	}
}
